package books

import(
	"context"
	"encoding/json"
	"net/http"
	"log"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"github.com/oasiscse/bookx/conn"
	"github.com/oasiscse/bookx/data/api"
)

func GetByYear(w http.ResponseWriter, r *http.Request) {
	var books []data.Book
	var params = mux.Vars(r)
	year := string(params["year"])
	collection := conn.ConnectDB()


	cur, err := collection.Find(context.TODO(), bson.M{"year":year})

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