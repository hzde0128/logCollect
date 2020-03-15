package models

import "time"

// Collect 日志收集表
type Collect struct {
	Id         int
	Path       string    // 日志文件路径
	Topic      string    // 日志对应的Topic
	Server     *Server   `orm:"rel(fk)"`                     // 服务器信息
	CreateTime time.Time `orm:"type(datetime);auto_now_add"` // 添加时间
}
