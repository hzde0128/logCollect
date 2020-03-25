package utils

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/astaxie/beego/config"
)

var (
	GServerName string // 应用程序名称
	GMysqlHost  string // mysql地址
	GMysqlPort  string // mysql端口
	GMysqlUser  string // mysql用户名
	GMysqlPass  string // mysql密码
	GMysqlDb    string // mysql数据库
	GRedisHost  string // redis地址
	GRedisPort  string // redis端口
	GRedisDb    string // redis库
)

func initConfig() {
	appconf, err := config.NewConfig("ini", "conf/app.conf")
	if err != nil {
		panic(err)
	}

	GServerName = appconf.String("appname")
	GMysqlHost = appconf.String("mysql_host")
	GMysqlPort = appconf.String("mysql_port")
	GMysqlUser = appconf.String("mysql_user")
	GMysqlPass = appconf.String("mysql_pass")
	GMysqlDb = appconf.String("mysql_db")
	GRedisHost = appconf.String("redis_host")
	GRedisPort = appconf.String("redis_port")
	GRedisDb = appconf.String("redis_db")

	return
}

// Md5String 加密函数
func Md5String(s string) string {
	w := md5.New()
	w.Write([]byte(s))
	return hex.EncodeToString(w.Sum(nil))
}

// RedisConn redis连接
func RedisConn() (bm cache.Cache, err error) {
	// 连接redis
	// 准备连接redis信息
	redisConf := map[string]string{
		"key":   GServerName,
		"conn":  GRedisHost + ":" + GRedisPort,
		"dbNum": GRedisDb,
	}

	// 将map转化为json
	redisConfJs, _ := json.Marshal(redisConf)

	// 创建redis句柄
	bm, err = cache.NewCache("redis", string(redisConfJs))
	if err != nil {
		return
	}
	return
}

func init() {
	initConfig()
}
