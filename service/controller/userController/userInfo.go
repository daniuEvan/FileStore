/**
 * @date: 2022/2/18
 * @desc: ...
 */

package userController

import (
	"FileStore/common/currentUser"
	"FileStore/common/formatError"
	"FileStore/common/response"
	"FileStore/common/smService"
	"FileStore/database"
	"FileStore/global"
	"FileStore/service/form/userForm"
	"FileStore/service/model/userModel"
	"FileStore/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

//
// EditUserInfo
// @Description: 编辑用户信息
// @param ctx:
//
func EditUserInfo(ctx *gin.Context) {
	db, err := database.GetDB(ctx)
	if err != nil {
		global.Logger.Error("编辑用户信息", zap.String("errorMsg", err.Error()))
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "服务异常")
		return
	}
	// 表单验证
	userInfoForm := userForm.UserInfoForm{}
	if err := ctx.ShouldBindJSON(&userInfoForm); err != nil {
		global.Logger.Error(err.Error())
		formatError.ValidatorErrorHandler(ctx, err)
		return
	}
	userId, ok := currentUser.GetCurrentUserID(ctx)
	if !ok {
		response.Response(ctx, http.StatusUnauthorized, 401, nil, "未登录")
		return
	}
	var user userModel.User
	db.Where("id = ?", userId).First(&user)
	if user.ID == 0 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户不存在")
		return
	}
	user.Username = userInfoForm.Username
	db.Save(&user)
	response.Success(ctx, gin.H{
		"id":       user.ID,
		"username": user.Username,
		"mobile":   user.Mobile,
	}, "更新成功")

}

//
// GetUserInfo
// @Description: 获取用户信息
// @param ctx:
//
func GetUserInfo(ctx *gin.Context) {
	db, err := database.GetDB(ctx)
	if err != nil {
		global.Logger.Error("获取用户信息", zap.String("errorMsg", err.Error()))
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "服务异常")
		return
	}
	userMobile := ctx.Param("mobile")
	var user userModel.User
	db.Where("mobile=?", userMobile).First(&user)
	if user.ID == 0 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户不存在")
		return
	}
	response.Success(ctx, gin.H{
		"id":       user.ID,
		"username": user.Username,
		"mobile":   user.Mobile,
	}, "")

}

//
// ChangePwd
// @Description: 修改密码
// @param ctx:
//
func ChangePwd(ctx *gin.Context) {
	db, err := database.GetDB(ctx)
	if err != nil {
		global.Logger.Error("修改密码", zap.String("errorMsg", err.Error()))
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "服务异常")
		return
	}
	// 表单验证
	userPwdInfoForm := userForm.UserPwdInfoForm{}
	if err := ctx.ShouldBindJSON(&userPwdInfoForm); err != nil {
		global.Logger.Error(err.Error())
		formatError.ValidatorErrorHandler(ctx, err)
		return
	}
	mobile := userPwdInfoForm.Mobile
	userId, ok := currentUser.GetCurrentUserID(ctx)
	if !ok {
		response.Response(ctx, http.StatusUnauthorized, 401, nil, "未登录或者登录异常")
		return
	}
	smCode := userPwdInfoForm.VerifyCode
	newPwd := userPwdInfoForm.Password
	var user userModel.User
	db.Where("id = ? and mobile = ?", userId, mobile).First(&user)
	if user.ID == 0 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户不存在")
		return
	}
	// 校验验证码
	sm := smService.NewSmService()
	ok, err = sm.VerifySmCode(mobile, smCode)
	if err != nil {
		global.Logger.Error("修改密码", zap.String("errorMsg", err.Error()))
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "验证码校验服务异常")
		return
	}
	if !ok {
		response.Response(ctx, http.StatusInternalServerError, 400, nil, "验证码错误")
		return
	}
	hashPassword, err := utils.GetHashPwd(newPwd)
	if err != nil {
		global.Logger.Error("修改密码", zap.String("errorMsg", err.Error()))
		response.Response(ctx, http.StatusUnprocessableEntity, 500, nil, "密码加密错误")
		return
	}
	user.Password = hashPassword
	db.Save(&user)
	response.Success(ctx, gin.H{}, "密码修改成功")
}

//
// GetUserList
// @Description: 获取用户列表
// @param ctx:
//
func GetUserList(ctx *gin.Context) {
	db, err := database.GetDB(ctx)
	if err != nil {
		global.Logger.Error("获取用户列表", zap.String("errorMsg", err.Error()))
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "服务异常")
		return
	}
	var userList []userModel.User
	result := db.Find(&userList)
	if result.Error != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "获取用户列表异常")
		return
	}
	for _, user := range userList {
		user.Password = ""
	}

	response.Response(ctx, http.StatusOK, 200, userList, "获取用户列表成功")
}
