grpcurl -protoset <(buf build -o -) -plaintext -d '{"name": "Jane"}'  localhost:8080 greet.v1.GreetService/Greet