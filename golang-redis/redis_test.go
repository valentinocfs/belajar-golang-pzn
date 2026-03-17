package main

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/go-openapi/testify/assert"
	"github.com/redis/go-redis/v9"
)

var client = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "",
	DB:       0,
})

func TestConnection(t *testing.T) {
	assert.NotNil(t, client)

	// err := client.Close()
	// assert.Nil(t, err)
}

var ctx = context.Background()

func TestPing(t *testing.T) {
	result, err := client.Ping(ctx).Result()
	assert.Nil(t, err)
	assert.Equal(t, "PONG", result)
	fmt.Println(result)
}

func TestString(t *testing.T) {
	client.Set(ctx, "name", "John Doe", 3*time.Second)
	result, err := client.Get(ctx, "name").Result()
	assert.Nil(t, err)
	assert.Equal(t, "John Doe", result)
	fmt.Println(result)

	time.Sleep(4 * time.Second)

	result, err = client.Get(ctx, "name").Result()
	assert.Equal(t, redis.Nil, err)
	fmt.Println(result)
}

func TestList(t *testing.T) {
	client.RPush(ctx, "names", "John")
	client.RPush(ctx, "names", "Doe")
	client.RPush(ctx, "names", "Kurniawan")

	assert.Equal(t, "John", client.LPop(ctx, "names").Val())
	assert.Equal(t, "Doe", client.LPop(ctx, "names").Val())
	assert.Equal(t, "Kurniawan", client.LPop(ctx, "names").Val())

	client.Del(ctx, "names")
}

func TestSet(t *testing.T) {
	client.SAdd(ctx, "students", "John")
	client.SAdd(ctx, "students", "Jane")
	client.SAdd(ctx, "students", "Adams")
	client.SAdd(ctx, "students", "John")
	client.SAdd(ctx, "students", "Adams")

	assert.Equal(t, int64(3), client.SCard(ctx, "students").Val())
	assert.Equal(t, []string{"John", "Jane", "Adams"}, client.SMembers(ctx, "students").Val())

	client.Del(ctx, "students")
}

func TestSortedSet(t *testing.T) {
	client.ZAdd(ctx, "scores", redis.Z{
		Score:  100,
		Member: "John",
	})
	client.ZAdd(ctx, "scores", redis.Z{
		Score:  90,
		Member: "Jane",
	})
	client.ZAdd(ctx, "scores", redis.Z{
		Score:  80,
		Member: "Adams",
	})

	assert.Equal(t, int64(3), client.ZCard(ctx, "scores").Val())
	assert.Equal(t, []string{"Adams", "Jane", "John"}, client.ZRange(ctx, "scores", 0, -1).Val())

	assert.Equal(t, "John", client.ZPopMax(ctx, "scores").Val()[0].Member)
	assert.Equal(t, "Jane", client.ZPopMax(ctx, "scores").Val()[0].Member)
	assert.Equal(t, "Adams", client.ZPopMax(ctx, "scores").Val()[0].Member)
}

func TestHash(t *testing.T) {
	client.HSet(ctx, "user:1", "name", "John")
	client.HSet(ctx, "user:1", "email", "john@gmail.com")
	client.HSet(ctx, "user:1", "age", 30)

	user := client.HGetAll(ctx, "user:1").Val()

	assert.Equal(t, "John", user["name"])
	assert.Equal(t, "john@gmail.com", user["email"])
	assert.Equal(t, "30", user["age"])

	client.Del(ctx, "user:1")
}

func TestGeoPoint(t *testing.T) {
	client.GeoAdd(ctx, "locations", &redis.GeoLocation{
		Name:      "John",
		Longitude: 106.8456,
		Latitude:  -6.2088,
	})

	assert.Equal(t, "John", client.GeoRadius(ctx, "locations", 106.8456, -6.2088, &redis.GeoRadiusQuery{
		Radius: 1000,
		Unit:   "m",
	}).Val()[0].Name)

	client.GeoAdd(ctx, "sellers", &redis.GeoLocation{
		Name:      "Seller A",
		Longitude: 106.8456,
		Latitude:  -6.2088,
	})

	client.GeoAdd(ctx, "sellers", &redis.GeoLocation{
		Name:      "Seller B",
		Longitude: 106.9457,
		Latitude:  -6.2089,
	})

	fmt.Println(client.GeoDist(ctx, "sellers", "Seller A", "Seller B", "km").Val())

	fmt.Println(client.GeoSearch(ctx, "sellers", &redis.GeoSearchQuery{
		Longitude:  106.8456,
		Latitude:   -6.2088,
		Radius:     20,
		RadiusUnit: "km",
	}).Val())
}

