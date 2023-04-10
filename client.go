package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	query := `
        query  {
            hello
        }
    `

	headers := http.Header{}
	headers.Set("Content-Type", "application/json")

	reqBody, err := json.Marshal(map[string]string{
		"query": query,
	})
	if err != nil {
		panic(err)
	}

	resp, err := http.Post("http://localhost:8080/graphql", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		panic(err)
	}

	var respData map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&respData)
	if err != nil {
		panic(err)
	}

	fmt.Println(respData)
}
