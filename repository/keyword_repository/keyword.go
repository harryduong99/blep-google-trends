package keyword_repository

import (
	"context"

	"github.com/harryduong99/google-trends/config"
	"github.com/harryduong99/google-trends/databasedriver"
	"github.com/harryduong99/google-trends/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type KeywordRepo struct{}

var KeywordRepository = &KeywordRepo{}

func (keywordRepo *KeywordRepo) Store(keyword models.Keyword) error {
	collection := databasedriver.Mongo.ConnectCollection(config.DB_NAME, config.COL_KEYWORD)

	bbytes, _ := bson.Marshal(keyword)
	if !KeywordRepository.IsKeywordExisting(context.Background(), keyword.Keyword) {
		_, err := collection.InsertOne(context.Background(), bbytes)
		if err != nil {
			return err
		}

		return nil
	}
	return nil
}

func (keywordRepo *KeywordRepo) GetKeyword(keyword string) models.Keyword {
	var document models.Keyword
	collection := databasedriver.Mongo.ConnectCollection(config.DB_NAME, config.COL_KEYWORD)
	data := collection.FindOne(context.Background(), bson.M{
		"keyword": keyword,
	})
	data.Decode(&document)
	return document
}

func (keywordRepo *KeywordRepo) Delete(id primitive.ObjectID) bool {
	collection := databasedriver.Mongo.ConnectCollection(config.DB_NAME, config.COL_KEYWORD)

	filter := bson.M{"_id": id}

	_, err := collection.DeleteOne(context.TODO(), filter)
	return err == nil
}

func (keywordRepo *KeywordRepo) IsKeywordExisting(ctx context.Context, keyword string) bool {
	collection := databasedriver.Mongo.ConnectCollection(config.DB_NAME, config.COL_KEYWORD)
	var document models.Keyword
	data := collection.FindOne(ctx, bson.M{"keyword": keyword})
	err := data.Decode(&document)

	if err == mongo.ErrNoDocuments {
		return false
	}
	return true
}
