package utility

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

// .....

type Student struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

func GET() {
	c := http.Client{Timeout: time.Duration(1) * time.Second}
	resp, err := c.Get("https://www.google.com")
	if err != nil {
		fmt.Printf("Error %s", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("Body : %s", body)
}

// .....
func POST(url string) error {

	body := &Student{
		Name:    "abc",
		Address: "xyz",
	}

	payloadBuf := new(bytes.Buffer)
	json.NewEncoder(payloadBuf).Encode(body)
	req, _ := http.NewRequest("POST", url, payloadBuf)

	client := &http.Client{}
	res, e := client.Do(req)
	if e != nil {
		return e
	}

	defer res.Body.Close()

	fmt.Println("response Status:", res.Status)
	// Print the body to the stdout
	io.Copy(os.Stdout, res.Body)

	return e

}

func PUT(url string) {

	body := &Student{
		Name:    "abc",
		Address: "xyz",
	}

	payloadBuf := new(bytes.Buffer)
	json.NewEncoder(payloadBuf).Encode(body)
	req, _ := http.NewRequest("PUT", url, payloadBuf)

	client := &http.Client{}
	res, e := client.Do(req)
	if e != nil {
		fmt.Printf("%s", "hello")
	}

	defer res.Body.Close()

	fmt.Println("response Status:", res.Status)
	// Print the body to the stdout
	io.Copy(os.Stdout, res.Body)

}

func Test(){
	fmt.Printf("%s", "string")
}
