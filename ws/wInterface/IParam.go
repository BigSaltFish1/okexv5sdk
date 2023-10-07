package wInterface

import . "github.com/BigSaltFish1/okexv5sdk/ws/wImpl"

// 请求数据
type WSParam interface {
	EventType() Event
	ToMap() *map[string]string
}
