package dal

import (
	"Ephyra-genesis-api/biz/dal/mysql"
	"Ephyra-genesis-api/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
