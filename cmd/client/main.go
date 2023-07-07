package main

import (
    "context"
    "log"
    "net/http"

    greetv1 "example/gen/greet/v1"
    "example/gen/greet/v1/greetv1connect"

    "github.com/bufbuild/connect-go"
)

func main() {
    client_with_connect_protocol()
    client_with_grpc_protocol()
    client_with_grpcweb_protocol()
}

func client_with_connect_protocol(){

    client := greetv1connect.NewGreetServiceClient(
        http.DefaultClient,
        "http://localhost:8080",
    )

    res, err := client.Greet(
        context.Background(),
        connect.NewRequest(&greetv1.GreetRequest{Name: "Jane"}),
    )
    if err != nil {
        log.Println(err)
        return
    }
    log.Println("With connect protocol:",res.Msg.Greeting)
}

func client_with_grpc_protocol(){

    client := greetv1connect.NewGreetServiceClient(
        http.DefaultClient,
        "http://localhost:8080",
        connect.WithGRPC(),
    )

    res, err := client.Greet(
        context.Background(),
        connect.NewRequest(&greetv1.GreetRequest{Name: "Jane"}),
    )
    if err != nil {
        log.Println(err)
        return
    }
    log.Println("With grpc protocol:",res.Msg.Greeting)
}

func client_with_grpcweb_protocol(){

    client := greetv1connect.NewGreetServiceClient(
        http.DefaultClient,
        "http://localhost:8080",
        connect.WithGRPCWeb(),
    )

    res, err := client.Greet(
        context.Background(),
        connect.NewRequest(&greetv1.GreetRequest{Name: "Jane"}),
    )
    if err != nil {
        log.Println(err)
        return
    }
    log.Println("With grpc-web protocol:",res.Msg.Greeting)
}