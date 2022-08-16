package plugin

import (
	"github.com/651016236/pos-plugin/plugin/ptools"
	"github.com/gogf/gf/net/ghttp"
)

var EventService = EventApi{}

type EventApi struct{}

func (*EventApi) CallEvent(r *ghttp.Request) {
	d := r.GetBodyString()
	ret := CallEvent(d)
	ptools.Success(r, ret)
}
