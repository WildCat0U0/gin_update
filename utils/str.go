package utils

import (
	"math/rand"
	"time"
)

//如果将 token 的有效期时间设置过短，到期后用户需要重新登录，过于繁琐且体验感差，
//这里我将采用服务端刷新 token 的方式来处理。先规定一个时间点，比如在过期前的 2 小时内
//，如果用户访问了接口，就颁发新的 token 给客户端（设置响应头），同时把旧 token 加入
//黑名单，在上一篇中，设置了一个黑名单宽限时间，目的就是避免并发请求中，刷新了 token
//，导致部分请求失败的情况；同时，我们也要避免并发请求导致 token 重复刷新的情况，这时候
//就需要上锁了，这里使用了 Redis 来实现，考虑到以后项目中可能会频繁使用锁，在篇头将简
//单做个封装

//用于生成锁标识，防止所有的客户端都能随意解锁

// RandString 这个代码定义了一个名为 RandString 的函数，它的作用是生成指定长度（由参数 len 指定）的随机字符串，其中每个字符都是大写字母（A 到 Z 中的一个）。
// 具体地，它使用了 Go 语言标准库中的 rand 包和 time 包，通过调用 rand.NewSource 和 rand.New 函数创建一个随机数生成器 r，然后在循环中生成指定长度的随机字符，并将它们存放在一个 []byte 类型的切片中，最后将该切片转换为字符串并返回。
func RandString(len int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}
