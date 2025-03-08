package dal

import (
	"pastebin/biz/dal/mysql"
)

func Init() {
	// redis.Init()
	mysql.Init()
}