func TestHyperLogLog(t *testing.T) {
	client.PFAdd(ctx, "visitors", "john", "jane", "bob")
	client.PFAdd(ctx, "visitors", "john", "jane", "bob", "budi")
	client.PFAdd(ctx, "visitors", "john", "jane", "bob", "budi", "siti")
	client.PFAdd(ctx, "visitors", "john", "jane", "bob", "budi", "siti", "ani")

	assert.Equal(t, int64(6), client.PFCount(ctx, "visitors").Val())
}

func TestPipeline(t *testing.T) {
	_, err := client.Pipelined(ctx, func(pipeliner redis.Pipeliner) error {
		pipeliner.SetEx(ctx, "name", "John", 5*time.Second)
		pipeliner.SetEx(ctx, "email", "john@gmail.com", 5*time.Second)
		pipeliner.SetEx(ctx, "age", 30, 5*time.Second)
		return nil
	})

	assert.Nil(t, err)
	assert.Equal(t, "John", client.Get(ctx, "name").Val())
	assert.Equal(t, "john@gmail.com", client.Get(ctx, "email").Val())
	assert.Equal(t, "30", client.Get(ctx, "age").Val())
}

func TestTransaction(t *testing.T) {
	_, err := client.TxPipelined(ctx, func(pipeliner redis.Pipeliner) error {
		pipeliner.SetEx(ctx, "name", "John", 5*time.Second)
		pipeliner.SetEx(ctx, "email", "john@gmail.com", 5*time.Second)
		pipeliner.SetEx(ctx, "age", 30, 5*time.Second)
		return nil
	})

	assert.Nil(t, err)
	assert.Equal(t, "John", client.Get(ctx, "name").Val())
	assert.Equal(t, "john@gmail.com", client.Get(ctx, "email").Val())
}

func TestPublishStream(t *testing.T) {
	for i := 0; i < 10; i++ {
		err := client.XAdd(ctx, &redis.XAddArgs{
			Stream: "members",
			Values: map[string]interface{}{
				"name":    "John Doe",
				"country": "Indonesia",
			},
		})
		assert.Nil(t, err)
	}
}

func TestCreateConsumerGroup(t *testing.T) {
	err := client.XGroupCreate(ctx, "members", "group-1", "0").Err()
	assert.Nil(t, err)

	err = client.XGroupCreateConsumer(ctx, "members", "group-1", "consumer-1").Err()
	assert.Nil(t, err)

	err = client.XGroupCreateConsumer(ctx, "members", "group-1", "consumer-2").Err()
	assert.Nil(t, err)
}

func TestGetStream(t *testing.T) {
	result := client.XReadGroup(ctx, &redis.XReadGroupArgs{
		Group:    "group-1",
		Consumer: "consumer-1",
		Streams:  []string{"members", ">"},
		Count:    2,
		Block:    5 * time.Second,
	}).Val()

	for _, stream := range result {
		for _, message := range stream.Messages {
			fmt.Println(message.Values)
		}
	}
}

func TestSubscribePubSub(t *testing.T) {
	subscriber := client.Subscribe(ctx, "channel-1")
	defer subscriber.Close()

	for i := 0; i < 10; i++ {
		message, err := subscriber.ReceiveMessage(ctx)
		assert.Nil(t, err)
		fmt.Println(message.Payload)
	}
}

func TestPublishPubSub(t *testing.T) {
	for i := 0; i < 10; i++ {
		err := client.Publish(ctx, "channel-1", "Hello "+strconv.Itoa(i)).Err()
		assert.Nil(t, err)
		fmt.Println("Published message", i+1)
	}
}
