package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	root := ""
	if len(os.Args) == 2 {
		root = os.Args[1]
	}

	http.Handle("/", http.FileServer(http.Dir(root)))

	if err := http.ListenAndServe("127.0.0.1:80", nil); err != nil {
		fmt.Println(err.Error())
		return
	}
}
