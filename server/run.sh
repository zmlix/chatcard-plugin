./grpcwebproxy \
    --backend_addr=localhost:8888 \
    --run_tls_server=false \
    --allow_all_origins \
    --server_http_debug_port 5005 &

./server