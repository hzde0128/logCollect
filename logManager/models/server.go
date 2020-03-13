package models

type Server struct {
	Id       int
	Hostname string `orm:"unique"` // 主机名称
	Address  string // 主机地址
	Collect   []*Collect `orm:"reverse(many)"`
}
