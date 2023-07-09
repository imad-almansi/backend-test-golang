package handlers

import (
	"fmt"
	"net/http"
)

func HandleFacts(rw http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()

	typeFilter := query.Get("type")
	if typeFilter != "" {
		fmt.Printf("%s\n", typeFilter)
		return
	}

	fmt.Printf("no filter\n")
}
