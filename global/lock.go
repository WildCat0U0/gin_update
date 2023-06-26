package global

import (
	"Gin_Start/utils"
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type Interface interface {
	Get() bool
	Block(seconds int64) bool
	Release() bool
	ForceRelease()
}

type lock struct {
	context context.Context
	name    string
	owner   string // 锁标识
	seconds int64  //有效期
}

func (l *lock) Get() bool {
	return App.Redis.SetNX(l.context, l.name, l.owner, time.Duration(l.seconds)*time.Second).Val()
}

// 阻塞一段时间，尝试获取锁

func (l *lock) Block(seconds int64) bool {
	starting := time.Now().Unix()
	for {
		if !l.Get() {
			time.Sleep(time.Duration(1) * time.Second)
			if time.Now().Unix()-seconds >= starting {
				return false
			}
		} else {
			return true
		}
	}
}

func (l *lock) Release() bool {
	luaScript := redis.NewScript(releaseLockLuaScript)
	result := luaScript.Run(l.context, App.Redis, []string{l.name}, l.owner).Val().(int64)
	return result != 0
}

// ForceRelease 强制释放锁
func (l *lock) ForceRelease() {
	App.Redis.Del(l.context, l.name).Val()
}

// 释放锁的lua脚本，防止所有客户端都能够解锁
const releaseLockLuaScript = `
if redis.call("get",KEY[1]) == ARGB[1] then
	return redis.call("del",KEY[1])
else
	return 0
end
`

func Lock(name string, seconds int64) Interface {
	return &lock{
		context: context.Background(),
		name:    name,
		owner:   utils.RandString(16),
		seconds: seconds,
	}
}
