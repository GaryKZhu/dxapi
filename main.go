// server.go created by @bzhu

package main

import (
	"doctorx/routers"
	"doctorx/pkg/setting"
)


func main() {
	
	//init the router
	router := routers.InitRouter()
	//start the server
	router.Run(setting.HTTPPort)
}
