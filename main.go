package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var v = `
	2023/03/17 17:13:47 request received from: &http.Request{
	  Method:"GET",
		URL:(*url.URL)(0xc000196000),
		Proto:"HTTP/1.1",
		ProtoMajor:1,
		ProtoMinor:1,
		Header:http.Header{"Accept":[]string{"*/*"},
		"User-Agent":[]string{"curl/7.79.1"}},
		Body:http.noBody{},
		GetBody:(func() (io.ReadCloser, error))(nil),
		ContentLength:0,
		TransferEncoding:[]string(nil),
		Close:false,
		Host:"example.com",
		Form:url.Values(nil),
		PostForm:url.Values(nil),
		MultipartForm:(*multipart.Form)(nil),
		Trailer:http.Header(nil),
		RemoteAddr:"127.0.0.1:53501",
		RequestURI:"/",
		TLS:(*tls.ConnectionState)(nil),
		Cancel:(<-chan struct {})(nil),
		Response:(*http.Response)(nil),
		ctx:(*context.cancelCtx)(0xc000184050),
	}
`

func main() {
	err := http.ListenAndServe(":80", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		headers := map[string]any{
			"Proto": r.Proto,
			"Host":  r.Host,
		}
		for key, value := range r.Header {
			if len(value) <= 1 {
				headers[key] = r.Header.Get(key)
				continue
			}
			headers[key] = r.Header[key]
		}
		b, err := json.Marshal(headers)
		if err != nil {
			log.Printf("unable to marshal headers as JSON: %s", err)
		}

		log := fmt.Sprintf("received request: %s %s%s\nwith headers: %s", r.Method, r.Host, r.URL, string(b))

		fmt.Print(log)
		fmt.Fprint(w, log)
	}))
	if err != nil {
		log.Printf("error when running the server: %s", err)
	}
}
