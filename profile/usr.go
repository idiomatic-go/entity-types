package profile

// Profile configuration for Go services. See : https://pkg.go.dev/net/http/pprof
type Configuration struct {
	Version    string
	Uri        string
	Names      []string // Profile names
	CPUSeconds int
}

type VersionedEntity struct {
}
