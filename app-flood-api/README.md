Generate the proto go code using this command
protoc -I . --go_out=plugins=grpc:. ./*.proto
