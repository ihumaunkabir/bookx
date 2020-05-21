## Bookx

Bookx is a simple API built in GO and MongoDB.

 - **Requirements:** 
    - Go
    - MongoDB
    - MongoDriver
    - gorilla/mux

#### Endpoints:
Create Book [POST]: /api/books/{title}/{author}/{psher}/{year}/{cat}/{bookid}
Get all books [GET]: /api/books/info/all
Get All Books By Year [GET]: /api/books/info/all/{year}
Get by Book ID [GET]: /api/books/info/{bookid}
Delete single Book by ID [DELETE]: /api/books/remove/{bookid}
