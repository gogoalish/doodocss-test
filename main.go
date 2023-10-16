package main

import (
	"github.com/gogoalish/doodocs-test/internal/api"
)

func main() {
	// service.Pack()
	server := api.NewServer()
	server.Start(":8080")
	// service.Unpack("./assets/arc.zip")
}

// func main() {
// 	fmt.Println(mime.TypeByExtension("xml"))
// }
