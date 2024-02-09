package main
import "fmt"
func main() {
http.HandleFunc("/", handler)
fmt.Println("Hello")
}
