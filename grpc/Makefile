generate:
	protoc -I. ./proto/remote.proto \
		--go_out=plugins=grpc:$$GOPATH/src

clean:
	rm -f ./proto/remote/* 