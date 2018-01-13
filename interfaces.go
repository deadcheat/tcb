package tcb

import (
	"github.com/couchbase/gocb"
)

// CouchBaseAdaptor CouchBase connect adaptor
type CouchBaseAdaptor interface {
	Open() error
	Env() *Config
	Cluster() *gocb.Cluster
	Bucket() *gocb.Bucket
	Get(key string, data interface{}) (gocb.Cas, error)
	Insert(key string, data interface{}, expire uint32) (gocb.Cas, error)
	Upsert(key string, data interface{}, expire uint32) (gocb.Cas, error)
	Remove(key string) (gocb.Cas, error)
	N1qlQuery(q string, params interface{}) (gocb.QueryResults, error)
	N1qlQueryWithMode(mode *gocb.ConsistencyMode, q string, params interface{}) (gocb.QueryResults, error)
}

// Configurable governs for configure and open connect
type Configurable interface {
	Open() error
	Env() *Config
}

// Loggerable logging interface
type Loggerable interface {
	Log(...interface{})
	Logf(format string, v ...interface{})
}
