package cloudhub

import (
	"github.com/kubeedge/beehive/pkg/core"
	"github.com/kubeedge/beehive/pkg/core/context"
	"github.com/kubeedge/kubeedge/cloud/edgecontroller/pkg/cloudhub/channelq"
	"github.com/kubeedge/kubeedge/cloud/edgecontroller/pkg/cloudhub/common/util"
	"github.com/kubeedge/kubeedge/cloud/edgecontroller/pkg/cloudhub/wsserver"
)

type cloudHub struct {
	context *context.Context
}

func init() {
	core.Register(&cloudHub{})
}

func (a *cloudHub) Name() string {
	return "cloudhub"
}

func (a *cloudHub) Group() string {
	return "cloudhub"
}

func (a *cloudHub) Start(c *context.Context) {
	a.context = c
	eventq, _ := channelq.NewChannelEventQueue(c)
	// start the cloudhub server
	wsserver.StartCloudHubNonTLS(util.HubConfig, eventq)
	wsserver.EventHandler.Context = c
	stopchan := make(chan bool)
	<-stopchan
}

func (a *cloudHub) Cleanup() {
	a.context.Cleanup(a.Name())
}
