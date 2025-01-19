package data

import (
	"context"

	"github.com/osamikoyo/portfiler/internal/data/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type PortfolioStorage struct{
	Collection *mongo.Collection
	ctx context.Context
}

func (storage *PortfolioStorage) Add(portfolio models.Portfolio) error {
	_, err := storage.Collection.InsertOne(storage.ctx, portfolio)
	return err
}

func (storage *PortfolioStorage) Get(title string) (models.Portfolio, error) {
	filter, err := storage.Collection.Find(storage.ctx, bson.M{
		"title" : title,
	})
	if err != nil{
		return models.Portfolio{}, err
	}

	var portfolias models.Portfolio

	if err := filter.All(storage.ctx, &portfolias);err != nil{
		return models.Portfolio{}, err
	}

	return portfolias, nil
}
