package tcb_test

import (
	"testing"

	"github.com/deadcheat/tcb"
)

// TestOpen tests tcb.Open
func TestOpenAndAddBucket(t *testing.T) {
	config := tcb.Configure{
		ConnectString: "ftp://localhost:8091",
	}
	var c tcb.CouchBaseAdapter
	c = tcb.NewCluster(config)
	if err := c.Open(); err == nil {
		t.Error("Open must return protocol error")
		t.Fail()
	}
	config.ConnectString = "http://localhost:8091"
	c = tcb.NewCluster(config)
	if err := c.Open(); err != nil {
		t.Error("Open must not return any error but:", err)
		t.Fail()
	}
	if _, err := c.AddBucket("default", ""); err == nil {
		t.Error("Before configuring auth param, must return err but:", err)
		t.Fail()
	}
	config.User = "Administrator"
	config.Password = "password"
	c = tcb.NewCluster(config)
	if err := c.Open(); err != nil {
		t.Error("Open must not return any error but:", err)
		t.Fail()
	}
	if _, err := c.AddBucket("default", ""); err != nil {
		t.Error("After configuring auth param, must not return err but:", err)
		t.Fail()
	}
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
		t.Error("Open must return invalid bucket error when bucket name not found")
		t.Fail()
	}

}
