package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func GitWebHook(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	check(err)
	fmt.Println(string(body))

	switch r.Method {
	case "POST":
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

// &{
// 	Method:POST
// URL: /webhook
// 	Proto:HTTP/1.1
// 	ProtoMajor:1
// 	ProtoMinor:1
// 	Header:map[Accept-Encoding:[gzip]
// 		Content-Length:[3163]
// 		Content-Type:[application/json]
// 		User-Agent:[Go-http-client/1.1]
// 		X-Gitea-Delivery:[460ed7e6-352d-440c-815a-6fcebe519a1d]
// 		X-Gitea-Event:[push]
// 		X-Gitea-Signature:[bc5fce2c947a37217fc30bed8a448e9464f77800a8824db15e7c50f4c4635d5b]
// 		X-Github-Delivery:[460ed7e6-352d-440c-815a-6fcebe519a1d]
// 		X-Github-Event:[push]
// 		X-Gogs-Delivery:[460ed7e6-352d-440c-815a-6fcebe519a1d]
// 		X-Gogs-Event:[push]
// 		X-Gogs-Signature:[bc5fce2c947a37217fc30bed8a448e9464f77800a8824db15e7c50f4c4635d5b]]
// 	Body:0xc00044c280
// 	GetBody:<nil>
// 		ContentLength:3163
// 	TransferEncoding:[]
// 	Close:false
// 	Host:george-devfu-net:8080
// 	Form:map[]
// 	PostForm:map[]
// 	MultipartForm:<nil>
// 		Trailer:map[]
// 	RemoteAddr:10.0.0.156:39836
// RequestURI:/webhook
// 	TLS:<nil>
// 		Cancel:<nil>
// 		Response:<nil>
// 		ctx:0xc00044c2c0
// }
