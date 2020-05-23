server:
	go run blog/blog_server/server.go

client:
	go run blog/blog_client/client.go

mongo:
	docker run -it --rm -p 27017:27017 --name mongo mongo

protoc:
	protoc blog/blogpb/blog.proto --go_out=plugins=grpc:.

update-deps:
	go get -u github.com/golang/protobuf/proto
	go get -u google.golang.org/grpc
	go get github.com/mongodb/mongo-go-driver/mongo
