package profile

// Profile configuration for Go services. See : https://pkg.go.dev/net/http/pprof
type Profile struct {
	Version    string
	Name       []string
	CPUSeconds int
}
