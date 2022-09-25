package accesslog

const LocalVersion = "0.0.0"

type Mutations struct {
	Version           string
	IngressAttributes []string
	EgressAttributes  []string
	RequestHeaders    []string
	ResponseHeaders   []string
	ResponseTrailers  []string
	Cookies           []string
}
