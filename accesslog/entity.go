package accesslog

import (
	"github.com/idiomatic-go/common-lib/util"
	"strings"
	"sync/atomic"
)

func CreateConfiguration(ingress *CSVAttributes, egress *CSVAttributes) Configuration {
	config := Configuration{Version: LocalVersion}
	config.Ingress = Attributes{App: tokenize(ingress.App), Custom: tokenize(ingress.Custom), RequestHeaders: tokenize(ingress.RequestHeaders), ResponseHeaders: tokenize(ingress.ResponseHeaders), ResponseTrailers: tokenize(ingress.ResponseTrailers), Cookies: tokenize(ingress.Cookies)}
	config.Egress = Attributes{App: tokenize(egress.App), Custom: tokenize(egress.Custom), RequestHeaders: tokenize(egress.RequestHeaders), ResponseHeaders: tokenize(egress.ResponseHeaders), ResponseTrailers: tokenize(egress.ResponseTrailers), Cookies: tokenize(egress.Cookies)}
	return config
}

func tokenize(attrs string) []string {
	if attrs == "" {
		return nil
	}
	tokens := strings.Split(attrs, ",")
	if tokens == nil {
		return nil
	}
	var list []string
	for _, s := range tokens {
		if s == "" {
			continue
		}
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

func (v *VersionedEntity) GetEntity() Configuration {
	return v.getState().config
}

func (v *VersionedEntity) SetEntity(config *Configuration) {
	if config == nil {
		return
	}
	index := atomic.LoadInt32(&v.index)
	// Toggle the index
	if index == 0 {
		index = 1
	} else {
		index = 0
	}
	v.state[index].config = *config
	v.state[index].hash = util.SimpleHash(config.Version)
	atomic.StoreInt32(&v.index, index)
}
