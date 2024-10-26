package main

import (
	gateway "root/gateway"
	db "root/mongo"
	service "root/service"
)

func main() {
	db.ConnectDB()
	db.Init()
	go func(){
		gateway.RunGQL()
	}()
	service.RunGRPC()
}
