package main
import (
	"tokenAuthGolang/server"
	"fmt"
	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load()
}
func main() {
	server.LoadRoutes()
	fmt.Println("hello world")
}	