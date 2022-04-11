/*
 * @date: 2021/12/15
 * @desc: ...
 */

package initialize

import (
	"FileStore/common/customValidator"
	"FileStore/global"
	"github.com/gin-gonic/gin/binding"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

// InitCustomValidator 初始化自定义校验器
func InitCustomValidator() {
	registerMobileValidator()
}

// registerMobileValidator 手机号码校验器
func registerMobileValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("mobile", customValidator.ValidateMobile)
		_ = v.RegisterTranslation(
			"mobile",
			global.Trans,
			func(ut ut.Translator) error {
				return ut.Add("mobile", "{0} 非法的手机号码!", true) // see universal-translator for details
			},
			func(ut ut.Translator, fe validator.FieldError) string {
				t, _ := ut.T("mobile", fe.Field())
				return t
			})
	}
}
