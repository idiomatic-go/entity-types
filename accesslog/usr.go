package accesslog

const LocalVersion = "0.0.0"

type View struct {
	Version          string
	Ingress          []string
	Egress           []string
	RequestHeaders   []string
	ResponseHeaders  []string
	ResponseTrailers []string
	Cookies          []string
}
