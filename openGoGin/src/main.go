package main

import "routers"

func main() {

	router := routers.RegisterRouters()

	router.Run(":8080")
}
