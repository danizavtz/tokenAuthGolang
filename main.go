package main
import (
	"tokenAuthGolang/server"
	"fmt"
)
func main() {
	server.LoadRoutes()
	fmt.Println("hello world")
}	