# gRPC Blog Demo

A simple gRPC blog demo.

Refer to the `Makefile` for project commands.

```bash
# Run Mongodb via Docker
make mongo

# Install/update the deps
make update-deps

# Run protoc
make protoc

# Run the server
make server

# Run the client
make client
```

## Evans CLI

[Evans](https://github.com/ktr0731/evans) is a universal gRPC client.

This example shows opening evans CLI and calling `CreateBlog` and `ListBlog`.

```bash
$ evans -r

  ______
 |  ____|
 | |__    __   __   __ _   _ __    ___
 |  __|   \ \ / /  / _. | | '_ \  / __|
 | |____   \ V /  | (_| | | | | | \__ \
 |______|   \_/    \__,_| |_| |_| |___/

 more expressive universal gRPC client

blog.BlogService@127.0.0.1:50051> call CreateBlog
blog::id (TYPE_STRING) => 
blog::author_id (TYPE_STRING) => Mike
blog::title (TYPE_STRING) => My second blog
blog::content (TYPE_STRING) => Content is king.
{
  "blog": {
    "id": "5ec86923664e455b67548166",
    "authorId": "Mike",
    "title": "My First Blog",
    "content": "Content is king."
  }
}

blog.BlogService@127.0.0.1:50051> call ListBlog
{
  "blog": {
    "id": "5ec8626cf8d99ff8cebff856",
    "authorId": "Mike",
    "title": "My First Blog",
    "content": "Content is king."
  }
}

```
