package main

import (
	"context"
	"log"



	"google.golang.org/grpc"

	"healthcheck/rpc"
	"os"
)


func main() {


	if len(os.Args) < 2 {
		log.Println("grpc server address is missing.\nUsages: ./grpc_health_check localhost:51000")
		os.Exit(1)
		return
	}
	server := os.Args[1]
	service := ""

	if len(os.Args) > 2 {
		service = os.Args[2]
	}


	// Set up a connection to the server.
	conn, err := grpc.Dial(server, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := rpc.NewGrpcHealthClient(conn)

	ok, err := client.Check(context.Background(),service)
		if !ok || err != nil {
			log.Printf("can't connect grpc server ( %v ) for service ( %v ) : \n%v, \ncode: %v\n", server,
				service, err, grpc.Code(err))
			os.Exit(1)
		} else {
			log.Printf("connect the grpc server ( %v ) for service ( %v ) successfully",server,service)
			os.Exit(0)
			}
}
