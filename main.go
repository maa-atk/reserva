package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

////
var client *mongo.Client

type Meet struct {
	ID    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title string             `json:"title,omitempty" bson:"Title,omitempty"`
	Start int                `json:"start,omitempty" bson:"Start,omitempty"`
	End   int                `json:"end," bson:"end,omitempty"`
	Ts    int                `json:"ts" bson:"ts,omitempty"`
}

func CreateMeeting(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var meet Meet
	_ = json.NewDecoder(request.Body).Decode(&meet)
	collection := client.Database("Meetings").Collection("Meet")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, _ := collection.InsertOne(ctx, meet)
	json.NewEncoder(response).Encode(result)
}

///
func main() {
	// client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://Reserva:abhi2000@cluster0.xgscm.mongodb.net/meetings?retryWrites=true&w=majority"))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// err = client.Connect(ctx)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer client.Disconnect(ctx)
	// err = client.Ping(ctx, readpref.Primary())
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// databases, err := client.ListDatabaseNames(ctx, bson.M{})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// //fmt.Println(databases)

	// meetingsDatabase := client.Database("Meetings")
	//meetCollection := meetingsDatabase.Collection("Meet")
	//participantCollection := meetingsDatabase.Collection("Participant")

	// //INSERT
	// meetResult, err := meetCollection.InsertOne(ctx, bson.D{
	// 	{Key: "title", Value: "Meet1"},
	// 	{Key: "start", Value: "9"},
	// 	{Key: "end", Value: "10"},
	// 	// {Key: "participants", Value: bson.M{
	// 	// 	"Name":  "abhi",
	// 	// 	"Email": "abc@123",
	// 	// 	"RSVP":  "Y"}},
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// partResult, err := participantCollection.InsertMany(ctx, []interface{}{
	// 	bson.D{
	// 		{Key: "Meet", Value: meetResult.InsertedID},
	// 		{Key: "Name", Value: "Abhi"},
	// 		{Key: "RSVP", Value: "Y"},
	// 	},
	// 	bson.D{
	// 		{Key: "Meet", Value: meetResult.InsertedID},
	// 		{Key: "Name", Value: "Aj"},
	// 		{Key: "RSVP", Value: "N"},
	// 	},
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("ppl %v in \n", len(partResult.InsertedIDs))

	//retrieve participants

	// cursor, err := participantCollection.Find(ctx, bson.M{})
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// defer cursor.Close(ctx)
	// for cursor.Next(ctx) {
	// 	var participants bson.M
	// 	if err = cursor.Decode(&participants); err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	//fmt.Println(participants["Name"])
	// }

	//some specific

	// filterCursor, err := participantCollection.Find(ctx, bson.M{"Name": "Abhi"})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// var participantsFiltered []bson.M
	// if err = filterCursor.All(ctx, &participantsFiltered); err != nil {
	// 	log.Fatal(err)
	// }
	//fmt.Println(participantsFiltered)

	//sorting
	// opts := options.Find()
	// opts.SetSort(bson.D{{Key: "duration", Value: -1}})
	// sortCursor, err := participantCollection.Find(ctx, bson.D{{"duration", bson.D{{"$gt", 24}}}}, opts)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// var episodesSorted []bson.M
	// if err = sortCursor.All(ctx, &episodesSorted); err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(episodesSorted)

	fmt.Println("Starting the application...")
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://Reserva:abhi2000@cluster0.xgscm.mongodb.net/meetings?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	router := mux.NewRouter()
	router.HandleFunc("/meeting", CreateMeeting).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}
