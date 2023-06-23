package bootstrap

import (
	"Gin_Start/global"
	"Gin_Start/routers"
	"context"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// setupRouter 初始化路由
func setupRouter() *gin.Engine {
	r := gin.Default() // 创建默认路由
	con := cors.DefaultConfig()
	con.AllowAllOrigins = true
	r.Use(cors.New(con))
	//前端项目静态文件
	//Static函数的作用是 将指定的文件夹下的文件暴露出来，供外部访问
	//StaticFile函数的作用是 将指定的文件暴露出来，供外部访问
	r.StaticFile("/", "./static/dist/index.html")
	r.Static("/assets", "./static/dist/assets")
	r.StaticFile("/favicon.ico", "./static/dist/favicon.ico")
	//其他静态资源
	r.Static("/public", "./static")
	r.Static("/storage", "./storage/app/public")

	api := r.Group("/api")          // 创建 api 路由组
	routers.SetApiGroupRouters(api) // 注册 api 路由
	return r
}

// RunServer 启动服务
func RunServer() {
	//r := setupRouter()                      // 初始化路由
	//r.Run(":" + global.App.Config.App.Port) // 启动服务
	r := setupRouter() // 初始化路由

	srv := &http.Server{ // 创建 http.Server
		Addr:    ":" + global.App.Config.App.Port,
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			//global.App.Log.Error("listen: %s\n", zap.Any("err",err))
			log.Fatalf("listen: %s\n", err)
		}
	}()
	//等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	log.Println("Shutdown Server ...")
	//context的内容是 一个请求的上下文，包含请求的截止时间、请求的绑定信息等
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
