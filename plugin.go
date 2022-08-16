package pos_plugin

import (
	"encoding/json"
	"github.com/gogf/gf/crypto/gaes"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/encoding/gbase64"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

var PluginId string
var PluginPort int

func Init(pluginId string, masterPort, pluginPort int, s *ghttp.Server) {
	// 主程序端口
	MasterPort = masterPort
	// 初始化插件Id
	PluginId = pluginId
	// 初始化插件端口
	PluginPort = pluginPort
	// 获取主程序数据
	_ = getMasterInfo()

	// 初始化插件路由
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.POST("/event", EventService.CallEvent)
	})
}

// 获取主程序配置项
func getMasterInfo() error {
	if MasterPort > 0 {
		if ret, err := g.Client().Get("http://localhost:"+gconv.String(MasterPort)+"/master/info", `{"sign":"1111"}`); err != nil {
			return err
		} else {
			loadJson, _ := gjson.LoadJson(ret.ReadAllString())
			data := loadJson.GetString("data")
			body := AesDecode(data)
			masterInfo := MasterInfo{}
			_ = json.Unmarshal([]byte(body), &masterInfo)
			gdb.SetConfig(gdb.Config{
				"default": gdb.ConfigGroup{
					gdb.ConfigNode{
						Host:             masterInfo.DbConfig.Host,
						Port:             masterInfo.DbConfig.Port,
						User:             masterInfo.DbConfig.User,
						Pass:             masterInfo.DbConfig.Pass,
						Name:             masterInfo.DbConfig.Name,
						Type:             "mysql",
						Role:             "master",
						MaxOpenConnCount: 5,
						Weight:           100,
					},
				},
			})
		}
	}
	return nil
}

func AesDecode(body string) string {
	decode, err := gbase64.Decode([]byte(body))
	if err != nil {
		return ""
	}
	res, _ := gaes.Decrypt(decode, []byte("1234567890123456"))
	return string(res)
}
