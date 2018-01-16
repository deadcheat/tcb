package tcb

import (
	"github.com/couchbase/gocb"
)

// Adapter to connect couchbase
type Adapter struct {
	ConnectString string
	User          string
	Password      string
	BucketConfigs []BucketConfig
	cluster       *gocb.Cluster
	bucketMap     *map[string]*gocb.Bucket
}

// BucketConfig tuple for bucket connection
type BucketConfig struct {
	Name     string
	Password string
}

// Open call this to open cluster connection
func (a *Adapter) Open() error {
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

// Cluster get cluster instance
func (a *Adapter) Cluster() *gocb.Cluster {
	return a.cluster
}
