package conf

// AppConf 主配置文件
type AppConf struct {
	CenterConf `ini:"center"`
	KafkaConf  `ini:"kafka"`
	EtcdConf   `ini:"etcd"`
	LogConf    `ini:"log"`
}

// CenterConf center节
type CenterConf struct {
	Address string `ini:"address"`
}

// KafkaConf kafka节
type KafkaConf struct {
	Address     string `ini:"address"`
	ChanMaxSize int    `ini:"chan_max_size"`
}

// EtcdConf etcd节
type EtcdConf struct {
	Address string `ini:"address"`
	Timeout int    `ini:"timeout"`
	Key     string `ini:"collect_log_key"`
}

// LogConf log节
type LogConf struct {
	FilePath string `ini:"filePath"`
	FileName string `ini:"filename"`
	LogLevel string `ini:"loglevel"`
	MaxAge   int    `ini:"max_age"`
}
