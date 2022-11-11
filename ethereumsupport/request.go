package ethereumsupport

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"time"
)

const (
	HeaderContent = "Content-Type"
	HeaderAccept  = "Accept-Language"
	HeaderHost    = "Host"
)

type Request struct {
	Url      string `json:"url"`
	Code     int    `json:"code"`
	ErrorMsg error  `json:"error_msg"`
	Response *http.Response
}

type Headerhandler map[string]string

func Get(url string, tlsEnable bool, header Headerhandler) *Request {
	request := new(Request)
	fmt.Println(url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		request.ErrorMsg = err
	}

	var tr http.Transport
	if tlsEnable {
		tr = http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	}

	client := &http.Client{Timeout: 30 * time.Second, Transport: &tr}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	if val, ok := header[HeaderContent]; ok {
		resp.Header.Add(HeaderContent, val)
	}

	if val, ok := header[HeaderAccept]; ok {
		resp.Header.Add(HeaderContent, val)
	}

	if val, ok := header[HeaderHost]; ok {
		resp.Header.Add(HeaderContent, val)
	}

	//defer resp.Body.Close()
	request.Response = resp
	request.Code = resp.StatusCode
	return request
}

/*
func Post(url string, header Headerhandler, data []byte) *Request {
	req := new(Request)

	resp, err := http.NewRequest("POST", url, bytes.NewReader(data))
	if err != nil {
		req.ErrorMsg = err
	}
	if val, ok := header[HeaderContent]; ok {
		resp.Header.Add(HeaderContent, val)
	}

	if val, ok := header[HeaderAccept]; ok {
		resp.Header.Add(HeaderContent, val)
	}

	if val, ok := header[HeaderHost]; ok {
		resp.Header.Add(HeaderContent, val)
	}

	defer resp.Body.Close()
	req.Response = resp
	return req
}
*/
