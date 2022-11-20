# fiber-mongo-example

A simple CRUD app that showcases how you can use Fiber and MongoDB together

### instructions

1. ensure you have the latest version of Go installed
2. ensure that you have MongoDB installed and running
3. clone this repo
4. create a .env file in the root of the project and add the following:

```
MONGO_URI=mongodb://localhost:27017
PORT=8080
```

5. run `go run main.go` in the root of the project

### endpoints

#### GET /books

Returns all books

#### GET /books/:id

Returns a single book

#### POST /books

Creates a new book

input:

```
{
    "title": "test book",
    "author": "me",
    "year": "2022"
}
```

#### PUT /books/:id

Updates a book

input:

```
{
    "title": "test book",
    "author": "me",
    "year": "2022"
}
```

#### DELETE /books/:id

Deletes a book
