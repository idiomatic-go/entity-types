package accesslog

type NameAlias struct {
	Name  string
	Alias string
}

type AccessLogView struct {
	Version          string
	RequestHeaders   []NameAlias
	ResponseHeaders  []NameAlias
	ResponseTrailers []NameAlias
	Cookies          []NameAlias
}
