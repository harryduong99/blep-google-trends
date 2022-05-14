package trend_repository

import (
	"context"

	"github.com/harryduong99/google-trends/config"
	"github.com/harryduong99/google-trends/databasedriver"
	"github.com/harryduong99/google-trends/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TrendRepo struct{}

var TrendRepository = &TrendRepo{}

func (TrendRepo *TrendRepo) Store(trend models.Trend) error {
	collection := databasedriver.Mongo.ConnectCollection(config.DB_NAME, config.COL_TREND)

	bbytes, _ := bson.Marshal(trend)
	if !TrendRepository.IsRecordExisting(context.Background(), trend.KeywordID, trend.Time) {
		_, err := collection.InsertOne(context.Background(), bbytes)
		if err != nil {
			return err
		}

		return nil
	} else {

	}
	return nil
}

func (TrendRepo *TrendRepo) IsRecordExisting(ctx context.Context, keywordID primitive.ObjectID, time string) bool {
	collection := databasedriver.Mongo.ConnectCollection(config.DB_NAME, config.COL_TREND)
	var document models.Keyword
	data := collection.FindOne(ctx, bson.M{"time": time, "keyword_id": keywordID})
	err := data.Decode(&document)

	if err == mongo.ErrNoDocuments {
		return false
	}
	return true
}
