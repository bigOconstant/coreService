package main

import (
	magic "coreService/Server"
	"fmt"
)


func main() {

	var server magic.Server = magic.Server{}
	fmt.Println("Begin Serving")
	server.Serve()
		

}
