// +build js

package grpc

import (
	"net/url"
	"strings"

	log "github.com/sirupsen/logrus"
	"honnef.co/go/js/dom"

	pbweb "github.com/kyeett/gameserver/proto/web"
)

var document = dom.GetWindow().Document().(dom.HTMLDocument)

func init() {
	document.AddEventListener("DOMContentLoaded", false, func(_ dom.Event) {
		log.Error("Document loaded!")
	})
}

func CorrectClient(serverAddr string, secure bool) (pbweb.BackendClient, error) {
	u2, err := url.Parse(document.DocumentURI())
	if err != nil {
		log.Error("unexpected error parsing URI", err)
		return nil, err
	}

	addr := "localhost:10001"
	if u2.Query().Get("addr") != "" {
		addr = u2.Query().Get("addr")

		if !strings.Contains(addr, "http") {
			addr = "https://" + addr
		}
	}

	log.Printf("Connecting to %s\n", addr)

	clas := pbweb.NewBackendClient(serverAddr)
	document.Body().SetInnerHTML(`<div><h2>GopherJS gRPC-Web is great!</h2></div>`)

	return clas, nil
}
