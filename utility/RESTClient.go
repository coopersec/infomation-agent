package utility

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	"bytes"
	"encoding/json"
	"io"
	"os"
)

// .....

func test() {
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

// ------------------------------  GET  ---------------------------------------------------
// Create a Resty Client
func GET(client resty.Client) {

	resp, err := client.R().
		EnableTrace().
		Get("https://httpbin.org/get")

	// Explore response object
	fmt.Println("Response Info:")
	fmt.Println("  Error      :", err)
	fmt.Println("  Status Code:", resp.StatusCode())
	fmt.Println("  Status     :", resp.Status())
	fmt.Println("  Proto      :", resp.Proto())
	fmt.Println("  Time       :", resp.Time())
	fmt.Println("  Received At:", resp.ReceivedAt())
	fmt.Println("  Body       :\n", resp)
	fmt.Println()

	// Explore trace info
	fmt.Println("Request Trace Info:")
	ti := resp.Request.TraceInfo()
	fmt.Println("  DNSLookup     :", ti.DNSLookup)
	fmt.Println("  ConnTime      :", ti.ConnTime)
	fmt.Println("  TCPConnTime   :", ti.TCPConnTime)
	fmt.Println("  TLSHandshake  :", ti.TLSHandshake)
	fmt.Println("  ServerTime    :", ti.ServerTime)
	fmt.Println("  ResponseTime  :", ti.ResponseTime)
	fmt.Println("  TotalTime     :", ti.TotalTime)
	fmt.Println("  IsConnReused  :", ti.IsConnReused)
	fmt.Println("  IsConnWasIdle :", ti.IsConnWasIdle)
	fmt.Println("  ConnIdleTime  :", ti.ConnIdleTime)
	fmt.Println("  RequestAttempt:", ti.RequestAttempt)
	fmt.Println("  RemoteAddr    :", ti.RemoteAddr.String())

	/* Output
	   Response Info:
	     Error      : <nil>
	     Status Code: 200
	     Status     : 200 OK
	     Proto      : HTTP/2.0
	     Time       : 457.034718ms
	     Received At: 2020-09-14 15:35:29.784681 -0700 PDT m=+0.458137045
	     Body       :
	     {
	       "args": {},
	       "headers": {
	         "Accept-Encoding": "gzip",
	         "Host": "httpbin.org",
	         "User-Agent": "go-resty/2.4.0 (https://github.com/go-resty/resty)",
	         "X-Amzn-Trace-Id": "Root=1-5f5ff031-000ff6292204aa6898e4de49"
	       },
	       "origin": "0.0.0.0",
	       "url": "https://httpbin.org/get"
	     }

	   Request Trace Info:
	     DNSLookup     : 4.074657ms
	     ConnTime      : 381.709936ms
	     TCPConnTime   : 77.428048ms
	     TLSHandshake  : 299.623597ms
	     ServerTime    : 75.414703ms
	     ResponseTime  : 79.337µs
	     TotalTime     : 457.034718ms
	     IsConnReused  : false
	     IsConnWasIdle : false
	     ConnIdleTime  : 0s
	     RequestAttempt: 1
	     RemoteAddr    : 3.221.81.55:443
	*/

	// ---------------------------- POST ---------------------------------------
	// Create a Resty Client
}

import (
"bytes"
"encoding/json"
"io"
"net/http"
"os"
)

// .....

type Student struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

// .....
func POST() {

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

}


func PUT() {

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
		return e
	}

	defer res.Body.Close()

	fmt.Println("response Status:", res.Status)
	// Print the body to the stdout
	io.Copy(os.Stdout, res.Body)

}
