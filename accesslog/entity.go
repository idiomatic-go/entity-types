package accesslog

import (
	"github.com/idiomatic-go/common-lib/util"
	"strings"
	"sync/atomic"
)

type versionedState struct {
	hash  uint32
	attrs View
}

type VersionedEntity struct {
	index int32
	state [2]versionedState
}

func CreateEntity(ingress string,
	egress string,
	requestHeaders string,
	responseHeaders string,
	responseTrailers string,
	cookies string) View {
	view := View{Version: LocalVersion}
	view.Ingress = parseView(ingress)
	view.Egress = parseView(egress)
	view.RequestHeaders = parseView(requestHeaders)
	view.ResponseHeaders = parseView(responseHeaders)
	view.ResponseTrailers = parseView(responseTrailers)
	view.Cookies = parseView(cookies)
	return view
}

func parseView(attrs string) []string {
	if attrs == "" {
		return nil
	}
	tokens := strings.Split(attrs, ",")
	if tokens == nil {
		return nil
	}
	var list []string
	for _, s := range tokens {
		list = append(list, s)
	}
	return list
}

func CreateVersionedEntity() *VersionedEntity {
	return &VersionedEntity{}
}

func (v *VersionedEntity) getState() *versionedState {
	return &v.state[atomic.LoadInt32(&v.index)]
}

func (v *VersionedEntity) IsEmpty() bool {
	return v.getState().hash == 0
}

func (v *VersionedEntity) IsNewVersion(version string) bool {
	return util.SimpleHash(version) != atomic.LoadUint32(&v.getState().hash)
}

func (v *VersionedEntity) GetEntity() View {
	return v.getState().attrs
}

func (v *VersionedEntity) SetEntity(attrs *View) {
	if attrs == nil {
		return
	}
	index := atomic.LoadInt32(&v.index)
	// Toggle the index
	if index == 0 {
		index = 1
	} else {
		index = 0
	}
	v.state[index].attrs = *attrs
	v.state[index].hash = util.SimpleHash(attrs.Version)
	atomic.StoreInt32(&v.index, index)
}
