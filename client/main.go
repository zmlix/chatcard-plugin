package main

import (
	"chatcard-plugin/pb/plugin"
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	l, _ := grpc.Dial("127.0.0.1:8888", grpc.WithTransportCredentials(insecure.NewCredentials()))
	client := plugin.NewPluginServiceClient(l)
	res, err := client.Connect(context.Background(), &plugin.ConnectRequest{})
	if err != nil {
		log.Println(err)
	}
	fmt.Println(res.Status, res.Plugins)

	args := `{"index":0,"id":"call_v9PfhwyjkYZ4bmLDIjrhQhZD","type":"function","function":{"name":"python","arguments":"from sympy import symbols, Eq, solve\n\nx, y = symbols('x y')\n\neq1 = Eq(x, 0.3*x + 5*y)\neq2 = Eq(100*x + 200*y, 1000)\n\neq2 = eq2.subs(x, solve(eq1, x)[0])\n\nsolution = solve(eq2, y)\n\nbrick_weight = solve(eq1.subs(y, solution[0]), x)[0]\n\nbrick_weight"}}`
	// args = `{"index":0,"id":"call_dMr1diDNhVwxsyI3RT8gDbrO","type":"function","function":{"name":"python","arguments":"import matplotlib.pyplot as plt\nimport numpy as np\n\n# Generate x values from 0 to 2*pi\nx = np.linspace(0, 2*np.pi, 100)\n\n# Calculate y values for sin(x)\ny = np.sin(x)\n\n# Plot the sin function\nplt.plot(x, y)\nplt.xlabel('x')\nplt.ylabel('sin(x)')\nplt.title('Plot of sin(x)')\nplt.grid(True)\nplt.show()"}}`
	c, _ := client.Call(context.Background(), &plugin.CallRequest{
		Name:      "codeInterpreter",
		Call:      "python",
		Arguments: &args,
	})

	for {
		req, err := c.Recv()
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf("%v\n%v\n\n", req.Status, req.Response)
	}

}
