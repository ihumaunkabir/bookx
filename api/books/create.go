package books

import(
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/oasiscse/bookx/conn"
	"github.com/oasiscse/bookx/data/api"
)

func Create(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type", "application/json")
	
	pms := mux.Vars(r)

	author := pms["author"]
	title := pms["title"]
	psher := pms["psher"]
	year := pms["year"]
	cat := pms["cat"]
	bookid := pms["bookid"]

	book := data.Book{
		Title : title,
		Publisher : psher,
		Author : author,
		Year : year,
		Category : cat,
		BookID : bookid,
	}	

	collection := conn.ConnectDB()
	result, err := collection.InsertOne(context.TODO(), book)

	if err != nil {
		w.Write([]byte(`{"message" : "Can not crate book"}`))
	} else{
		json.NewEncoder(w).Encode(result)
	}

}