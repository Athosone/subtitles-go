package infrastructure

import (
	"context"

	"github.com/mongodb/mongo-go-driver/mongo"
)

func failOnError(error string) {
	panic(error)
}

func Connect() {
	client, _ := mongo.Connect(context.TODO(), "mongodb://subzero:15choup30flix@ds223343.mlab.com:23343/heroku_bvv4blc7")
	print(client.ConnectionString)
}
