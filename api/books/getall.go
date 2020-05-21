package books

import(
	"context"
	"encoding/json"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"github.com/oasiscse/bookx/conn"
	"github.com/oasiscse/bookx/data/api"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	var books []data.Book
	collection := conn.ConnectDB()

	cur, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		conn.GetError(err, w)
		return
	}

	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {

		var book data.Book

		err := cur.Decode(&book) 
		if err != nil {
			log.Fatal(err)
		}

		books = append(books, book)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(books) 
}
