package tool

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"os"
	"strings"
)

type Tool struct {
}

var tool *Tool

func init() {
	tool = new(Tool)
}

func GetTool() *Tool {
	return tool
}
func (t *Tool) GetEnv() string {
	env := os.Getenv("environment")
	return env
}

//func NowString() string {
//	return time.Now().Format("2006-01-02 15:04:05")
//}

func (t *Tool) ConObjectIDToString(obj primitive.ObjectID) string {
	i := strings.TrimLeft(obj.Hex(), "0")
	return i
}

func (t *Tool) ConStringToObjectID(s string) primitive.ObjectID {
	obj, _ := primitive.ObjectIDFromHex(s)
	return obj
}

func (t *Tool) SprintfErr(err error) string {
	return fmt.Sprintf("%s", err)
}

//func ReverseString(s string) string {
//	runes := []rune(s)
//	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
//		runes[from], runes[to] = runes[to], runes[from]
//	}
//	return string(runes)
//}
