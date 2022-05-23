package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Student struct {
	Username string `json:"name"`
	Password string `json:"address"`
}

func main() {

	body := &Student{
		Username: "jt",
		Password: "foobar",
	}

	payloadBuf := new(bytes.Buffer)
	json.NewEncoder(payloadBuf).Encode(body)
	req, _ := http.NewRequest("POST", "google.com", payloadBuf)

	client := &http.Client{}
	res, e := client.Do(req)
	if e != nil {
		fmt.Println("12")
	}

	defer res.Body.Close()

	fmt.Println("response Status:", res.Status)
	// Print the body to the stdout
	io.Copy(os.Stdout, res.Body)

}
