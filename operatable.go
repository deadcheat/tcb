package tcb

import (
	"log"

	"github.com/couchbase/gocb"
)

// BucketOperator operatable instance
type BucketOperator struct {
	Bucket *gocb.Bucket
	Loggerable
}

// NewBucketOperator return new operatable instance
func NewBucketOperator(b *gocb.Bucket, l Loggerable) *BucketOperator {
	return &BucketOperator{b, l}
}

// Get invoke gocb.Bucket.Get
func (b *BucketOperator) Get(key string, data interface{}) (cas gocb.Cas, err error) {
	if b == nil || b.Bucket == nil {
		log.Printf("CouchBase Connections may not be establlished. skip this process.")
		return 0, nil
	}
	bucket := b.Bucket
	cas, err = bucket.Get(key, data)
	if err != nil {
		b.Logf("Didn't hit any data for key: %s or err: %+v \n", key, err)
		return cas, err
	}
	b.Logf("hit key: %s", key)
	return cas, nil
}

// Insert invoke gocb.Bucket.Insert
func (b *BucketOperator) Insert(k string, d interface{}, e uint32) (cas gocb.Cas, err error) {
	return b.update(insert, k, d, e)
}

// Upsert invoke gocb.Bucket.Upsert
func (b *BucketOperator) Upsert(k string, d interface{}, e uint32) (cas gocb.Cas, err error) {
	return b.update(upsert, k, d, e)
}

type updateMode int

const (
	insert updateMode = iota
	upsert
)

func (b *BucketOperator) update(mode updateMode, key string, data interface{}, expire uint32) (cas gocb.Cas, err error) {
	bucket := *b.Bucket
	switch mode {
	case insert:
		cas, err = bucket.Insert(key, data, expire)
	case upsert:
		cas, err = bucket.Upsert(key, data, expire)
	}
	if err != nil {
		b.Logf("Couldn't send data for key: %s or err: %+v \n", key, err)
		return cas, err
	}
	b.Logf("sent data to b.CouchBucket key: %s", key)
	return cas, nil
}

// N1qlQuery prepare query and execute
func (b *BucketOperator) N1qlQuery(q string, params interface{}) (r gocb.QueryResults, err error) {
	return b.N1qlQueryWithMode(nil, q, params)
}

// Remove remove data
func (b *BucketOperator) Remove(key string) (cas gocb.Cas, err error) {
	var dummy interface{}
	bucket := b.Bucket
	if cas, err = bucket.Get(key, dummy); err != nil {
		return cas, err
	}
	if cas, err = bucket.Remove(key, cas); err != nil {
		b.Logf("Couldn't remove any data for key: %s by err: %+v \n", key, err)
	}
	return cas, nil
}

// N1qlQuery prepare query and execute
func (b *BucketOperator) N1qlQueryWithMode(m *gocb.ConsistencyMode, q string, params interface{}) (r gocb.QueryResults, err error) {
	if b == nil || b.Bucket == nil {
		return nil, nil
	}
	nq := gocb.NewN1qlQuery(q)
	if m != nil {
		nq.Consistency(*m)
	}
	bucket := *b.Bucket
	r, err = bucket.ExecuteN1qlQuery(nq, params)
	if err != nil {
		b.Logf("Couldn't execute query for query: %s params: %+v or err: %+v \n", q, params, err)
		return r, err
	}
	b.Logf("succeeded to execute query: %s , params: %+v", q, params)
	return r, err
}
