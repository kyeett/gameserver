generate:
	protoc -I. -Ivendor/ ./proto/remote.proto \
		--gopherjs_out=plugins=grpc,Mgoogle/protobuf/empty.proto=github.com/johanbrandhorst/protobuf/ptypes/empty:$$GOPATH/src \
		--go_out=plugins=grpc:$$GOPATH/src
	go generate ./grpcweb/

clean:
	rm -f ./proto/web/* ./proto/*.pb.go \
		./grpcweb/html/frontend.js ./grpcweb/html/frontend.js.map

install:
	go install ./vendor/github.com/golang/protobuf/protoc-gen-go \
		./vendor/github.com/johanbrandhorst/protobuf/protoc-gen-gopherjs \
		./vendor/github.com/foobaz/go-zopfli \
		./vendor/github.com/gopherjs/gopherjs

generate_cert:
	go run "$$(go env GOROOT)/src/crypto/tls/generate_cert.go" \
		--host=localhost,127.0.0.1 \
		--ecdsa-curve=P256 \
		--ca=true

serve:
	go run main.go
