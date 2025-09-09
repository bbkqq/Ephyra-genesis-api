package redis

import (
	"fmt"
	"time"
)

// UserLoginTokenKey 用户登录的时候用于保存token
func UserLoginTokenKey(uid, address interface{}) string {
	return fmt.Sprintf("user_login_token/%v_%v", uid, address)
}

func ChainEventListenKey(chainName string) string {
	return fmt.Sprintf("last_listen_on_block_number_%v", chainName)
}

// UserDailyPointsZSetKey 用户每日积分排名的ZSet key
func UserDailyPointsZSetKey(date time.Time) string {
	return fmt.Sprintf("user_daily_points_zset_%v", date.Format(time.DateOnly))
}

// UserWeeklyPointsZSetKey 用户周积分排名的ZSet key
func UserWeeklyPointsZSetKey(weekStart time.Time) string {
	return fmt.Sprintf("user_weekly_points_zset_%v", weekStart.Format(time.DateOnly))
}
