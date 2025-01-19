package data

import (
	"context"
	"time"

	"github.com/osamikoyo/portfiler/internal/data/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type PortfolioStorage struct{
	Collection *mongo.Collection
	ctx context.Context
}

func New() (*PortfolioStorage, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	conn, err := mongo.Connect(ctx)
	defer conn.Disconnect(ctx)

	collection := conn.Database("portfiles").Collection("portfiles")
	return &PortfolioStorage{
		Collection: collection,
		ctx: ctx,
	}, err
}

func (storage *PortfolioStorage) Add(portfolio models.Portfolio) error {
	_, err := storage.Collection.InsertOne(storage.ctx, portfolio)
	return err
}

func (storage *PortfolioStorage) Get(title string) ([]models.Portfolio, error) {
	filter, err := storage.Collection.Find(storage.ctx, bson.M{
		"title" : title,
	})
	if err != nil{
		return nil, err
	}

	var portfolias []models.Portfolio

	if err := filter.All(storage.ctx, &portfolias);err != nil{
		return nil, err
	}

	return portfolias, nil
}
