package accesslog

import (
	"github.com/idiomatic-go/common-lib/util"
	"sync/atomic"
)

type versionedState struct {
	hash uint32
	view AccessLogView
}

type VersionedEntity struct {
	index int32
	state [2]versionedState
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

func (v *VersionedEntity) GetEntity() AccessLogView {
	return v.getState().view
}

func (v *VersionedEntity) SetEntity(view *AccessLogView) {
	if view == nil {
		return
	}
	index := atomic.LoadInt32(&v.index)
	// Toggle the index
	if index == 0 {
		index = 1
	} else {
		index = 0
	}
	v.state[index].view = *view
	v.state[index].hash = util.SimpleHash(view.Version)
	atomic.StoreInt32(&v.index, index)
}
