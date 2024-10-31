package core

import (
	"context"
	"fmt"
	"testing"
)

func TestEEtcd(t *testing.T) {
	client := InitEtcd("127.0.0.1:2379")
	res, err := client.Put(context.Background(), "auth_api2", "127.0.0.1:20021")
	fmt.Println(res, err)
	res1, err1 := client.Get(context.Background(), "auth_api2")
	fmt.Println(res1, err1)
	for _, ev := range res1.Kvs {
		fmt.Printf("%s:%s\n", ev.Key, ev.Value)
	}
}
