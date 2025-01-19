package data

import (
	"github.com/osamikoyo/portfiler/internal/data/models"
	"go.mongodb.org/mongo-driver/bson"
)

func (storage *PortfolioStorage) AddReview(title string, review models.Review) error {
	filter := bson.M{
		"title" : title,
	}

	update := bson.M{
		"$push" : bson.M{
			"reviews" : review,
		},
	}

	_, err := storage.Collection.UpdateOne(storage.ctx, filter, update)
	return err
}