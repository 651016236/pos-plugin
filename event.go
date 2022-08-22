package pos_plugin

import (
	"encoding/json"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"github.com/gookit/event"
)

var (
	eventList = make([]EventInfo, 0)
)

// CallEvent 调用事件
func CallEvent(data string) interface{} {
	eventKV := EventKV{}
	_ = json.Unmarshal([]byte(data), &eventKV)
	// 调用事件
	e := event.MustFire(eventKV.EventInfo.Event, event.M{"key": eventKV.Key, "data": eventKV.Data})
	return e.Get("callback")
}

// SetCallback 处理完以后设置事件回调
// e 当前事件
// data 需要返回的data数据
// extData 扩展数据
// over 是否覆盖 false
// stop 调用事件后是否return
func SetCallback(e event.Event, data interface{}, extData interface{}, over bool, stop ...bool) {
	_stop := false
	if len(stop) > 0 {
		_stop = stop[0]
	}

	marshal, err := json.Marshal(EventCallback{Stop: _stop, Data: data, ExtData: extData, Over: over})
	if err != nil {
		e.Set("callback", "")
		return
	}
	e.Set("callback", string(marshal))
}

func GetCallbackStruct(data string) EventCallback {
	eve := EventCallback{}
	_ = json.Unmarshal([]byte(data), &eve)
	return eve
}

// EventOn 监听事件
// name 事件名称
// evt 事件函数
// sync 是否为异步调用，默认为true
func EventOn(name string, evt func(e event.Event) error, sync ...bool) {
	e := EventInfo{
		PluginId: PluginId,
		Event:    name,
		Sync:     true,
		Port:     PluginPort,
	}
	if len(sync) > 0 {
		e.Sync = sync[0]
	}
	eventList = append(eventList, e)
	event.On(name, event.ListenerFunc(evt), event.Normal)
}

// EventRegisterMaster 注册事件到mater主程序
func EventRegisterMaster() {
	if MasterPort > 0 {
		jsonStr, _ := json.Marshal(eventList)
		g.Log().Info(string(jsonStr))
		if ret, err := g.Client().Post("http://"+MasterHost+":"+gconv.String(MasterPort)+"/plugin_new/event/register", string(jsonStr)); err != nil {
			g.Log().Info("获取主程序配置错误", err)
		} else {
			g.Log().Info("设置主程序数据库配置成功", ret.ReadAllString())
		}
	}
}
