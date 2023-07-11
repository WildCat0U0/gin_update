package bootstrap

import (
	"Gin_Start/global"
	"github.com/jassue/go-storage/local"
)

func InitializeStorage() {
	_, _ = local.Init(global.App.Config.Storage.Disks.Local)
	//_, _ = kodo.Init(global.App.Config.Storage.Disks.QiNiu)
	//_, _ = oss.Init(global.App.Config.Storage.Disks.AliOss)
}
