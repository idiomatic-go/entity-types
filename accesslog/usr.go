package accesslog

const LocalVersion = "0.0.0"

type Attributes struct {
	App              []string
	Custom           []string
	RequestHeaders   []string
	ResponseHeaders  []string
	ResponseTrailers []string
	Cookies          []string
}

type Configuration struct {
	Version string
	Ingress Attributes
	Egress  Attributes
}

type CSVAttributes struct {
	App              string
	Custom           string
	RequestHeaders   string
	ResponseHeaders  string
	ResponseTrailers string
	Cookies          string
}
