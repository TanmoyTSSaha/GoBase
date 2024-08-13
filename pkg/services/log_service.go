package services

import (
	"context"
	"encoding/json"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type LogService struct {
	Collection *mongo.Collection
}

type LogStruct struct {
	URL              string        `json:"url"`
	Timestamp        time.Time     `json:"timestamp"`
	Method           string        `json:"method"`
	StatusCode       int           `json:"status_code"`
	ResponseDuration time.Duration `json:"response_duration"`
	Message          string        `json:"message"`
}

func NewLogService(db *mongo.Database) *LogService {
	return &LogService{Collection: db.Collection("APIGatewayLogs")}
}

func (s *LogService) StoreLogs(logStruct LogStruct) error {
	var logEntry map[string]interface{}
	jsonBytes, err := json.Marshal(logStruct)
	if err != nil {
		return err
	}

	err = json.Unmarshal(jsonBytes, &logEntry)
	if err != nil {
		return err
	}

	_, err = s.Collection.InsertOne(context.Background(), logEntry)
	if err != nil {
		return err
	}

	return nil
}
