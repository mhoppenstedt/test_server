// package main

// import (
// 	"fmt"
// 	"net/http"
// )

// func indexHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Whoa, Go is neat!")
// }

// func aboutHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Mitchell Hoppenstedt")
// }

// func main() {
// 	http.HandleFunc("/", indexHandler)
// 	http.HandleFunc("/about/", aboutHandler)
// 	http.ListenAndServe(":8000", nil)
// }
