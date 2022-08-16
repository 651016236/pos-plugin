package pos_plugin

import (
	"encoding/json"
	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"github.com/gookit/event"
)

var (
	instances = gmap.NewStrAnyMap(true)
	eventList = make([]EventInfo, 0)
)

func CallEvent(data string) interface{} {
	eventKV := EventKV{}
	_ = json.Unmarshal([]byte(data), &eventKV)
	event.MustFire(data, event.M{"key": eventKV.Key, "data": eventKV.Data})
	return GetData(eventKV.Key)
}

func SetData(key string, data interface{}) {
	instances.Set(key, data)
}

func GetData(key string) interface{} {
	return instances.Get(key)
}

func EventOn(name string, evt func(e event.Event) error, sync ...bool) {
	e := EventInfo{
		PluginId: "",
		Event:    name,
		Sync:     false,
	}
	if len(sync) > 0 {
		if sync[0] {
			e.Sync = true
		}
	}
	eventList = append(eventList, e)
	event.On(name, event.ListenerFunc(evt), event.Normal)
}

func EventRegisterMaster() {
	if MasterPort > 0 {
		if ret, err := g.Client().Post("http://localhost:"+gconv.String(MasterPort)+"/event/register", eventList); err != nil {
			g.Log().Info("获取主程序配置错误", err)
		} else {
			g.Log().Info("设置主程序数据库配置成功", ret.ReadAllString())
		}
	}
}
