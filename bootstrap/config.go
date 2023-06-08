package bootstrap

import (
	"Gin_Start/global"
	"fmt"
	"github.com/fsnotify/fsnotify"
	Viper "github.com/spf13/viper"
	"os"
)

func InitializeConfig() *Viper.Viper {
	//设置配置文件路径
	config := "config.yaml"
	//生产环境可以通过设置环境变量来改变配置文件的路径
	if configEnv := os.Getenv("VIPER_CONFIG"); configEnv != "" {
		config = configEnv // 优先级更高
	}
	//初始化Viper
	v := Viper.New()
	//设置配置文件的名字
	v.SetConfigFile(config)
	//设置配置文件类型
	v.SetConfigType("yaml")
	//err 处理的是什么错误 读取配置文件的错误
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchRemoteConfig() // 监控远程配置文件变化
	//fsnotify.Event 事件
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("Config file changed:", in.Name) // 配置文件发生变化
		//重载配置
		if err := v.ReadInConfig(); err != nil {
			fmt.Println(err)
		}
	})
	//将配置赋值给全局变量
	//unmarshal 将配置信息绑定到结构体上
	if err := v.Unmarshal(&global.App.Config); err != nil { // 将配置文件中的配置信息绑定到全局变量上
		fmt.Println(err) // 配置文件解析失败
	}
	return v
}
