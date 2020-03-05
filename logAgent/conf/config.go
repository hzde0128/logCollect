package conf

type AppConf struct {
	CenterConf `ini:"center"`
	KafkaConf `ini:"kafka"`
	EtcdConf  `ini:"etcd"`
}

type CenterConf struct {
	Address string `ini:"address"`
}

type KafkaConf struct {
	Address     string `ini:"address"`
	ChanMaxSize int    `ini:"chan_max_size"`
}

type EtcdConf struct {
	Address string `ini:"address"`
	Timeout int    `ini:"timeout"`
	Key     string `ini:"collect_log_key"`
}
