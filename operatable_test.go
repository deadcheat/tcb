package tcb_test

import (
	"testing"

	"github.com/deadcheat/tcb"
)

func TestGet(t *testing.T) {
	t.Log("=== Case 1. return error when key did not hit")
	o := prepareOperator()

	testKey := "test"
	var dummy interface{}
	if _, err := o.Get(testKey, dummy); err == nil {
		t.Error("Get should return error but not returned")
		t.Fail()
	}
	t.Log("=== Case 2. return data when key did hit")
	testData := 1000
	_, _ = o.Insert(testKey, testData, 0)
	if _, err := o.Get(testKey, dummy); err != nil {
		t.Error("Get should not return error but returned ", err)
		t.Fail()
	}
	_, _ = o.Remove(testKey)
	t.Log("=== Case 3. return error when bucket or operator are nil")
	var o_nil *tcb.BucketOperator
	if _, err := o_nil.Get("somekey", dummy); err == nil {
		t.Error("Get should return error but not returned")
		t.Fail()
	}
	bo := o.(*tcb.BucketOperator)
	bo.Bucket = nil
	if _, err := bo.Get("somekey", dummy); err == nil {
		t.Error("Get should return error but not returned")
		t.Fail()
	}
}

func prepareOperator() tcb.Operatable {
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
	o, _ := c.Operator("default")
	return o
}
