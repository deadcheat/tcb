package tcb

// Configconfiguration to connect couchbase
type Config struct {
	User           string
	Password       string
	BucketName     string
	BucketPassword string
	ConnectString  string
}
