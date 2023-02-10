package firebasedb

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	firebase "firebase.google.com/go/v4"
	"github.com/labstack/echo"
	"google.golang.org/api/option"
)

func Getfirebase(c echo.Context) error {
	ctx := context.Background()
	conf := &firebase.Config{
		DatabaseURL: "https://pth-it-default-rtdb.asia-southeast1.firebasedatabase.app",
	}
	opt := option.WithCredentialsFile("pth-it-firebase-adminsdk-i11h0-4333a623a3.json")
	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		return c.String(http.StatusOK, err.Error())
	}
	client, err := app.Database(ctx)
	if err != nil {
		log.Fatalln("Error initializing database client:", err)
	}

	// The app only has access as defined in the Security Rules
	ref := client.NewRef("user_scores")
	var get []map[string]interface{}
	if err := ref.Get(context.TODO(), &get); err != nil {
		log.Fatal(err)
	}

	jsonbody, _ := json.Marshal(get)
	return c.String(http.StatusOK, string(jsonbody))
}
func Putfirebase(c echo.Context) error {
	ctx := context.Background()
	conf := &firebase.Config{
		DatabaseURL: "https://pth-it-default-rtdb.asia-southeast1.firebasedatabase.app",
	}
	opt := option.WithCredentialsFile("pth-it-firebase-adminsdk-i11h0-4333a623a3.json")
	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		return c.String(http.StatusOK, err.Error())
	}
	client, err := app.Database(ctx)
	if err != nil {
		log.Fatalln("Error initializing database client:", err)
	}

	// The app only has access as defined in the Security Rules
	ref := client.NewRef("user_scores/" + fmt.Sprint(1))

	if err := ref.Set(context.TODO(), map[string]interface{}{"score": 40}); err != nil {
		log.Fatal(err)
	}

	fmt.Println("score added/updated successfully!")
	return c.String(http.StatusOK, "hello")
}
