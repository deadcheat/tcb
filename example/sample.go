package main

import (
	"fmt"

	"github.com/deadcheat/tcb"
)

type Data struct {
	Text string
}

func main() {
	config := tcb.Configure{
		ConnectString: "couchbase://localhost",
		BucketConfigs: []tcb.BucketConfig{
			tcb.BucketConfig{
				Name:     "default",
				Password: "",
			},
		},
		User:     "Administrator",
		Password: "password",
	}
	c := tcb.NewCluster(config)
	if err := c.Open(); err != nil {
		panic(err)
	}
	o, _ := c.Operator("default")
	data := Data{
		Text: "TestData",
	}
	if _, err := o.Insert("key", &data, 0); err != nil {
		panic(err)
	}
	var m Data
	if _, err := o.Get("key", &m); err != nil {
		panic(err)
	}
	fmt.Println("data : ", m)
	if _, err := o.Remove("key"); err != nil {
		panic(err)
	}
}
