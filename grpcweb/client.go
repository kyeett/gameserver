package main

import (
	"context"
	"log"
	"net/url"
	"strings"

	pbweb "github.com/kyeett/gameserver/proto/web"
	"honnef.co/go/js/dom"
	// "github.com/kyeett/grpcweb-boilerplate/proto/client"
)

// Build this snippet with GopherJS, minimize the output and
// write it to html/frontend.js.
//go:generate gopherjs build client.go -m -o html/frontend.js

// Zopfli compress static files.
////go:generate find ./html/ -name *.gz -prune -o -type f -exec go-zopfli {} +

// Integrate generated JS into a Go file for static loading.
////go:generate bash -c "go run assets_generate.go"

// This constant is very useful for interacting with
// the DOM dynamically
var document = dom.GetWindow().Document().(dom.HTMLDocument)

// Define no-op main since it doesn't run when we want it to
func main() {}

// Ensure our setup() gets called as soon as the DOM has loaded
func init() {
	document.AddEventListener("DOMContentLoaded", false, func(_ dom.Event) {
		go setup()
	})
}

// Setup is where we do the real work.
func setup() {

	u2, err := url.Parse(document.DocumentURI())
	if err != nil {
		log.Fatal("unexpected error parsing URI", err)
		return
	}

	addr := "localhost:10001"
	if u2.Query().Get("addr") != "" {
		addr = u2.Query().Get("addr")

		if !strings.Contains(addr, "http") {
			addr = "https://" + addr
		}
	}

	log.Printf("Connecting to %s\n", addr)

	clas := pbweb.NewBackendClient(addr)
	document.Body().SetInnerHTML(`<div><h2>GopherJS gRPC-Web is great!</h2></div>`)

	resp, err := clas.NewPlayer(context.Background(), &pbweb.Empty{})
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp, "YAY mf")
}
