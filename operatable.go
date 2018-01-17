package tcb

import (
	"github.com/couchbase/gocb"
)

// BucketOperator operatable instance
type BucketOperator struct {
	Bucket *gocb.Bucket
	Loggerable
}

func NewBucketOperator(b *gocb.Bucket, l Loggerable) *BucketOperator {
	return &BucketOperator{b, l}
}

// Get invoke gocb.a.Bucket.Get
func (b *BucketOperator) Get(key string) (cas gocb.Cas, data interface{}, err error) {
	if b == nil || b.Bucket == nil {
		b.Logf("CouchBase Connections may not be establlished. skip this process.")
		return 0, nil, nil
	}
	bucket := b.Bucket
	cas, err = bucket.Get(key, data)
	if err != nil {
		b.Logf("Didn't hit any data for key: %s or err: %+v \n", key, err)
		return cas, nil, err
	}
	b.Logf("hit key: %s", key)
	return cas, data, nil
}
