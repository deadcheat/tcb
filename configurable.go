package tcb

import (
	"errors"

	"github.com/couchbase/gocb"
)

const (
	ErrBucketMissing error = errors.New("bucket is missing")
)

// ClusterConfig to connect couchbase
type ClusterConfig struct {
	ConnectString string
	User          string
	Password      string
	BucketConfigs []BucketConfig
	Loggerable
	cluster   *gocb.Cluster
	bucketMap map[string]*gocb.Bucket
}

// BucketConfig tuple for bucket connection
type BucketConfig struct {
	Name     string
	Password string
}

// NewClusterConfig return new instance
func NewClusterConfig() *ClusterConfig {
	return &ClusterConfig{Loggerable: NewDefaultLogger(true)}
}

// Open call this to open cluster connection
func (a *ClusterConfig) Open() error {
	cluster, err := gocb.Connect(a.ConnectString)
	if err == gocb.ErrAuthError {
		err = cluster.Authenticate(gocb.PasswordAuthenticator{
			Username: a.User,
			Password: a.Password,
		})
	}
	if err != nil {
		return err
	}
	a.cluster = cluster
	return nil
}

// Cluster return inner cluster instance
func (a *ClusterConfig) Cluster() *gocb.Cluster {
	return a.cluster
}

// AddBucket add a bucket to
func (a *ClusterConfig) AddBucket(bucket, password string) (*gocb.Bucket, error) {
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
func (a *ClusterConfig) Bucket(bucket string) *gocb.Bucket {
	b, _ := a.bucketMap[bucket]
	return b
}

// Operator return operator instance
func (a *ClusterConfig) Operator(bucket string) (Operatable, error) {
	b := a.Bucket(bucket)
	if b == nil {
		return nil, ErrBucketMissing
	}
	return NewBucketOperator(b, a.Loggerable), nil
}
