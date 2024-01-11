package main

import (
	"chatcard-plugin/pb/plugin"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

var BASE_PATH string

type PluginServer struct {
	plugin.UnimplementedPluginServiceServer
}

type TreeNode struct {
	Title    string      `json:"title"`
	Key      string      `json:"key"`
	Children []*TreeNode `json:"children"`
	IsLeaf   bool        `json:"isLeaf"`
}

func getDirStructure(root string) *TreeNode {
	var tree *TreeNode

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if path == root {
			tree = &TreeNode{
				Title:    info.Name(),
				Key:      "files",
				Children: []*TreeNode{},
				IsLeaf:   false,
			}
		} else {
			relPath, err := filepath.Rel(root, path)
			if err != nil {
				return err
			}

			node := &TreeNode{
				Title:    info.Name(),
				Key:      relPath,
				Children: []*TreeNode{},
				IsLeaf:   !info.IsDir(),
			}

			parentNode, err := findParentNode(tree, filepath.Dir(relPath))
			if err != nil {
				return err
			}
			parentNode.Children = append(parentNode.Children, node)
		}

		return nil
	})

	if err != nil {
		return nil
	}

	return tree
}

func findParentNode(tree *TreeNode, path string) (*TreeNode, error) {
	if path == "." || path == string(filepath.Separator) {
		return tree, nil
	}

	for _, child := range tree.Children {
		if child.Key == path {
			return child, nil
		} else if strings.HasPrefix(path, child.Key+string(filepath.Separator)) {
			return findParentNode(child, path)
		}
	}

	return nil, fmt.Errorf("Parent node not found: %s", path)
}

func (*PluginServer) Connect(ctx context.Context, req *plugin.ConnectRequest) (*plugin.ConnectResponse, error) {
	filesPath := path.Join(BASE_PATH, "../files")
	_, err := os.Stat(filesPath)
	if err != nil {
		log.Println(err)
		os.MkdirAll(filesPath, 0755)
	}
	tree := getDirStructure(filesPath)
	directory, _ := json.Marshal(tree)
	fmt.Println(string(directory))
	fmt.Println("connect...")
	pluginsPath := path.Join(BASE_PATH, "../plugins")
	var res plugin.ConnectResponse
	_, err = os.Stat(pluginsPath)
	if err != nil {
		log.Println(err)
		os.MkdirAll(pluginsPath, 0755)
		res = plugin.ConnectResponse{
			Status:    plugin.Status_SUCCESS,
			Plugins:   []*plugin.Plugin{},
			Directory: string(directory),
			Web:       os.Getenv("WEB_HTTP") + ":" + os.Getenv("WEB_PORT"),
		}
		return &res, nil
	}
	pluginsDir, _ := os.ReadDir(pluginsPath)
	Plugins := []*plugin.Plugin{}
	for _, p := range pluginsDir {
		conf_, _ := os.Open(path.Join(pluginsPath, p.Name(), "conf.json"))
		conf := &PluginConfigure{}
		err := json.NewDecoder(conf_).Decode(conf)
		if err != nil {
			continue
		}
		pluginInfo_, err := json.Marshal(conf.Plugins)
		if err != nil {
			continue
		}
		pluginInfo := string(pluginInfo_)
		var pluginOptions []*plugin.PluginOption
		for _, option := range conf.Options {
			pluginOptions = append(pluginOptions, &plugin.PluginOption{
				Type:  plugin.PluginOptionValueType(option.Type),
				Key:   option.Key,
				Value: option.Value,
			})
		}
		Plugins = append(Plugins, &plugin.Plugin{
			Name:    conf.Name,
			Display: &conf.Display,
			Version: conf.Version,
			Info:    &pluginInfo,
			Options: pluginOptions,
		})
		res = plugin.ConnectResponse{
			Status:    plugin.Status_SUCCESS,
			Plugins:   Plugins,
			Directory: string(directory),
			Web:       os.Getenv("WEB_HTTP") + ":" + os.Getenv("WEB_PORT"),
		}
	}
	return &res, nil
}

