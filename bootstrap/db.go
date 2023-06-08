package bootstrap

import (
	"Gin_Start/app/models"
	"Gin_Start/global"
	"go.uber.org/zap"
	"gopkg.in/natefinch/lumberjack.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

// gorm 有一个默认的 logger ，由于日志内容是输出到控制台的，我们需要自定义一个写入器，将默认logger.Writer
//接口的实现切换为自定义的写入器，上一篇引入了 lumberjack ，将继续使用它
// Path: bootstrap\log.go
// Compare this snippet from bootstrap\log.go:

// 自定义 gorm Writer
func getGormLogWriter() logger.Writer {
	var writer io.Writer

	// 是否启用日志文件
	if global.App.Config.Database.EnableFileLogWriter {
		// 自定义 Writer
		writer = &lumberjack.Logger{
			Filename:   global.App.Config.Log.RootDir + "/" + global.App.Config.Database.LogFilename, // 日志文件路径
			MaxSize:    global.App.Config.Log.MaxSize,                                                // 每个日志文件保存的最大尺寸 单位：M
			MaxBackups: global.App.Config.Log.MaxBackups,                                             // 日志文件最多保存多少个备份
			MaxAge:     global.App.Config.Log.MaxAge,                                                 // 文件最多保存多少天
			Compress:   global.App.Config.Log.Compress,                                               // 是否压缩
		}
	} else {
		// 默认 Writer
		writer = os.Stdout
	}
	return log.New(writer, "\r\n", log.LstdFlags) // 将 Writer 转换为 logger.Writer
}

// getGormLogger 自定义 gorm Logger
func getGormLogger() logger.Interface {
	var logMode logger.LogLevel

	switch global.App.Config.Database.LogMode {
	case "silent":
		logMode = logger.Silent
	case "error":
		logMode = logger.Error
	case "warn":
		logMode = logger.Warn
	case "info":
		logMode = logger.Info
	default:
		logMode = logger.Info
	}
	return logger.New(getGormLogWriter(), logger.Config{
		SlowThreshold:             200 * time.Millisecond,                          // 慢 SQL 阈值
		LogLevel:                  logMode,                                         // 日志级别
		Colorful:                  !global.App.Config.Database.EnableFileLogWriter, //彩色打印
		IgnoreRecordNotFoundError: false,                                           // 忽略ErrRecordNotFound（记录未找到）错误
	})
	//这个函数的作用是： 通过配置文件中的 logMode 字段，设置 gorm 的日志级别，以及是否开启彩色打印
	// gorm 的日志级别有 5 种，分别是 silent、error、warn、info、debug，对应的数字分别是 0、1、2、3、4
	// gorm默认的日志级别是 info，即 2
}

// InitializeDB 初始化数据库
func InitializeDB() (*gorm.DB, error) {
	// 根据配置文件进行初始化
	switch global.App.Config.Database.Driver {
	case "mysql":
		return initMySqlGorm()
	default:
		return initMySqlGorm()
	}
}

// initMySqlGorm 初始化 mysql 数据库
func initMySqlGorm() (*gorm.DB, error) {
	dbConfig := global.App.Config.Database
	if dbConfig.Database == "" {
		return nil, nil
	}
	dsn := dbConfig.UserName + ":" + dbConfig.Password + "@tcp(" + dbConfig.Host + ":" + strconv.Itoa(dbConfig.Port) + ")/" +
		dbConfig.Database + "?charset=" + dbConfig.Charset + "&parseTime=True&loc=Local"
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时删除旧的索引，然后创建同名的新索引，MySQL 不支持
		DontSupportRenameColumn:   true,  // 用 change 修改列名，MySQL 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}

	if db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,            // 禁用外键约束以便于自动迁移
		Logger:                                   getGormLogger(), // 使用默认 logger，显示详细信息
		//什么是外键约束？外键约束是用来保证数据的完整性，外键约束可以保证数据的一致性，防止数据不一致
	}); err != nil {
		global.App.Log.Error("MySQL启动异常", zap.Any("err", err))
		//zap.any 用于将任意类型的字段添加到日志中
		return nil, err
	} else {
		sqlDB, _ := db.DB()                          // 获取通用数据库对象 sql.DB ，然后使用其提供的功能
		sqlDB.SetMaxIdleConns(dbConfig.MaxIdleConns) // 设置空闲连接池中连接的最大数量
		sqlDB.SetMaxOpenConns(dbConfig.MaxOpenConns) // 设置打开数据库连接的最大数量
		initMySqlTables(db)
		return db, nil
	}
}

// 数据库初始化

// initMySqlTables 函数用于初始化数据库表，这里我们使用了 gorm 的自动迁移功能，只需要在初始化数据库连接时调用
// AutoMigrate 方法，传入需要迁移的表对应的 model 即可，gorm 会自动检查表是否存在，不存在则创建，存在则检查字段是否有变化，有变化则更新
func initMySqlTables(db *gorm.DB) {

	err := db.AutoMigrate(
		models.User{},
	)
	if err != nil {
		global.App.Log.Error("migrate table failed", zap.Any("err", err))
		os.Exit(0)
	}
}
