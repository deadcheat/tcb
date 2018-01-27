package tcb

import (
	"errors"

	"github.com/couchbase/gocb"
)

var (
	ErrBucketMissing error = errors.New("bucket is missing")
)

// Configure struct for config
type Configure struct {
	ConnectString string
	User          string
	Password      string
	BucketConfigs []BucketConfig
}

// Cluster to connect couchbase
type Cluster struct {
	Configure
	Loggerable
	cluster   *gocb.Cluster
	bucketMap map[string]*gocb.Bucket
}

// BucketConfig tuple for bucket connection
type BucketConfig struct {
	Name     string
	Password string
}

// NewCluster return new instance
func NewCluster(c Configure) *Cluster {
	bm := make(map[string]*gocb.Bucket)
	return &Cluster{Configure: c, Loggerable: NewDefaultActiveLogger(), bucketMap: bm}
}

// Open call this to open cluster connection
func (a *Cluster) Open() error {
	cluster, err := gocb.Connect(a.ConnectString)
	if err != nil {
		return err
	}
	if a.User != "" {
		_ = cluster.Authenticate(gocb.PasswordAuthenticator{
			Username: a.User,
			Password: a.Password,
		})
	}
	a.cluster = cluster
	for _, b := range a.BucketConfigs {
		if _, err := a.AddBucket(b.Name, b.Password); err != nil {
			return err
		}
	}
	return nil
}

// Cluster return inner cluster instance
func (a *Cluster) Cluster() *gocb.Cluster {
	return a.cluster
}

// AddBucket add a bucket to
func (a *Cluster) AddBucket(bucket, password string) (*gocb.Bucket, error) {
	if b, ok := a.bucketMap[bucket]; ok {
		return b, nil
	}
	b, err := a.cluster.OpenBucket(bucket, password)
	if err != nil {
		return nil, err
	}
	a.bucketMap[bucket] = b
	return b, nil
}

// Bucket return from bucketmap
func (a *Cluster) Bucket(bucket string) *gocb.Bucket {
	b, _ := a.bucketMap[bucket]
	return b
}

// Operator return operator instance
func (a *Cluster) Operator(bucket string) (Operatable, error) {
	b := a.Bucket(bucket)
	if b == nil {
		return nil, ErrBucketMissing
	}

	return NewBucketOperator(b, a.Loggerable), nil
}
