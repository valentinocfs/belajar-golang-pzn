# Golang Redis

## Redis
Redis adalah database yang menggunakan key-value store. Redis memiliki beberapa jenis data yang bisa disimpan, seperti string, hash, list, set, zset, dan bitmap.

## Redis Client
Redis client adalah library yang digunakan untuk menghubungkan ke Redis server. Redis client biasanya menggunakan protocol Redis. Di golang, ada beberapa redis client yang bisa digunakan, seperti redigo, go-redis, dan redis.

## Go Redis
Go Redis adalah redis client yang paling populer di golang. Go Redis menggunakan redis protocol. Untuk menginstall go redis, bisa menggunakan perintah:
> go get github.com/redis/go-redis/v9

## Client
> object redis.NewClient()

## Command
Tiap command di redis menggunakan format PascalCase, contoh:
> client.Get("key")
> client.Ping()

## String
> client.Set("key", "value", 0)

> client.Get("key")

> client.Del("key")

> client.SetEx("key", "value", 10)

> client.SetNX("key", "value", 10)

> client.SetXX("key", "value", 10)

## List
> client.RPush("key", "value")

> client.LRange("key", 0, -1)

> client.LPop("key")

> client.RPop("key")

> client.BLPop("key")

> client.RPopLPush("key1", "key2")

## Set
> client.SAdd("key", "value")

> client.SMembers("key")

> client.SIsMember("key", "value")

> client.SRem("key", "value")

> client.SUnion("key1", "key2")

> client.SInter("key1", "key2")

> client.SDiff("key1", "key2")

> client.SIsMember("key", "value")

## Sorted Set
> client.ZAdd(ctx, "scores", redis.Z{
	Score: 10,
	Member: "John",
})

> client.ZRange(ctx, "scores", 0, -1)

> client.ZRank(ctx, "scores", "John")

> client.ZRevRange(ctx, "scores", 0, -1)

> client.ZRevRank(ctx, "scores", "John")

> client.ZScore(ctx, "scores", "John")

> client.ZRem(ctx, "scores", "John")

> client.ZCard(ctx, "scores")

> client.ZPopMax(ctx, "scores")

> client.ZPopMin(ctx, "scores")

> client.ZRangeByScore(ctx, "scores", "0", "10")

> client.ZRevRange(ctx, "scores", 0, -1)

## Hash
> client.HSet(ctx, "user:1", "name", "John")

> client.HGet(ctx, "user:1", "name")

> client.HGetAll(ctx, "user:1")

> client.HDel(ctx, "user:1", "name")

> client.HExists(ctx, "user:1", "name")

> client.HIncrBy(ctx, "user:1", "age", 1)

> client.HIncrByFloat(ctx, "user:1", "age", 1.5)

> client.HSetNX(ctx, "user:1", "name", "John")

## Geo Point
> client.GeoAdd(ctx, "locations", &redis.GeoLocation{
	Name: "John",
	Lon:  106.8456,
	Lat:  -6.2088,
})

> client.GeoRadius(ctx, "locations", 106.8456, -6.2088, &redis.GeoRadiusQuery{
	Radius: 1000,
	Unit:   "m",
})

> client.GeoDist(ctx, "locations", "John", "Jane", "km")

> client.GeoSearch(ctx, "locations", &redis.GeoSearchQuery{
	
})

## Hyper Log Log
> client.PFAdd(ctx, "visitors", "john", "jane", "bob")

> client.PFCount(ctx, "visitors")

## Pipeline
Pipeline digunakan untuk mengirim beberapa command ke redis dalam satu kali request.
> client.Pipeline(callback)

## Transaction
> client.TxPipelined(ctx, func(pipeliner redis.Pipeliner) error {
	
})

> client.Watch(ctx, func(tx *redis.Tx) error {
	
})

## Stream
> client.XAdd(ctx, &redis.XAddArgs{
	Stream: "logs",
	Values: []interface{}{"message", "Hello World"},
})

> client.XRead(ctx, &redis.XReadArgs{
	Streams: []string{"logs", "0"},
})

## PubSub
> client.Subscribe(ctx, "channel")

> client.Publish(ctx, "channel", "message")
