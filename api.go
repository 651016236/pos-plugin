package pos_plugin

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

var EventService = EventApi{}

type EventApi struct{}

func (*EventApi) CallEvent(r *ghttp.Request) {
	d := r.GetBodyString()
	ret := CallEvent(d)
	_ = r.Response.WriteJsonExit(gconv.String(ret))
}
