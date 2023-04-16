package main

import "belajar-gin/routes"

func main() {
	var port = ":8080"

	routes.StartServer().Run(port)
}
