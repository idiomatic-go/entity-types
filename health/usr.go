package health

// Configuration for application health endpoints
type Configuration struct {
	Version  string
	Uri      []string
	Timeout  int // seconds
	Interval int // seconds
	Content  string
	Protocol string // TCP,UDP
}

type VersionedEntity struct {
}
