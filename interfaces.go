package tcb

import (
	"github.com/couchbase/gocb"
)

// Configurable governs for configure and open connect
type CouchBaseAdapter interface {
	Open() error
	Cluster() *gocb.Cluster
	AddBucket(name string) (*gocb.Bucket, error)
	Bucket(name string) (*gocb.Bucket, error)
	Operator(name string) (Operatable, error)
}

// Operatable abstracted for operating on couchbase
type Operatable interface {
	Get(key string, data interface{}) (gocb.Cas, error)
	Insert(key string, data interface{}, expire uint32) (gocb.Cas, error)
	Upsert(key string, data interface{}, expire uint32) (gocb.Cas, error)
	Remove(key string) (gocb.Cas, error)
	N1qlQuery(q string, params interface{}) (gocb.QueryResults, error)
	N1qlQueryWithMode(mode *gocb.ConsistencyMode, q string, params interface{}) (gocb.QueryResults, error)
}

// Loggerable logging interface
type Loggerable interface {
	Log(...interface{})
	Logf(format string, v ...interface{})
}
