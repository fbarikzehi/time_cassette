package repository

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/yaml.v3"
)

type Configs struct {
	ConnectionUrl string `yaml:"connectionUrl"`
	DatabaseName  string `yaml:"databaseName"`
	UserPass      string `yaml:"userPass"`
}

func DbConfiguration() (Configs, error) {
	var configs Configs
	configFile, _ := os.ReadFile("./configs/db.config.yaml")
	err := yaml.Unmarshal(configFile, &configs)
	if err != nil {
		log.Fatalf("DbConfiguration faild: %v", err)
		return configs, err
	}

	return configs, nil
}

func DBinstance() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	config, err := DbConfiguration()
	if err != nil {
		log.Fatalf("DBinstance faild: %v", err)
		return nil
	}
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.ConnectionUrl))
	if err != nil {
		fmt.Println("Something is wrrong to Connecte to databse!")
		log.Fatal(err)
		return nil
	}
	fmt.Println("Connected to databse!")
	return client
}

var Client *mongo.Client = DBinstance()

func Db(client *mongo.Client) *mongo.Database {
	config, err := DbConfiguration()
	if err != nil {
		log.Fatalf("DBinstance faild: %v", err)
		return nil
	}
	var collection *mongo.Database = client.Database(config.DatabaseName)
	return collection
}

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = Db(client).Collection(collectionName)
	return collection
}
