package plugin

type InstallYaml struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Path        string `json:"path"`
}

type MasterInfo struct {
	DbConfig DbConfig `json:"db_config"`
}

type EventKV struct {
	Key       string      `json:"key"`
	Data      interface{} `json:"data"`
	EventInfo EventInfo   `json:"event_info"`
}

type EventInfo struct {
	PluginId string `json:"plugin_id"`
	Event    string `json:"event"`
	Sync     bool   `json:"sync"`
}

// DbConfig 数据库配置结构
type DbConfig struct {
	Host string `json:"host"`
	Port string `json:"port"`
	User string `json:"user"`
	Pass string `json:"pass"`
	Name string `json:"name"`
}

// MasterPort 主程序端口号
var MasterPort int
