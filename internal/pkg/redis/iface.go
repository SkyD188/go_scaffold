package redis

import "time"

type Iface interface {
	Set(key, val string, time time.Duration)
	Get(key string) string

	// 补充
	HSet(key, field string, value interface{})
	HGet()
	LPush()
	RPush()
	Eval()
}
