package books

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/oasiscse/bookx/conn"
	"go.mongodb.org/mongo-driver/bson"
)

//DeleteByID removes data from database
func DeleteByID(w http.ResponseWriter, r *http.Request) {
	var params = mux.Vars(r)
	bookid := params["bookid"]
	collection := conn.ConnectDB()

	filter := bson.M{"bookid": bookid}
	delRes, err := collection.DeleteOne(context.TODO(), filter)

	if err != nil {
		//dbcon.GetError(err, w)
		//return
		w.Write([]byte(`{"message": "Something went wrong."}`))
	} else {
		json.NewEncoder(w).Encode(delRes)
	}
}
