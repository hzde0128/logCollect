package conf

type AppConf struct {
	KafkaConf `ini:"kafka"`
	TailConf  `ini:"tail"`
}

type KafkaConf struct {
	Address string `ini:"address"`
	Topic   string `ini:"topic"`
}

type TailConf struct {
	Filename string `ini:"filename"`
}
