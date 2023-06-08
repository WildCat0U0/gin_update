package global

import (
	"Gin_Start/config"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"gorm.io/gorm"
)
import Viper "github.com/spf13/viper"

//中文

type Application struct {
	ConfigViper *Viper.Viper          // Viper
	Config      config.Configurations // Configurations
	Log         *zap.Logger
	DB          *gorm.DB
	Redis       *redis.Client
}

var App = new(Application)
