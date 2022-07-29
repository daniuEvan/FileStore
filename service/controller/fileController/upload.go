/**
 * @date: 2022/4/11
 * @desc: ...
 */

package fileController

import (
	"FileStore/common/response"
	"FileStore/global"
	"FileStore/service/model/fileModel"
	"FileStore/utils"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io/ioutil"
	"os"
	"time"
)

func UploadFileHandler(ctx *gin.Context) {
	// 单个文件
	file, err := ctx.FormFile("file")
	if err != nil {
		global.Logger.Error("文件上传失败", zap.String("err:", err.Error()))
		response.Failed(ctx, nil, err.Error())
		return
	}
	tempFile, err := file.Open()
	defer tempFile.Close()
	if err != nil {
		global.Logger.Error("文件上传失败", zap.String("err:", err.Error()))
		response.Failed(ctx, nil, err.Error())
		return
	}

	filePath := fmt.Sprintf("uploadFiles/%s", file.Filename)
	newFile, err := os.Create(filePath)
	newFile.Seek(0, 0)
	if err != nil {
		if err != nil {
			global.Logger.Error("文件上传失败", zap.String("err:", err.Error()))
			response.Failed(ctx, nil, err.Error())
			return
		}
	}
	fileMeta := fileModel.FileMeta{
		FileSha1: utils.FileSha1(newFile),
		FileName: file.Filename,
		FileSize: file.Size,
		Location: filePath,
		UploadAt: time.Now().Format("2006-01-02 15:04:05"),
	}
	fileModel.UpdateFileMeta(fileMeta)

	// 上传文件到指定的目录
	err = ctx.SaveUploadedFile(file, filePath)
	if err != nil {
		global.Logger.Error("文件上传失败", zap.String("err:", err.Error()))
		response.Failed(ctx, nil, err.Error())
		return
	}
	response.Success(ctx, fileMeta, "上传成功")
}

func GetFileMetaHandler(ctx *gin.Context) {
	fileHash := ctx.Query("filehash")
	fileMeta := fileModel.GetFileMeta(fileHash)
	_, err := json.Marshal(fileMeta)
	if err != nil {
		global.Logger.Error("获取文件元信息失败", zap.String("err:", err.Error()))
		response.Failed(ctx, nil, err.Error())
		return
	}
	response.Success(ctx, fileMeta, "获取成功")
}

func DownloadHandler(ctx *gin.Context) {
	fileHash := ctx.Query("filehash")
	fileMeta := fileModel.GetFileMeta(fileHash)
	file, err := os.Open(fileMeta.Location)
	if err != nil {
		global.Logger.Error("文件下载失败", zap.String("err:", err.Error()))
		response.Failed(ctx, nil, err.Error())
		return
	}
	defer file.Close()
	_, err = ioutil.ReadAll(file)
	if err != nil {
		global.Logger.Error("文件下载失败", zap.String("err:", err.Error()))
		response.Failed(ctx, nil, err.Error())
		return
	}
	ctx.Header("Content-Type", "application/octet-stream")
	ctx.Header("Content-Disposition", "attachment; filename="+fileMeta.FileName)
	ctx.Header("Content-Transfer-Encoding", "binary")
	ctx.File(fileMeta.Location)
}

//
// UpdateHandler
// @Description: 重命名
// @param ctx:
//
func UpdateHandler(ctx *gin.Context) {
	fileSha1 := ctx.Query("filehash")
	fileMeta := fileModel.GetFileMeta(fileSha1)
	newFileName := ctx.Query("filename")
	fileMeta.FileName = newFileName
	fileModel.UpdateFileMeta(fileMeta)
	//os.Rename(fileMeta.Location, fileMeta.Location)
	response.Success(ctx, fileMeta, "重命名成功")
}

//
// DeleteHandler
// @Description: 删除文件
// @param ctx:
//
func DeleteHandler(ctx *gin.Context) {
	fileSha1 := ctx.Query("filehash")
	fileMeta := fileModel.GetFileMeta(fileSha1)
	fileModel.RemoveFile(fileSha1)
	os.Remove(fileMeta.Location)
	response.Success(ctx, fileMeta, "删除文件成功")
}
