package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/locales/zh_Hant_TW"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_transactions "github.com/go-playground/validator/v10/translations/en"
	zh_transactions "github.com/go-playground/validator/v10/translations/zh"
)

func Transactions() gin.HandlerFunc {
	return func(c *gin.Context) {
		uni := ut.New(en.New(), zh.New(), zh_Hant_TW.New())
		locale := c.GetHeader("locale")
		trans, _ := uni.GetTranslator(locale)
		v, ok := binding.Validator.Engine().(*validator.Validate)
		if ok {
			switch locale {
			case "zh":
				_ = zh_transactions.RegisterDefaultTranslations(v, trans)
				break
			case "en":
				_ = en_transactions.RegisterDefaultTranslations(v, trans)
				break
			default:
				_ = zh_transactions.RegisterDefaultTranslations(v, trans)
				break
			}
			c.Set("trans", trans)
		}
		c.Next()
	}
}
