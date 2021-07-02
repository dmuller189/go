
**
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    chat/chat.proto

protoc --go_out=chat --go_opt=paths=source_relative \
    --go-grpc_out=chat --go-grpc_opt=paths=source_relative \
    chat.proto


export PATH=$PATH:$HOME/go/bin 