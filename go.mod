module logCollect

go 1.14

require (
	github.com/Shopify/sarama v1.26.1
	github.com/astaxie/beego v1.12.3
	github.com/coreos/bbolt v1.3.6 // indirect
	github.com/coreos/etcd v3.3.25+incompatible
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/go-systemd v0.0.0-20191104093116-d3cd4ed1dbcf // indirect
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/fastly/go-utils v0.0.0-20180712184237-d95a45783239 // indirect
	github.com/go-sql-driver/mysql v1.6.0
	github.com/golang/groupcache v0.0.0-20200121045136-8c9f03a8e57e // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/btree v1.1.2 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0 // indirect
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.16.0 // indirect
	github.com/hpcloud/tail v1.0.0
	github.com/jehiah/go-strftime v0.0.0-20171201141054-1d33003b3869 // indirect
	github.com/jonboulle/clockwork v0.3.0 // indirect
	github.com/lestrrat-go/file-rotatelogs v2.3.0+incompatible
	github.com/lestrrat-go/strftime v1.0.1 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/olivere/elastic/v7 v7.0.22
	github.com/rifflock/lfshook v0.0.0-20180920164130-b9218ef580f5
	github.com/sirupsen/logrus v1.8.1
	github.com/smartystreets/goconvey v1.6.4 // indirect
	github.com/soheilhy/cmux v0.1.5 // indirect
	github.com/tebeka/strftime v0.1.5 // indirect
	github.com/tmc/grpc-websocket-proxy v0.0.0-20220101234140-673ab2c3ae75 // indirect
	github.com/xiang90/probing v0.0.0-20221125231312-a49e3df8f510 // indirect
	go.etcd.io/bbolt v1.3.6 // indirect
	go.uber.org/zap v1.14.0 // indirect
	golang.org/x/lint v0.0.0-20210508222113-6edffad5e616 // indirect
	golang.org/x/sys v0.1.0 // indirect
	golang.org/x/time v0.3.0 // indirect
	golang.org/x/tools v0.1.4 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
	gopkg.in/ini.v1 v1.52.0
	sigs.k8s.io/yaml v1.3.0 // indirect
)

replace (
	github.com/coreos/bbolt v1.3.6 => go.etcd.io/bbolt v1.3.6
	go.etcd.io/bbolt v1.3.6 => github.com/coreos/bbolt v1.3.6
)
