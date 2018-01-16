package tcb

// Configconfiguration to connect couchbase
type Config struct {
	ConnectString string
	User          string
	Password      string
	BucketConfigs []BucketConfig
}

// BucketConfig tuple for bucket connection
type BucketConfig struct {
	Name     string
	Password string
}
