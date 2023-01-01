package models

import (
	"github.com/netwatcherio/netwatcher-agent/checks"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type SpeedTestData struct {
	ID        primitive.ObjectID `bson:"_id, omitempty"json:"id"`
	Agent     primitive.ObjectID `bson:"agent"json:"agent"`
	Data      checks.SpeedTest   `bson:"data"json:"data"`
	Timestamp time.Time          `bson:"timestamp"json:"timestamp"`
}