type CallResponseJson struct {
	Message string `json:"message"`
	Log     string `json:"log"`
	Finish  bool   `json:"finish"`
	Level   int16  `json:"level"`
}

func (c CallResponseJson) Json() string {
	res, _ := json.Marshal(c)
	return string(res)
}

type CallArguments struct {
	Function Arguments `json:"function"`
	ID       string    `json:"id"`
	Index    int64     `json:"index"`
	Name     string    `json:"name"`
}

type Arguments struct {
	Name      string `json:"name"`
	Arguments string `json:"arguments"`
}

type Output struct {
	Type    string `json:"type"`
	Content string `json:"content"`
}

func RunPlugin(file string, call string, arguments string, conf *PluginConfigure, server plugin.PluginService_CallServer, wg *sync.WaitGroup) {
	// fmt.Println("RunPlugin", file, call, arguments)
	os.Chdir(path.Join("../plugins", conf.Name))
	defer func() {
		os.Chdir(path.Join("../../", "server"))
	}()
	var allCallArguments CallArguments
	err := json.Unmarshal([]byte(arguments), &allCallArguments)
	if err != nil {
		server.Send(&plugin.CallResponse{
			Status:   plugin.Status_FAILED,
			Response: CallResponseJson{Log: "调用参数不正确" + err.Error(), Level: -2, Finish: true}.Json(),
		})
		wg.Done()
		return
	}

	server.Send(&plugin.CallResponse{
		Status:   plugin.Status_PROCESS,
		Response: CallResponseJson{Log: "执行插件的前置命令:" + conf.Cmd, Level: 2, Finish: false}.Json(),
	})
	CMD_TIMEOUT, err := strconv.ParseInt(os.Getenv("CMD_TIMEOUT"), 10, 64)
	if err != nil {
		log.Println(err)
		CMD_TIMEOUT = 10
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(CMD_TIMEOUT)*time.Second)
	defer cancel()
	cmd := exec.CommandContext(ctx, os.Getenv("PLUGIN_SHELL"), os.Getenv("PLUGIN_SHELL_ARGS"), fmt.Sprintf("%v", conf.Cmd))
	cmdLog, err := cmd.Output()
	if err != nil {
		server.Send(&plugin.CallResponse{
			Status:   plugin.Status_PROCESS,
			Response: CallResponseJson{Log: err.Error(), Level: -2, Finish: true}.Json(),
		})
		wg.Done()
		return
	}
	server.Send(&plugin.CallResponse{
		Status:   plugin.Status_PROCESS,
		Response: CallResponseJson{Log: string(cmdLog), Level: 3, Finish: false}.Json(),
	})

	server.Send(&plugin.CallResponse{
		Status:   plugin.Status_PROCESS,
		Response: CallResponseJson{Log: "执行插件的" + call + "方法", Level: 2, Finish: false}.Json(),
	})

	var confCmd string
	if conf.Cmd != "" {
		cmds := strings.Split(conf.Cmd, ";")
		for _, c := range cmds {
			if c == "" {
				continue
			}
			confCmd += fmt.Sprintf("%v >> %v.log 2>&1;", c, conf.Name)
		}
	}
	RUN_TIMEOUT, err := strconv.ParseInt(os.Getenv("RUN_TIMEOUT"), 10, 64)
	if err != nil {
		log.Println(err)
		RUN_TIMEOUT = 600
	}
	ctx, cancel = context.WithTimeout(context.Background(), time.Duration(RUN_TIMEOUT)*time.Second)
	defer cancel()
	arguments_base64 := base64.StdEncoding.EncodeToString([]byte(allCallArguments.Function.Arguments))
	pythonCmd := fmt.Sprintf("%v %v %v %v %v %v %v", os.Getenv("PYTHON"), "-u", conf.Name+".py", "--call", call, "--arguments", arguments_base64)
	cmd = exec.CommandContext(ctx, os.Getenv("PLUGIN_SHELL"), os.Getenv("PLUGIN_SHELL_ARGS"), fmt.Sprintf("%v%v", confCmd, pythonCmd))
	fmt.Println(cmd.Path + " " + strings.Join(cmd.Args[1:], " "))

	out, err := cmd.StdoutPipe()
	if err != nil {
		server.Send(&plugin.CallResponse{
			Status:   plugin.Status_FAILED,
			Response: CallResponseJson{Log: err.Error(), Level: -2, Finish: true}.Json(),
		})
		wg.Done()
		return
	}

	if err := cmd.Start(); err != nil {
		server.Send(&plugin.CallResponse{
			Status:   plugin.Status_FAILED,
			Response: CallResponseJson{Log: err.Error(), Level: -2, Finish: true}.Json(),
		})
		wg.Done()
		return
	}

	success := make(chan error, 1)

	fmt.Println("start")
	go func() {
		output := Output{}
		fmt.Println("decoding...")
		dec := json.NewDecoder(out)
		for dec.More() {
			if err := dec.Decode(&output); err != nil {
				fmt.Println("decoding failed:", err)
				server.Send(&plugin.CallResponse{
					Status:   plugin.Status_FAILED,
					Response: CallResponseJson{Log: err.Error(), Level: -3, Finish: true}.Json(),
				})
				wg.Done()
				success <- err
				return
			}
			// fmt.Println(output)
			if output.Type == "stdout" || output.Type == "error" || output.Type == "text/plain" {
				server.Send(&plugin.CallResponse{
					Status:   plugin.Status_PROCESS,
					Response: CallResponseJson{Message: output.Content, Level: 3, Finish: false}.Json(),
				})
			} else if output.Type == "image/png" || output.Type == "image/jpeg" {
				server.Send(&plugin.CallResponse{
					Status:   plugin.Status_PROCESS,
					Response: CallResponseJson{Message: output.Content, Log: "[" + output.Type + "]", Level: 3, Finish: false}.Json(),
				})
			}
		}
		success <- nil
	}()

	fmt.Println("wait...")
	if err := cmd.Wait(); err != nil {
		fmt.Println("failed:", ctx, err)
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			server.Send(&plugin.CallResponse{
				Status:   plugin.Status_FAILED,
				Response: CallResponseJson{Log: fmt.Sprintf("执行超时(%vs): %v %v", os.Getenv("RUN_TIMEOUT"), ctx.Err().Error(), err.Error()), Level: -3, Finish: true}.Json(),
			})
			wg.Done()
			return
		} else {
			server.Send(&plugin.CallResponse{
				Status:   plugin.Status_FAILED,
				Response: CallResponseJson{Log: err.Error(), Level: -3, Finish: true}.Json(),
			})
			wg.Done()
			return
		}
	}
	if err := <-success; err == nil {
		fmt.Println("success")
		server.Send(&plugin.CallResponse{
			Status:   plugin.Status_SUCCESS,
			Response: CallResponseJson{Log: "执行完成", Level: 4, Finish: true}.Json(),
		})
		wg.Done()
	}
}

