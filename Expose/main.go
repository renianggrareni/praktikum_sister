package main

import (
	"fmt"
	"net/http"
)

func main(){
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.request){
	fmt.Fprint(w, "Hello World!")
}