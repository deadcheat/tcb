package tcb_test

import (
	"reflect"
	"testing"

	"github.com/deadcheat/tcb"
)

// TestOpen tests tcb.Open
func TestOpenAndAddBucket(t *testing.T) {
	t.Log("=== Case 1. Open should return error when illegal protocol given")
	config := tcb.Configure{
		ConnectString: "ftp://localhost:8091",
	}
	var c tcb.CouchBaseAdapter
	c = tcb.NewCluster(config)
	if err := c.Open(); err == nil {
		t.Error("Open should return protocol error")
		t.Fail()
	}
	t.Log("=== Case 2. Open should not return any error")
	config.ConnectString = "http://localhost:8091"
	c = tcb.NewCluster(config)
	if err := c.Open(); err != nil || c.Cluster() == nil {
		t.Error("Open should not return any error but:", err)
		t.Fail()
	}
	t.Log("=== Case 3. AddBucket should return error when auth params are not given in spite of auth is required")
	if _, err := c.AddBucket("default", ""); err == nil {
		t.Error("Before configuring auth param, should return err but:", err)
		t.Fail()
	}
	config.User = "Administrator"
	config.Password = "password"
	c = tcb.NewCluster(config)
	if err := c.Open(); err != nil {
		t.Error("Open should not return any error but:", err)
		t.Fail()
	}
	t.Log("=== Case 4. AddBucket should not return any error when auth params are given")
	if _, err := c.AddBucket("default", ""); err != nil {
		t.Error("After configuring auth param, should not return err but:", err)
		t.Fail()
	}
	t.Log("=== Case 5. AddBucket should return error when illegal bucket name given")
	config.BucketConfigs = []tcb.BucketConfig{
		tcb.BucketConfig{
			Name:     "default",
			Password: "",
		},
		tcb.BucketConfig{
			Name:     "not-found-bucket",
			Password: "",
		},
	}
	c = tcb.NewCluster(config)
	if err := c.Open(); err == nil {
		t.Error("Open should return invalid bucket error when bucket name not found")
		t.Fail()
	}
	t.Log("=== Case 6. AddBucket should return a registered bucket when same bucket name given")
	config.BucketConfigs = []tcb.BucketConfig{
		tcb.BucketConfig{
			Name:     "default",
			Password: "",
		},
	}
	c = tcb.NewCluster(config)
	_ = c.Open()
	buc := c.Bucket("default")
	dup, _ := c.AddBucket("default", "")
	if !reflect.DeepEqual(buc, dup) {
		t.Error("AddBucket should return a registered bucket but returned", buc, "and", dup)
		t.Fail()
	}
}

func TestOperator(t *testing.T) {
	t.Log("=== Case 1. Operator should return operatable instance successfully")
	config := tcb.Configure{
		ConnectString: "http://localhost:8091",
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
	_ = c.Open()
	var o tcb.Operatable
	var err error
	o, err = c.Operator("default")
	if err != nil || o == nil {
		t.Error("Operator should not return any error but ", err)
		t.Fail()
	}
	t.Log("=== Case 2. Operator should return tcb.ErrBucketMissing when unknown bucket name given")

	o, err = c.Operator("unknown")
	if (err == nil || err != tcb.ErrBucketMissing) || o != nil {
		t.Error("Operator should return tcb.ErrBucketMissing but", err, o)
		t.Fail()
	}
}
