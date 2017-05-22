package main

import (
	"fmt"
	"time"

	"encoding/json"

	"github.com/go-redis/redis"
)

var client *redis.Client

func ExampleNewClient() {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pug, err := client.Ping().Result()
	fmt.Println(pug, err)
	// Output: PONG <nil>
}

type Ex struct {
	Foo  int       `json:"foo"`
	Bar  string    `json:"bar"`
	Time time.Time `json:"time"`
}

func ExampleClient() {

	ex := Ex{
		Foo:  134,
		Bar:  "onetwothree",
		Time: time.Now(),
	}

	exJson, err := json.Marshal(ex)
	if err != nil {
		panic(err)
	}

	sjson := string(exJson)

	err = client.Set("key", sjson, 0).Err()
	if err != nil {
		panic(err)
	}

	fmt.Println(sjson)

	jsonval, err := client.Get("key").Bytes()
	if err != nil {
		panic(err)
	}

	var res *Ex

	json.Unmarshal(jsonval, res)

	// val := &Ex{}

	// err = json.Unmarshal(jsonval, val)
	// if err != nil {
	// 	panic(err)
	// }

	fmt.Println("key", res)
}

func main() {
	ExampleNewClient()
	ExampleClient()
}