func (*PluginServer) Call(req *plugin.CallRequest, server plugin.PluginService_CallServer) error {
	wg := sync.WaitGroup{}
	name := req.Name
	call := req.Call
	fmt.Println("BASE_PATH", BASE_PATH)
	os.Chdir(BASE_PATH)

	py := path.Join(BASE_PATH, "../plugins", name, name+".py")
	_, err := os.Stat(py)
	if err != nil {
		server.Send(&plugin.CallResponse{
			Status:   plugin.Status_FAILED,
			Response: CallResponseJson{Log: "插件" + name + "不存在", Level: -1, Finish: true}.Json(),
		})
		return nil
	}
	conf_, err := os.Open(path.Join(BASE_PATH, "../plugins", name, "conf.json"))
	if err != nil {
		server.Send(&plugin.CallResponse{
			Status:   plugin.Status_FAILED,
			Response: CallResponseJson{Log: "插件" + name + "配置文件不存在", Level: -1, Finish: true}.Json(),
		})
		return nil
	}
	conf := &PluginConfigure{}
	err = json.NewDecoder(conf_).Decode(conf)
	if err != nil {
		server.Send(&plugin.CallResponse{
			Status:   plugin.Status_FAILED,
			Response: CallResponseJson{Log: "插件" + name + "配置文件错误", Level: -1, Finish: true}.Json(),
		})
		return nil
	}

	server.Send(&plugin.CallResponse{
		Status:   plugin.Status_PROCESS,
		Response: CallResponseJson{Log: "找到插件" + name, Level: 2, Finish: false}.Json(),
	})

	wg.Add(1)
	go RunPlugin(py, call, *req.Arguments, conf, server, &wg)
	wg.Wait()
	return nil
}

