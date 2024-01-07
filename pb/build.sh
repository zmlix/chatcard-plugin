export PATH="$PATH:/home/n104/go/bin"
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./plugin/plugin.proto
# protoc --js_out=import_style=commonjs,binary:. --grpc-web_out=import_style=commonjs,mode=grpcwebtext:. ./plugin/plugin.proto