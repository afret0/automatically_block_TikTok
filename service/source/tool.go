package source

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"os"
	"strings"
	"time"
)

func GetEnv() string {
	env := os.Getenv("environment")
	return env
}

func NowString() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func ConObjectIDToString(obj primitive.ObjectID) string {
	i := strings.TrimLeft(obj.Hex(), "0")
	return i
}

func ConStringToObjectID(s string) (primitive.ObjectID, error) {
	obj, err := primitive.ObjectIDFromHex(s)
	return obj, err
}
