protoc ./wire/wire.proto --go_out ./ && protoc $GOPATH/src/github.com/jtremback/pb-bs/client/client.proto --go_out $GOPATH/src -I $GOPATH/src