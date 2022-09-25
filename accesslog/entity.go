package accesslog

import (
	"github.com/idiomatic-go/common-lib/util"
	"strings"
	"sync/atomic"
)

type versionedState struct {
	hash uint32
	mu   Mutations
}

type VersionedEntity struct {
	index int32
	state [2]versionedState
}

func CreateEntity(ingressAttributes string,
	egressAttributes string,
	requestHeaders string,
	responseHeaders string,
	responseTrailers string,
	cookies string) Mutations {
	view := Mutations{Version: LocalVersion}
	view.IngressAttributes = parseAttributes(ingressAttributes)
	view.EgressAttributes = parseAttributes(egressAttributes)
	view.RequestHeaders = parseAttributes(requestHeaders)
	view.ResponseHeaders = parseAttributes(responseHeaders)
	view.ResponseTrailers = parseAttributes(responseTrailers)
	view.Cookies = parseAttributes(cookies)
	return view
}

func parseAttributes(attrs string) []string {
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

func (v *VersionedEntity) GetEntity() Mutations {
	return v.getState().mu
}

func (v *VersionedEntity) SetEntity(mu *Mutations) {
	if mu == nil {
		return
	}
	index := atomic.LoadInt32(&v.index)
	// Toggle the index
	if index == 0 {
		index = 1
	} else {
		index = 0
	}
	v.state[index].mu = *mu
	v.state[index].hash = util.SimpleHash(mu.Version)
	atomic.StoreInt32(&v.index, index)
}
