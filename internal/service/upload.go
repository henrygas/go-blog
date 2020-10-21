package service

import (
	"errors"
	"go-blog/global"
	"go-blog/pkg/upload"
	"mime/multipart"
	"os"
)

type FileInfo struct {
	Name      string
	AccessUrl string
}

func (svc *Service) UploadFile(fileType upload.FileType, file multipart.File, fileHeader *multipart.FileHeader) (*FileInfo, error) {
	fileName := upload.GetFileName(fileHeader.Filename)
	uploadSavePath := upload.GetSavePath()
	dst := uploadSavePath + "/" + fileName

	// 检查上传文件扩展名是否合法
	if !upload.CheckContainExt(fileType, fileName) {
		return nil, errors.New("File suffix is not supported.")
	}

	// 创建保存目录
	if upload.CheckSavePath(uploadSavePath) {
		if err := upload.CreateSavePath(uploadSavePath, os.ModePerm); err != nil {
			return nil, errors.New("Failed to create save directory.")
		}
	}

	// 检查文件大小是否合法
	if upload.CheckMaxSize(fileType, file) {
		return nil, errors.New("Exceeded maximum file limit.")
	}

	// 检查文件权限是否合法
	if upload.CheckPermission(uploadSavePath) {
		return nil, errors.New("Insufficient file permissions.")
	}

	// 保存文件
	if err := upload.SaveFile(fileHeader, dst); err != nil {
		return nil, err
	}

	accessUrl := global.AppSetting.UploadServerUrl + "/" + fileName
	return &FileInfo{Name: fileName, AccessUrl: accessUrl}, nil
}
