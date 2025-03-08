package dal

import (
	"pastebin/biz/dal/mysql"
	"pastebin/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
