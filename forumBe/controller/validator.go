package controller

import (
	"fmt"
	"forum/models"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

var trans ut.Translator

func InitTrans(locale string) (err error) {
	// 修改gin 中的binding方法
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 注册一个获取json tag 的自定义方法
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
		// 给SignupParams 注册自定义校验的方法
		v.RegisterStructValidation(SignUpParamStructLevelValidation)
		zhT := zh.New()
		enT := en.New()
		uni := ut.New(enT, zhT, enT)
		var ok bool
		trans, ok = uni.GetTranslator(locale)
		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s) failed", locale)
		}
		// 注册翻译器
		switch locale {
		case "en":
			err = enTranslations.RegisterDefaultTranslations(v, trans)
		case "zh":
			err = zhTranslations.RegisterDefaultTranslations(v, trans)
		default:
			err = enTranslations.RegisterDefaultTranslations(v, trans)
		}
		return
	}
	return
}

// 定义一个去掉结构体名称前缀的自定义方法：
func removeTopStruct(fields map[string]string) map[string]string {
	res := map[string]string{}
	for field, err := range fields {
		res[field[strings.Index(field, ".")+1:]] = err
	}
	return res
}

// SignUpParamStructLevelValidation 自定义SignUpParam结构体校验函数
func SignUpParamStructLevelValidation(sl validator.StructLevel) {
	su := sl.Current().Interface().(models.User)
	if su.ConfirmPassword != su.Password {
		// 输出错误提示信息，最后一个参数就是param
		sl.ReportError(su.ConfirmPassword, "confirm_password", "ConfirmPassword", "eqfield", "password")
	}
}
