package api

import (
	"github.com/gin-gonic/gin"
	"go-blog/global"
	"go-blog/internal/service"
	"go-blog/pkg/app"
	"go-blog/pkg/convert"
	"go-blog/pkg/errcode"
	"go-blog/pkg/upload"
)

type Upload struct {

}

func NewUpload() Upload {
	return Upload{}
}

func (u Upload) UploadFile(c *gin.Context) {
	response := app.NewResponse(c)

	// 获取上传文件信息
	file, fileHeader, err := c.Request.FormFile("file")

	// 获得上传文件信息失败
	if err != nil {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Error()))
		return
	}

	fileType := convert.StrTo(c.PostForm("type")).MustInt()

	// 文件类型有误
	if fileHeader == nil || fileType <= 0 {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}

	svc := service.New(c.Request.Context())
	// 处理文件保存时报错
	fileInfo, err := svc.UploadFile(upload.FileType(fileType), file, fileHeader)
	if err != nil {
		global.Logger.Errorf("svc.UploadFile err: %v", err)
		response.ToErrorResponse(errcode.ErrorUploadFileFail.WithDetails(err.Error()))
		return
	}

	// 将文件下载链接返回给客户端
	response.ToResponse(gin.H{
		"file_access_url": fileInfo.AccessUrl,
	})
}
