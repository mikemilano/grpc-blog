package main

import (
	"context"
	"fmt"
	"github.com/mikemilano/grpc-blog/blog/blogpb"
	"google.golang.org/grpc"
	"io"
	"log"
)

func main() {
	fmt.Println("Blog Client")

	opts := grpc.WithInsecure()
	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer cc.Close()

	c := blogpb.NewBlogServiceClient(cc)

	// Create blog
	fmt.Println("Creating the blog")
	blog := &blogpb.Blog{
		AuthorId: "Mike",
		Title: "My First Blog",
		Content: "Content of the first blog",
	}
	createBlogRes, err := c.CreateBlog(context.Background(), &blogpb.CreateBlogRequest{Blog: blog})
	if err != nil {
		log.Fatalf("unexpected error: %v", err)
	}
	blogId := createBlogRes.Blog.GetId()

	fmt.Printf("Blog has been created: %v\n", createBlogRes)

	// read blog
	fmt.Println("Reading the blog")
	_, err2 := c.ReadBlog(context.Background(), &blogpb.ReadBlogRequest{
		BlogId: "5ec791e9555c6982a74b0eaa",
	})
	if err2 != nil {
		fmt.Printf("Error while reading: %v\n", err2)
	}

	readBlogReq := &blogpb.ReadBlogRequest{BlogId: blogId}
	readBlogRes, readBlogErr := c.ReadBlog(context.Background(), readBlogReq)
	if readBlogErr != nil {
		fmt.Printf("Error while reading: %v\n", readBlogErr)
	}

	fmt.Printf("Blog was read: %v\n", readBlogRes)

	// update blog
	fmt.Println("Updating the blog")

	updatedBlog := &blogpb.Blog{
		Id: blogId,
		AuthorId: "Mia",
		Title: "My First Blog (edited)",
		Content: "Updated content",
	}

	updateRes, updateErr := c.UpdateBlog(context.Background(), &blogpb.UpdateBlogRequest{
		Blog: updatedBlog,
	})
	if updateErr != nil {
		fmt.Printf("Error while updating: %v\n", updateErr)
	}

	fmt.Printf("blog was updated: %v\n", updateRes)

	// List blogs
	fmt.Println("Listing blogs")

	stream, err := c.ListBlog(context.Background(), &blogpb.ListBlogRequest{})

	if err != nil {
		log.Fatalf("Error while calling ListBlog RPC: %v", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Something happened: %v", err)
		}
		fmt.Println(res.GetBlog())
	}

	// Delete blog
	fmt.Println("Deleting the blog")

	deleteRes, deleteErr := c.DeleteBlog(context.Background(), &blogpb.DeleteBlogRequest{BlogId: updatedBlog.GetId()})
	if deleteErr != nil {
		fmt.Printf("Error while deleting: %v\n", deleteErr)
	}

	fmt.Printf("blog was deleted: %v\n", deleteRes)
}
