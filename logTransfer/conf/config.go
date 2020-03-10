package conf

type AppConf struct {
	KafkaConf `ini:"kafka"`
	EsConf    `ini:"es"`
	LogConf   `ini:"log"`
}

type KafkaConf struct {
	Address string `ini:"address"`
	Topic   string `ini:"topic"`
	Group   string `ini:"group"`
}

type EsConf struct {
	Address     string `ini:"address"`
	MaxChanSize int    `ini:"max_chan_size"`
	Nums        int    `ini:"nums"`
}

type LogConf struct {
	FilePath string `ini:"filePath"`
	FileName string `ini:"filename"`
	LogLevel string `ini:"loglevel"`
	MaxAge   int    `ini:"max_age"`
}
