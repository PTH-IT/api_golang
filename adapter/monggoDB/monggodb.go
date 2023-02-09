package monggodb

import (
	"PTH-IT/api_golang/config"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getmongo(c echo.Context) error {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	uri := fmt.Sprintf(config.Getconfig().Monggo.Host, config.Getconfig().Monggo.User, config.Getconfig().Monggo.Pass)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	coll := client.Database(config.Getconfig().Monggo.Db).Collection("movies")
	var result []bson.M

	s, err := coll.Find(context.TODO(), bson.D{})
	s.All(context.TODO(), &result)
	if err != nil {
		panic(err)
	}
	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", jsonData)
	return c.String(http.StatusOK, string(jsonData))
}
func putmongo(c echo.Context) error {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	uri := fmt.Sprintf(config.Getconfig().Monggo.Host, config.Getconfig().Monggo.User, config.Getconfig().Monggo.Pass)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	coll := client.Database(config.Getconfig().Monggo.Db).Collection("movies")
	title := "Back to the Future"

	coll.InsertOne(context.TODO(), bson.D{{"title", title}})

	return c.String(http.StatusOK, "susscess")
}
