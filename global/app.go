package global

import (
	"Gin_Start/config"
	"github.com/go-redis/redis/v8"
	"github.com/jassue/go-storage/storage"
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

func (app *Application) Disk(disk ...string) storage.Storage {
	diskName := app.Config.Storage.Default
	if len(disk) > 0 {
		diskName = storage.DiskName(disk[0])
	}
	s, err := storage.Disk(diskName)
	if err != nil {
		panic(err)
	}
	return s
}
