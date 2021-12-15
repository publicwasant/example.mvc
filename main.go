package main

import (
	"example/mvc/controller"
	"fmt"
	"net/http"
)

func main() {
	controller.DB_CONFIG = controller.GetDatabaseConnection()

	if controller.DB_CONFIG.Conn == nil {
		fmt.Println("Database Connection Failed!")
		return
	}

	fmt.Println("Database Connection Done.")
	http.ListenAndServe(`:3000`, GetRouters())
}

/*
	https://dev.to/mecode4food/simple-golang-mux-router-api-service-3nal
	https://medium.com/garenadev/%E0%B8%A5%E0%B8%AD%E0%B8%87%E0%B9%80%E0%B8%82%E0%B8%B5%E0%B8%A2%E0%B8%99-rest-api-%E0%B8%9A%E0%B8%99-go-%E0%B8%81%E0%B8%B1%E0%B8%99-%E0%B8%95%E0%B8%AD%E0%B8%99%E0%B8%97%E0%B8%B5%E0%B9%88-2-d4d410bd5098
*/
