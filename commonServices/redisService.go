package commonServices

import (
	"context"
	"encoding/json"
	"log"
	"restaurent-mng-bc/database"
	"time"
)

func SetStringVal(ctx context.Context, key string, value any, timeOut int) {

	rClient := database.RedisClient
	//set redis
	parseJson, _ := json.Marshal(value)
	err := rClient.Set(ctx, key, string(parseJson), time.Duration(timeOut)*time.Minute).Err()
	if err != nil {
		log.Println("Error in store string value in redis")
		log.Println("Error:", err.Error())
	}
}
func GetRedisStringValue(ctx context.Context, key string, value interface{}) error {
	rClient := database.RedisClient
	rResult, err := rClient.Get(ctx, key).Result()
	if err != nil {
		return err
	}
	_ = json.Unmarshal([]byte(rResult), &value)
	return nil
}

func DeleteRedisStringValue(ctx context.Context, key string) error {
	rClient := database.RedisClient
	_, err := rClient.Del(ctx, key).Result()
	if err != nil {
		return err
	}
	return nil
}
