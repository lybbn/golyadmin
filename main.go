package main
 
import (
	"fmt"
	"net/http"
)
 
func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello Go-Vue-Lyadmin")
}
 
func main() {
 
	myMux := http.NewServeMux()
 
	myMux.HandleFunc("/", sayHello)
 
	http.ListenAndServe(":9000", myMux)
}