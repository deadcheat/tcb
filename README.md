# tcb
[![Build Status](https://travis-ci.org/deadcheat/tcb.svg?branch=master)](https://travis-ci.org/deadcheat/tcb) [![Coverage Status](https://coveralls.io/repos/github/deadcheat/tcb/badge.svg?branch=master&service=github)](https://coveralls.io/github/deadcheat/tcb?branch=master) [![GoDoc](https://godoc.org/github.com/deadcheat/tcb?status.svg)](https://godoc.org/github.com/deadcheat/tcb)

takin' care of business

![1rqml0](https://user-images.githubusercontent.com/2797681/34908037-93dbb5fe-f8cc-11e7-82fb-cf60a2da6234.gif)

## what is this ?
this is an adapter library for couchbase client [gocb](https://github.com/couchbase/gocb)
i made [gocbadapter](https://github.com/deadcheat/gocbadaptor) for that before, i want refactor and refine by recreating.


## How to Use

### Install

it's just same as traditional style
`go get github.com/deadcheat/tcb`

### Example

```Go
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
```