package pos_plugin

import (
	"encoding/json"
	"github.com/gogf/gf/crypto/gaes"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/encoding/gbase64"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gproc"
	"github.com/gogf/gf/os/gtimer"
	"github.com/gogf/gf/util/gconv"
	"time"
)

var PluginId string
var PluginPort int
var MasterData MasterInfo
var MasterDataStr string

// MasterPort 主程序端口号
var MasterPort int
var MasterHost string

// Init 		初始化插件
// pluginId 	插件ID
// masterPort 	主程序端口号
// pluginPort 	插件端口号
// s 		  	插件server
func Init(pluginId string, masterPort, pluginPort int, s *ghttp.Server, mh ...string) {
	// 主程序端口
	MasterPort = masterPort
	// 初始化插件Id
	PluginId = pluginId
	// 初始化插件端口
	PluginPort = pluginPort
	MasterHost = "localhost"
	if len(mh) > 0 {
		MasterHost = mh[0]
	}

	ppid := gproc.PPid()
	// 先设置为0防止启动讲主程序杀死
	//g.Log().Info("pid=", gproc.PPid(), ",isChild=", gproc.IsChild())
	gproc.SetPPid(0)
	//g.Log().Info("pid=", gproc.PPid(), ",isChild=", gproc.IsChild())
	gtimer.AddOnce(3*time.Second, func() {
		gproc.SetPPid(ppid)
	})

	// 获取主程序数据
	_ = getMasterInfo()

	// 初始化插件路由
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.POST("/event", EventService.CallEvent)
	})
}

// getMasterInfo 获取主程序配置项
func getMasterInfo() error {
	if MasterPort > 0 {
		if ret, err := g.Client().Timeout(30*time.Second).Get("http://"+MasterHost+":"+gconv.String(MasterPort)+"/master/info", `{"sign":"1111"}`); err != nil {
			return err
		} else {
			MasterDataStr = ret.ReadAllString()
			loadJson, _ := gjson.LoadJson(MasterDataStr)
			data := loadJson.GetString("data")
			body := AesDecode(data)
			MasterData = MasterInfo{}
			_ = json.Unmarshal([]byte(body), &MasterData)

			gdb.SetConfig(gdb.Config{
				"default": gdb.ConfigGroup{
					gdb.ConfigNode{
						Host:             MasterData.DbConfig.Host,
						Port:             MasterData.DbConfig.Port,
						User:             MasterData.DbConfig.User,
						Pass:             MasterData.DbConfig.Pass,
						Name:             MasterData.DbConfig.Name,
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
