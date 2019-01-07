generate:
	protoc -I. ./proto/remote.proto \
		--gopherjs_out=plugins=grpc:$$GOPATH/src \
		--go_out=plugins=grpc:$$GOPATH/src
	go generate ../grpcweb/

install:
	go install ./vendor/github.com/golang/protobuf/protoc-gen-go \
		./vendor/github.com/johanbrandhorst/protobuf/protoc-gen-gopherjs \
		./vendor/github.com/foobaz/go-zopfli \
		./vendor/github.com/gopherjs/gopherjs

clean:
	rm -f ./proto/*.pb.go 