package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/ivpusic/neo"
	"github.com/satori/go.uuid"
	"gopkg.in/redis.v2"
)

type MyJsonName struct {
	Active bool    `json:"Active"`
	Value  float64 `json:"Value"`
	Name   string  `json:"name"`
}

func (a *MyJsonName) UpperName() string {
	return strings.ToUpper(a.Name)
}

func main() {
	fmt.Println("In main")

	app := neo.App()

	route := app.Get("/", func(ctx *neo.Ctx) {
		ctx.Res.Text("I am Neo Programmer", 200)
	})

	route.Use(func(ctx *neo.Ctx, next neo.Next) {
		if 1 == 1 {
			next()
		} else {
			ctx.Res.Status = 401
		}
	})
	app.Post("/post/id/:id/key/:key", func(ctx *neo.Ctx) {
		fmt.Println("\n\tIn post\n")
		fmt.Println("\n\tUpdating image with id " + ctx.Req.Params.Get("id"))
		fmt.Println("\n\tUpdating image with key " + ctx.Req.Params.Get("key") + "\n\n\n")

		//  == * == * == * == * ==   H E A D E R S     S E C T I O N   == * == * == * == * == * == * ==

		head := ctx.Req.Header.Get("Host")

		for ky, value := range ctx.Req.Header {

			fmt.Printf("\n\tKy = %s   \n", ky)
			fmt.Printf("\tvalue = %s   \n", value)

		}

		if len(head) == 0 {
			fmt.Printf("\n\n\thead is  empty, it is %d long!! ", len(head))
		} else {
			fmt.Printf("\n\n\theader is not empty, it is %d long ", len(head))
		}

		//  == * == * == * == * ==   H E A D E R S     S E C T I O N   == * == * == * == * == * == * ==

		//  == * == * == * == * ==   J S O N      S E C T I O N   == * == * == * == * == * == * ==

		jsDesc := &MyJsonName{}

		err := ctx.Req.JsonBody(jsDesc)
		if err != nil {
			fmt.Println(err)
		}

		//  == * == * == * == * ==   J S O N      S E C T I O N   == * == * == * == * == * == * ==

		fmt.Printf("\n\n\n\tJSON Body parsed name => %s , Name set to upper is => %s 	, Value => %.2f and an active flag of %t \n", jsDesc.Name, jsDesc.UpperName(), jsDesc.Value, jsDesc.Active)

		//  == * == * == * == * ==   J S O N      S E C T I O N   == * == * == * == * == * == * ==

		//  == * == * == * == * ==   R E D I S      S E C T I O N   == * == * == * == * == * == * ==
		uuid := uuid.NewV4()
		out, err := json.Marshal(jsDesc)
		if err != nil {
			fmt.Println(err)
		}
		insertOK := AddRedis(uuid.String(), string(out))

		//  == * == * == * == * ==   R E D I S       S E C T I O N   == * == * == * == * == * == * ==

		//  == * == * == * == * ==   R E S P O N S E       S E C T I O N   == * == * == * == * == * == * ==

		fmt.Println("\n\t" + PostHandler(ctx) + " \n")
		fmt.Printf("\n\trecord inserted into Redis? => %t \n\n", insertOK)
		ctx.Res.Text(PostHandler(ctx), 200)

		//  == * == * == * == * ==   R E S P O N S E      S E C T I O N   == * == * == * == * == * == * ==

	})

	//  == * == * == * == * ==   M I D D L E W A R E      S E C T I O N   == * == * == * == * == * == * ==
	app.Use(func(ctx *neo.Ctx, next neo.Next) {
		start := time.Now()
		fmt.Printf("\n\n\n\t--> [Req] %s to %s\n\n\n", ctx.Req.Method, ctx.Req.URL.Path)

		next()

		elapsed := float64(time.Now().Sub(start) / time.Millisecond)
		fmt.Printf("\n\n\n\t<-- [Res] (%d) %s to %s Took %vms\n\n\n", ctx.Res.Status, ctx.Req.Method, ctx.Req.URL.Path, elapsed)
	})
	//  == * == * == * == * ==   M I D D L E W A R E      S E C T I O N   == * == * == * == * == * == * ==
	app.Start()

}

func PostHandler(ctx *neo.Ctx) string {

	return "I am newer Neo posting Programmer called => " + ctx.Req.Params.Get("id") + " with a key value of " + ctx.Req.Params["key"]
}

func AddRedis(uuid string, out string) bool {

	client := getRedisConnection()

	keys := client.SetNX(uuid, out)
	fmt.Printf("\n\n\tRedis key  => %s , Redis Value => %s 	with a return code of %t \n", uuid, out, keys.Val())
	return keys.Val()
}

func getRedisConnection() *redis.Client {

	return redis.NewTCPClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
