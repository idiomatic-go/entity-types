package accesslog

const LocalVersion = "0.0.0"

type Attributes struct {
	App              []string
	RequestHeaders   []string
	ResponseHeaders  []string
	ResponseTrailers []string
	Cookies          []string
}

type View struct {
	Version string
	Ingress Attributes
	Egress  Attributes
}

type CSVAttributes struct {
	App              string
	RequestHeaders   string
	ResponseHeaders  string
	ResponseTrailers string
	Cookies          string
}
