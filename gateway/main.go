package main

import (
	"log"

	"net/http"

	greeter "root/protobuf/greeter"

	"github.com/ysugimoto/grpc-graphql-gateway/runtime"
)

func main() {
	mux := runtime.NewServeMux()

	if err := greeter.RegisterGreeterGraphql(mux); err != nil {
		log.Fatalln(err)
	}
	http.Handle("/graphql", mux)
	log.Fatalln(http.ListenAndServe(":8888", nil))
}
