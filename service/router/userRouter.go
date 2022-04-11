/**
 * @date: 2022/3/19
 * @desc: ...
 */

package router

import (
	"FileStore/global"
	"FileStore/middleware"
	"FileStore/service/controller/userController"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(routerGroup *gin.RouterGroup) {
	userRouter := routerGroup.Group("user").Use(middleware.GinLogger(global.Logger))
	{
		userRouter.POST("/login/", userController.PasswordLogin)
		userRouter.POST("/login_ldap/", userController.LdapLogin)
		userRouter.POST("/register/", userController.Register)
	}
	editUserRouter := routerGroup.Group("user_info").Use(middleware.GinLogger(global.Logger)).Use(middleware.JWTAuth())
	{
		editUserRouter.GET("/user_list/", middleware.AdminFilter(), userController.GetUserList)
		editUserRouter.GET("/:mobile/", middleware.AdminFilter(), userController.GetUserInfo)
		editUserRouter.PUT("/:mobile/", userController.EditUserInfo)
		editUserRouter.POST("/change_pwd/", userController.ChangePwd)
	}
}
