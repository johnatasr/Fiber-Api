# Fiber Api Example

Api developed to test the microframework Fiber, Fiber is a built-in of FastHttp very simple and fast than others
frameworks like NodeJs, Nestjs or FastAPI. 

## Requisites

* Docker

## Tecnologies

* Go
* GoRM 
* Fiber
* SQLite

## Starts up

Steps to configure the project:

1. `cd` in the root path
2. `docker-compose up --build`


As default the project will run in localhost:8000 

## How to use ? 

### Get books

```
curl --request GET \
  --url http://localhost:8000/api/books

```

### Get books by ID

```
curl --request GET \
  --url http://localhost:8000/api/book/<ID>

```

### Insert a book

```
curl --request POST \
  --url http://localhost:8000/api/book-insert \
  --header 'Content-Type: application/json' \
  --data '{
	"title": "Bible", 
	"author": "God", 
	"rating": 10
}'

```

### Delete a book

```
curl --request DELETE \
  --url http://localhost:8000/api/book-delete/<ID>

```


