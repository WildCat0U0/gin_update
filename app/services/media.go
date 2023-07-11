package services

import (
	"Gin_Start/app/common/request"
	"Gin_Start/app/models"
	"Gin_Start/global"
	"context"
	"errors"
	"github.com/jassue/go-storage/storage"
	uuid "github.com/satori/go.uuid"
	"path"
	"strconv"
	"time"
)

type mediaService struct {
}

var MediaService = new(mediaService)

type outPut struct {
	Id   int64  `json:"id"`
	Path string `json:"path"`
	Url  string `json:"url"`
}

const mediaCacheKeyPre = "media:"

// makeFaceDir 文件存储目录
func (mediaService *mediaService) makeFaceDir(business string) string {
	return global.App.Config.App.Env + "/" + business
}

// HashName 生成文件名称，使用的是uuid
func (mediaService *mediaService) HashName(filename string) string {
	fileSuffix := path.Ext(filename)
	return uuid.NewV4().String() + fileSuffix
}

// SaveImage 保存图片（公共读
// 接收一个类型为request.ImageUpload的参数params，其中包含了图片的信息，包括文件名、文件大小、文件对象等。
//
// 打开图片文件并读取其中的内容，如果读取失败则返回一个错误信息。
//
// 根据全局变量global.App.Config.Storage.Default判断使用的是本地存储还是云存储，如果是本地存储，则将图片存储到本地文件系统中，并在图片的路径前添加"public/"前缀，以保证后续处理静态资源时能够正确访问到图片。
//
// 根据查询参数params中的业务信息，生成一个图片存储路径key，使用该key将图片存储到指定的存储介质中（本地或云存储）。
//
// 将图片的相关信息（包括存储介质类型、图片路径、图片大小等）保存到数据库中。
//
// 返回一个类型为outPut的结果，其中包括图片的相关信息，包括图片ID、图片路径和图片URL。
func (mediaService *mediaService) SaveImage(params request.ImageUpload) (result outPut, err error) {
	file, err := params.Image.Open()
	defer file.Close()
	if err != nil {
		err = errors.New("上传失败")
		return
	}
	localPrefix := ""
	//// 本地文件存放路径为 storage/app/public，由于在『（五）静态资源处理 & 优雅重启服务器』中，
	//    // 配置了静态资源处理路由 router.Static("/storage", "./storage/app/public")
	//    // 所以此处不需要将 public/ 存入到 mysql 中，防止后续拼接文件 Url 错误
	if global.App.Config.Storage.Default == storage.Local {
		localPrefix = "public" + "/"
	}
	key := mediaService.makeFaceDir(params.Business) + "/" + mediaService.HashName(params.Image.Filename)
	disk := global.App.Disk()
	err = disk.Put(localPrefix+key, file, params.Image.Size)
	if err != nil {
		return
	}
	image := models.Media{
		DiskType: string(global.App.Config.Storage.Default),
		SrcType:  1,
		Src:      key,
	}
	err = global.App.DB.Create(&image).Error
	if err != nil {
		return
	}
	result = outPut{
		Id:   int64(image.ID.ID),
		Path: key,
		Url:  disk.Url(key),
	}
	return
}

// GetUrlById 通过 id 获取文件 url
func (mediaService *mediaService) GetUrlById(id int64) string {
	if id == 0 {
		return ""
	}
	var url string
	cacheKey := mediaCacheKeyPre + strconv.FormatInt(id, 10)

	exist := global.App.Redis.Exists(context.Background(), cacheKey).Val()
	if exist == 1 {
		url = global.App.Redis.Get(context.Background(), cacheKey).Val()
	} else {
		media := models.Media{}
		err := global.App.DB.First(&media, id).Error
		if err != nil {
			return ""
		}
		url = global.App.Disk(media.DiskType).Url(media.Src)
		global.App.Redis.Set(context.Background(), cacheKey, url, time.Second*3*24*3600)
	}

	return url
}
