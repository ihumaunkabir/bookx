## Bookx

Bookx is a simple API built in GO and MongoDB.


## Getting Started

To get a local copy up and running follow these simple steps.

### Prerequisites
* Go 1.10 and higher
* MongoDB Server 4.2.3
* MongoDriver
```sh
go get go.mongodb.org/mongo-driver/mongo
```
* gorilla/mux
```sh
go get -u github.com/gorilla/mux
```
* Cobra
```sh
go get -u github.com/spf13/cobra/cobra
```


### Installation
 
1. Clone the repo
```sh
git clone https://github.com/oasiscse/bookx.git
```
2. Go to project directory and run -
```sh
go run main.go serve
```

## Endpoints
```sh
Create Book [POST]: /api/books/{title}/{author}/{psher}/{year}/{cat}/{bookid}
Get all books [GET]: /api/books/info/all
Get All Books By Year [GET]: /api/books/info/all/{year}
Get by Book ID [GET]: /api/books/info/{bookid}
Delete single Book by ID [DELETE]: /api/books/remove/{bookid}
```
## Contributing

Contributions are what make the open source community such an amazing place to be learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Open a Pull Request

## License

Distributed under the MIT License.

## Contact

Tweet me - [@oasiscse](https://twitter.com/oasiscse)  
Find me on [ihumaun.com](http://ihumaun.com)
