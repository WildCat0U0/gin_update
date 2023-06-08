package global

import (
	"Gin_Start/config"
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
}

var App = new(Application)
