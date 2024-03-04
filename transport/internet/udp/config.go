package udp

import (
	"github.com/toxeh/xray-core/common"
	"github.com/toxeh/xray-core/transport/internet"
)

func init() {
	common.Must(internet.RegisterProtocolConfigCreator(protocolName, func() interface{} {
		return new(Config)
	}))
}
