package gateway

import (
	"log"
	"net/http"

	greeter "root/protobuf/greeter"

	"github.com/ysugimoto/grpc-graphql-gateway/runtime"
)

func RunGQL(){
	mux := runtime.NewServeMux()

	if err := greeter.RegisterGreeterGraphql(mux); err != nil {
		log.Printf("register grpc service to graphql error: %v\n", err)
		return
	}

	log.Println("start graphql server at :8888")
	http.Handle("/graphql", mux)
	log.Fatalln(http.ListenAndServe(":8888", nil))
}