func (*PluginServer) Directory(ctx context.Context, req *plugin.DirectoryRequest) (*plugin.DirectoryResponse, error) {
	event := req.Event
	paths := req.Paths
	fmt.Println(">>>>>", event, paths)
	filesPath := path.Join(BASE_PATH, "../files")
	var res plugin.DirectoryResponse
	if event == "delete" {
		for _, file := range paths {
			err := os.Remove(path.Join(filesPath, file))
			if err != nil {
				log.Println(err)
				tree := getDirStructure(filesPath)
				directory, _ := json.Marshal(tree)
				res = plugin.DirectoryResponse{
					Status:    plugin.Status_FAILED,
					Directory: string(directory),
				}
				return &res, nil
			}
		}
	}
	tree := getDirStructure(filesPath)
	directory, _ := json.Marshal(tree)
	res = plugin.DirectoryResponse{
		Status:    plugin.Status_SUCCESS,
		Directory: string(directory),
	}
	return &res, nil
}

func download(portPtr, dirPtr string) {
	http.HandleFunc("/download/", func(w http.ResponseWriter, r *http.Request) {
		filename := r.URL.Path[len("/download/"):]
		filepath := path.Join(dirPtr, filename)
		startTime := time.Now()
		log.Printf("Received request for %s from %s at %s", filename, r.RemoteAddr, time.Now().Format("2006-01-02 15:04:05"))

		file, err := http.Dir(dirPtr).Open(filename)
		if err != nil {
			http.Error(w, "File not found", http.StatusNotFound)
			return
		}
		defer file.Close()

		fi, err := os.Stat(filepath)
		if err != nil {
			http.Error(w, "File not found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Length", fmt.Sprintf("%d", fi.Size()))
		w.Header().Set("Last-Modified", fi.ModTime().UTC().Format(http.TimeFormat))

		http.ServeContent(w, r, filename, fi.ModTime(), file)
		log.Printf("File %s sent successfully to IP %s in %.2fs\n", filename, r.RemoteAddr, time.Since(startTime).Seconds())
	})

	http.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		r.ParseMultipartForm(50 << 20)
		file, handler, err := r.FormFile("file")
		if err != nil {
			log.Println(err)
			return
		}
		defer file.Close()

		data, err := io.ReadAll(file)
		if err != nil {
			log.Println(err)
			return
		}

		for _, dir := range r.Form["dirs[]"] {
			if dir == "files" {
				dir = ""
			}
			uploadDir := path.Join(dirPtr, dir)
			log.Println(uploadDir)
			if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
				log.Println(err)
				resp := map[string]any{
					"status":  "failed",
					"message": err,
				}
				jsonData, _ := json.Marshal(resp)
				w.Header().Set("Content-Type", "application/json")
				w.Write(jsonData)
				return
			}
			filePath := path.Join(uploadDir, handler.Filename)
			err = os.WriteFile(filePath, data, 0644)
			if err != nil {
				log.Println(err)
				resp := map[string]any{
					"status":  "failed",
					"message": err,
				}
				jsonData, _ := json.Marshal(resp)
				w.Header().Set("Content-Type", "application/json")
				w.Write(jsonData)
				return
			}
		}

		resp := map[string]any{
			"status":  "success",
			"message": "成功上传文件至目录" + strings.Join(r.Form["dirs[]"], ","),
		}
		jsonData, _ := json.Marshal(resp)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	})

	fmt.Printf("Starting server on port %s, serving files from %s...\n", portPtr, dirPtr)
	http.ListenAndServe(":"+portPtr, nil)
}

