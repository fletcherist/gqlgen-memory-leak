package main

import (
	"fmt"
	"gqlgen-memory-leak/graph"
	"gqlgen-memory-leak/graph/generated"
	"runtime"
	"time"

	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
)

func main() {
	c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}})))

	count := 0
	go func() {
		for {
			var m runtime.MemStats
			runtime.ReadMemStats(&m)

			fmt.Printf("count: %d", count)
			fmt.Printf("\tgoroutines num: %d", runtime.NumGoroutine())
			fmt.Printf("\tAlloc = %v MiB", m.Alloc/1024/1024)
			fmt.Printf("\tTotalAlloc = %v MiB", m.TotalAlloc/1024/1024)
			fmt.Printf("\tSys = %v MiB", m.Sys/1024/1024)
			fmt.Printf("\tNumGC = %v\n", m.NumGC)

			time.Sleep(time.Second * 1)
		}
	}()

	// var msg struct {
	// 	resp struct {
	// 		Test struct {
	// 			ID   string
	// 			Text string
	// 			Done bool
	// 		}
	// 	}
	// 	err error
	// }

	for {
		func() {
			count += 1
			sub := c.Websocket(`subscription { test(name:"test") { id text done } }`)
			defer sub.Close()
			time.Sleep(time.Millisecond * 50)
			// msg.err = sub.Next(&msg.resp)
			// log.Println(msg)
		}()
	}
}
