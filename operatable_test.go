package tcb_test

import (
	"testing"

	"github.com/couchbase/gocb"

	"github.com/deadcheat/tcb"
)

func TestGet(t *testing.T) {
	t.Log("=== Case 1. return error when key did not hit")
	o := prepareOperator()

	testKey := "test1"
	_, _ = o.Remove(testKey)
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

func TestInsertAndUpsert(t *testing.T) {
	t.Log("=== Case 1. Insert sucesses to insert")
	o := prepareOperator()

	testKey := "test2"
	_, _ = o.Remove(testKey)
	testData := 1000
	if _, err := o.Insert(testKey, testData, 0); err != nil {
		t.Error("Insert should not return error but returned ", err)
		t.Fail()
	}
	t.Log("=== Case 2. Insert failed to insert when key is duplicated")
	if _, err := o.Insert(testKey, testData, 0); err == nil {
		t.Error("Insert should return error but not returned ")
		t.Fail()
	}
	t.Log("=== Case 3. Upsert sucesses to upsert data")
	upsertData := 2000
	if _, err := o.Upsert(testKey, upsertData, 0); err != nil {
		t.Error("Insert should not return error but returned ", err)
		t.Fail()
	}
	t.Log("=== Case 4. Upsert/Insert return error when nil operator")
	var o_nil *tcb.BucketOperator
	if _, err := o_nil.Upsert(testKey, upsertData, 0); err != tcb.ErrOperationUnenforceable {
		t.Error("Upsert/Insert should return error but not returned or is not ErrOperationUnenforceable", err)
		t.Fail()
	}
	_, _ = o.Remove(testKey)
}

func TestRemove(t *testing.T) {
	t.Log("=== Case 1. Remove sucesses to remove")
	o := prepareOperator()

	testKey := "test3"
	testData := 1000
	_, _ = o.Remove(testKey)
	_, _ = o.Insert(testKey, testData, 0)
	if _, err := o.Remove(testKey); err != nil {
		t.Error("Remove should not return error but returned ", err)
		t.Fail()
	}
	t.Log("=== Case 2. Remove failed to remove when data has been removed already")
	if _, err := o.Remove(testKey); err == nil {
		t.Error("Remove should return error but not returned ")
		t.Fail()
	}
	t.Log("=== Case 3. Remove return error when nil operator")
	var o_nil *tcb.BucketOperator
	if _, err := o_nil.Remove(testKey); err != tcb.ErrOperationUnenforceable {
		t.Error("Remove should return error but not returned or is not ErrOperationUnenforceable", err)
		t.Fail()
	}
	_, _ = o.Remove(testKey)
}

func TestN1qlQuery(t *testing.T) {
	t.Log("=== Case 1. N1qlQuery sucesses to execute query")
	c := prepareCluster()
	b := c.Bucket("default")
	if err := b.Manager("", "").CreatePrimaryIndex("", true, false); err != nil {
		t.Error("prepare primary index failed")
		t.Fail()
	}
	o := tcb.NewBucketOperator(b, tcb.NewDefaultActiveLogger())

	testKey := "test4"
	testData := 1000
	_, _ = o.Insert(testKey, testData, 0)
	if _, err := o.N1qlQuery("SELECT * FROM default d WHERE meta(d).id = $1", []interface{}{testKey}); err != nil {
		t.Error("N1qlQuery should not return error but returned ", err)
		t.Fail()
	}

	t.Log("=== Case 2. N1qlQueryWithMode sucesses to execute query")
	if _, err := o.N1qlQueryWithMode(gocb.RequestPlus, "SELECT * FROM default d WHERE meta(d).id = $1", []interface{}{testKey}); err != nil {
		t.Error("N1qlQuery should not return error but returned ", err)
		t.Fail()
	}

	t.Log("=== Case 3. N1qlQueryWithMode return error when nil operator")
	var o_nil *tcb.BucketOperator
	if _, err := o_nil.N1qlQueryWithMode(gocb.RequestPlus, "SELECT * FROM default d WHERE meta(d).id = $1", []interface{}{testKey}); err != tcb.ErrOperationUnenforceable {
		t.Error("Remove should return error but not returned or is not ErrOperationUnenforceable", err)
		t.Fail()
	}

	t.Log("=== Case 4. N1qlQueryWithMode fail to execute query")
	if _, err := o.N1qlQueryWithMode(gocb.RequestPlus, "ELECT * FROM default d WHERE meta(d).id = $1", []interface{}{testKey}); err == nil {
		t.Error("N1qlQuery should return error but returned nil")
		t.Fail()
	}
	_, _ = o.Remove(testKey)
}
func prepareOperator() tcb.Operatable {
	c := prepareCluster()
	o, err := c.Operator("default")
	if err != nil {
		panic(err)
	}
	return o
}

func prepareCluster() *tcb.Cluster {
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
	return c
}
