package services

import (
	"github.com/SofyanHadiA/linq/core/utils"

	"gopkg.in/redis.v4"
)

var client redis.Client

type redisClient struct {
	hostAddress string
	port        string
	password    string
	client      *redis.Client
}

func RedisService(hostAddress string, port string, password string) (*redisClient, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     hostAddress + ":" + port,
		Password: password,
		DB:       0, // use default DB
	})

	_, err := client.Ping().Result()

	if err == nil {
		utils.Log.Info("Connected to redis server", hostAddress+":"+port)

		return &redisClient{
			hostAddress: hostAddress,
			port:        port,
			password:    password,
			client:      client,
		}, err
	}

	return nil, err
}

func (client redisClient) KeyNil(err error) bool {
	return err == redis.Nil
}

func (redis redisClient) Set(key string, value interface{}) error {
	err := redis.client.Set(key, value, 0).Err()
	utils.Log.Info("Set new redis key "+key+": ", err)
	return err
}

func (redis redisClient) Get(key string) (string, error) {
	result, err := redis.client.Get(key).Result()
	utils.Log.Info("Get new redis key "+key+": ", err)
	return result, err
}
