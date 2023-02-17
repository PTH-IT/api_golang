package firebasedb

import (
	"context"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/db"
	"google.golang.org/api/option"
)

func Connect(ctx context.Context) (*db.Client, error) {
	conf := &firebase.Config{
		DatabaseURL: "https://pth-it-default-rtdb.asia-southeast1.firebasedatabase.app",
	}
	opt := option.WithCredentialsFile("pth-it-firebase-adminsdk-i11h0-4333a623a3.json")
	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		return nil, err
	}
	client, err := app.Database(ctx)
	return client, err
}
