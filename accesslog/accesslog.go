package accesslog

import (
	"github.com/idiomatic-go/common-lib/util"
)

type AccessLogView struct {
	Version string
	Headers []string
	Cookies []string
}

func CreateVersionedEntity() *util.VersionedEntity {
	return util.CreateVersionedEntity(AccessLogView{}, GetVersion)
}

var GetVersion = func(a any) string {
	e, ok := GetEntity(a)
	if !ok {
		return ""
	}
	return e.Version
}

var GetEntity = func(a any) (AccessLogView, bool) {
	if a == nil {
		return AccessLogView{}, false
	}
	if data, ok := a.(AccessLogView); ok {
		return data, ok
	}
	return AccessLogView{}, false
}
