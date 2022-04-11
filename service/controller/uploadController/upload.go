/**
 * @date: 2022/4/11
 * @desc: ...
 */

package uploadController

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func UploadFile(ctx *gin.Context) {
	// 单个文件
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	log.Println(file.Filename)
	dst := fmt.Sprintf("uploadFiles/%s", file.Filename)
	// 上传文件到指定的目录
	err = ctx.SaveUploadedFile(file, dst)
	if err != nil {
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("'%s' uploaded!", file.Filename),
	})
}
