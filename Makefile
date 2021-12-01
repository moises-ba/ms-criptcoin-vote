GOPATH=/home/moises/go
PATH := $(Shell pwd)/$(GOPATH)/bin:$(PATH)

gen:
	protoc -I. -I${GOPATH}/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative  --grpc-gateway_out ./ --swagger_out=:swagger   criptcoinvote/*.proto

clean:
	rm criptcoinvote/*.go && rm -f swagger/*