from functools import reduce
import jupyter_client
import re
import argparse
import base64
import json

class JupyterKernel:
    def __init__(self, work_dir):
        (
            self.kernel_manager,
            self.kernel_client,
        ) = jupyter_client.manager.start_new_kernel(kernel_name="python3")
        self.work_dir = work_dir
        self.create_work_dir()

    def execute_code(self, code):
        self.kernel_client.execute(code)
        res_list = []
        while True:
            try:
                iopub_msg = self.kernel_client.get_iopub_msg(timeout=1)
                res_list.append(
                    {"msg_type": iopub_msg["msg_type"], "content": iopub_msg["content"]}
                )
                if (
                    iopub_msg["msg_type"] == "status"
                    and iopub_msg["content"].get("execution_state") == "idle"
                ):
                    break
            except:
                continue
        res_list = filter(
            lambda x: x["msg_type"]
            in ["stream", "execute_result", "display_data", "error"],
            res_list,
        )
        res_list = map(
            lambda x: (
                [
                    {
                        "type": "error",
                        "content": re.compile(r"(\x9B|\x1B\[)[0-?]*[ -\/]*[@-~]").sub(
                            "", "\n".join(x["content"].get("traceback", []))
                        ),
                    }
                ]
                if x["msg_type"] == "error"
                else (
                    [{"type": x["content"]["name"], "content": x["content"]["text"]}]
                    if x["msg_type"] == "stream"
                    else [
                        {"type": data[0], "content": data[1]}
                        for data in x["content"].get("data", {}).items()
                    ]
                )
            ),
            res_list,
        )
        res_list = reduce(lambda x, y: x + y, res_list, [])
        res_list = list(res_list)
        # for res in res_list:
        #     print(res)
        return json.dumps(res_list)

    def create_work_dir(self):
        code = (
            f"import os\n"
            f"if not os.path.exists('{self.work_dir}'):\n"
            f"    os.makedirs('{self.work_dir}')\n"
            f"os.chdir('{self.work_dir}')\n"
            f"del os"
        )
        self.execute_code(code)


def getArgs():
    parser = argparse.ArgumentParser(description="Python代码解释器")
    parser.add_argument("--call", type=str, default="")
    parser.add_argument("--arguments", type=str, default="")
    args_ = parser.parse_args()
    return {
        "call": args_.call,
        "arguments": base64.b64decode(args_.arguments).decode("utf-8"),
    }


def run_python_code(code):
    jk = JupyterKernel("../../files/codeInterpreter")
    out = jk.execute_code(code)
    jk.kernel_client.shutdown()
    print(out)


if __name__ == "__main__":
    args = getArgs()
    call_functions = {"python": run_python_code}
    call_functions[args["call"]](args["arguments"])
