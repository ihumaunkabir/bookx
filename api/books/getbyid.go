package books

import(
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"github.com/oasiscse/bookx/conn"
	"github.com/oasiscse/bookx/data/api"
)

func GetByID(w http.ResponseWriter, r *http.Request){
	var book data.Book
	var params = mux.Vars(r)
	bookid := string(params["bookid"])
	collection := conn.ConnectDB()

	filter := bson.M{"bookid": bookid}
	err := collection.FindOne(context.TODO(), filter).Decode(&book)

	if err != nil {
		w.Write([]byte(`{"message": "Book not found, param=id"}`))
	} else{
		json.NewEncoder(w).Encode(book)
	}

}