func main() {
	BASE_PATH, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(BASE_PATH)
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	WEB_PORT, ok := os.LookupEnv("WEB_PORT")
	if !ok {
		WEB_PORT = "5006"
	}
	GRPC_PORT, ok := os.LookupEnv("GRPC_PORT")
	if !ok {
		GRPC_PORT = "8888"
	}
	PROXY_PORT, ok := os.LookupEnv("PROXY_PORT")
	if !ok {
		PROXY_PORT = "5005"
	}
	PROXY_TLS, ok := os.LookupEnv("PROXY_TLS")
	if !ok {
		PROXY_TLS = "false"
	}
	HTTP_MAX_TIMEOUT, ok := os.LookupEnv("HTTP_MAX_TIMEOUT")
	if !ok {
		HTTP_MAX_TIMEOUT = "3600s"
	}

	if runtime.GOOS == "linux" {
		fmt.Println("Running in linux, port ", PROXY_PORT)
		_, ok := os.LookupEnv("PLUGIN_SHELL")
		if !ok {
			os.Setenv("PLUGIN_SHELL", "bash")
		}
		_, ok = os.LookupEnv("PLUGIN_SHELL_ARGS")
		if !ok {
			os.Setenv("PLUGIN_SHELL_ARGS", "-c")
		}
		go func() {
			cmd := exec.Command("./grpcwebproxy",
				"--backend_addr=localhost:"+GRPC_PORT,
				"--run_tls_server="+PROXY_TLS,
				"--allow_all_origins",
				"--server_http_debug_port="+PROXY_PORT,
				"--server_http_max_write_timeout="+HTTP_MAX_TIMEOUT,
				"--server_http_max_read_timeout="+HTTP_MAX_TIMEOUT,
			)
			out, err := cmd.Output()
			log.Println(string(out))
			if err != nil {
				log.Fatalln(err)
			}
		}()
	} else if runtime.GOOS == "windows" {
		fmt.Println("Running in windows, port ", PROXY_PORT)
		_, ok := os.LookupEnv("PLUGIN_SHELL")
		if !ok {
			os.Setenv("PLUGIN_SHELL", "powershell")
		}
		_, ok = os.LookupEnv("PLUGIN_SHELL_ARGS")
		if !ok {
			os.Setenv("PLUGIN_SHELL_ARGS", "/c")
		}
		go func() {
			cmd := exec.Command("./grpcwebproxy.exe",
				"--backend_addr=localhost:"+GRPC_PORT,
				"--run_tls_server="+PROXY_TLS,
				"--allow_all_origins",
				"--server_http_debug_port="+PROXY_PORT,
				"--server_http_max_write_timeout="+HTTP_MAX_TIMEOUT,
				"--server_http_max_read_timeout="+HTTP_MAX_TIMEOUT,
			)
			out, err := cmd.Output()
			log.Println(string(out))
			if err != nil {
				log.Fatalln(err)
			}
		}()
	}

	filesPath := path.Join(BASE_PATH, "../files")
	pluginsPath := path.Join(BASE_PATH, "../plugins")

	for _, dir := range []string{pluginsPath, filesPath} {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			err = os.Mkdir(dir, 0755)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	go download(WEB_PORT, filesPath)
	l, _ := net.Listen("tcp", ":"+GRPC_PORT)
	server := grpc.NewServer()
	plugin.RegisterPluginServiceServer(server, &PluginServer{})
	server.Serve(l)
}
