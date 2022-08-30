package pos_plugin

type MasterInfo struct {
	DbConfig DbConfig `json:"db_config"`
	DataInfo DataInfo `json:"data_info"`
}

// 公共数据信息
type DataInfo struct {
	BusinessDay string `json:"business_day"`  // 营业日 Y-m-d
	PosUserName string `json:"pos_user_name"` // 收银员名称
	TerminalNo  string `json:"terminal_no"`   // 机器编号
	BranchNo    string `json:"branch_no"`     // 门店编号
	BranchName  string `json:"branch_name"`   // 门店名称
	PageSize    int    `json:"page_size"`     // 纸张大小
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
	Port     int    `json:"port"`
}

// DbConfig 数据库配置结构
type DbConfig struct {
	Host string `json:"host"`
	Port string `json:"port"`
	User string `json:"user"`
	Pass string `json:"pass"`
	Name string `json:"name"`
}

// EventCallback 事件回调结构体
type EventCallback struct {
	Stop    bool        `json:"stop"`     // 回调后直接return，不调用后面的方法
	Data    interface{} `json:"data"`     // 回调后的数据
	ExtData interface{} `json:"ext_data"` // 回调后的扩展数据
	Over    bool        `json:"over"`     // 是否覆盖数据
}
