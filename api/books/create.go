package books

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/oasiscse/bookx/conn"
	data "github.com/oasiscse/bookx/data/api"
)

func Create(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var book data.Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	collection := conn.ConnectDB()
	result, err := collection.InsertOne(context.TODO(), book)

	if err != nil {
		w.Write([]byte(`{"message" : "Can not Create book"}`))
	} else {
		json.NewEncoder(w).Encode(result)
	}

}
