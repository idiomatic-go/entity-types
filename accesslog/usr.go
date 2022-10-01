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

type versionedState struct {
	hash   uint32
	config Configuration
}

type VersionedEntity struct {
	index int32
	state [2]versionedState
}

type CSVAttributes struct {
	App              string
	Custom           string
	RequestHeaders   string
	ResponseHeaders  string
	ResponseTrailers string
	Cookies          string
}
