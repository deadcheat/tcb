package tcb

import (
	"github.com/couchbase/gocb"
)

// BucketOperator operatable instance
type BucketOperator struct {
	Bucket *gocb.Bucket
}

func NewBucketOperator(b *gocb.Bucket) *BucketOperator {
	return &BucketOperator{b}
}

// Get invoke gocb.a.Bucket.Get
func (b *BucketOperator) Get(key string) (cas gocb.Cas, data interface{}, err error) {
	if a == nil || a.CouchBucket == nil {
		a.Logf("CouchBase Connections may not be establlished. skip this process.")
		return 0, nil, nil
	}
	bucket := a.Bucket
	cas, err = bucket.Get(key, data)
	if err != nil {
		a.Logf("Didn't hit any data for key: %s or err: %+v \n", key, err)
		return cas, nil, err
	}
	a.Logf("hit key: %s", key)
	return cas, data, nil
}
