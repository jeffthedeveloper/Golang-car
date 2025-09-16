package car_model

import (
	"context"

	"github.com/jeffthedeveloper/go-car-shop/api/database/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (c *carModel) Update(ctx context.Context, id string, car *entities.TCar) (*entities.TCar, error) {
	objectId, _ := primitive.ObjectIDFromHex(id)
	mId := bson.M{"_id": objectId}
	update := bson.M{"$set": car}
	opts := options.FindOneAndUpdate()

	r := c.collection.FindOneAndUpdate(ctx, mId, update, opts.SetReturnDocument(1))
	var updatedCar *entities.TCar
	r.Decode(&updatedCar)

	return updatedCar, r.Err()
}
