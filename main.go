package main

import (
	gateway "root/gateway"
	db "root/mongo"
	service "root/service"
)

func main() {
	db.ConnectDB()
	service.RunGRPC()
	gateway.RunGQL()
}
