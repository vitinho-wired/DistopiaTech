package main

import (
	"fmt"
	"net/http"
)

func hellohandler(resp http.ResponseWriter, req *http.Request) {
	resp.WriteHeader(http.StatusOK)
	fmt.Fprint(resp, "Hello world")
	return
}

func main() {
	http.HandleFunc("/", hellohandler)
	if err := http.ListenAndServe(":9999", nil); err != nil {
		fmt.Println(err.Error())
	}

}
