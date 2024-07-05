package public

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales"
	en2 "github.com/go-playground/locales/en"
	zh2 "github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	en1 "gopkg.in/go-playground/validator.v9/translations/en"
	zh1 "gopkg.in/go-playground/validator.v9/translations/zh"
	"reflect"
	"strings"
)

var (
	Validate *validator.Validate
	Uni      *ut.UniversalTranslator
)

// 初始化过滤器
func InitValidate() {
	var (
		en, zh locales.Translator
	)
	en = en2.New()
	zh = zh2.New()
	Uni = ut.New(en, zh)

	Validate = validator.New()
	transEn, _ := Uni.GetTranslator("en")
	transZh, _ := Uni.GetTranslator("zh")
	_ = zh1.RegisterDefaultTranslations(Validate, transZh)
	_ = en1.RegisterDefaultTranslations(Validate, transEn)
}

func BindingValidParams(c *gin.Context, param interface{}, t reflect.Type) (err error) {
	v := c.Value("trans")
	lang := c.Value("lang")
	if len(lang.(string)) < 1 {
		lang = "zh"
	}
	trans, ok := v.(ut.Translator)
	if !ok {
		trans, _ = Uni.GetTranslator("zh")
	}
	if err = Validate.Struct(param); err != nil {
		errs := err.(validator.ValidationErrors)
		sliceErrs := ""
		for _, e := range errs {
			field, _ := t.FieldByName(e.Field())
			newTrans := field.Tag.Get(lang.(string))
			sliceErrs = strings.Replace(e.Translate(trans), e.Field(), newTrans, -1)
			break
		}
		return errors.New(sliceErrs)
	}
	return
}
