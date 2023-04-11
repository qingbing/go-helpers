package validator

import (
	"fmt"
	"reflect"

	CN_ZH "github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTran "github.com/go-playground/validator/v10/translations/zh"
)

// 初始化验证器实例
func New() *validate {
	valid := validator.New()
	// 设置一个函数，获取 struct-tag 里的自定义 label 作为字段名
	valid.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := fld.Tag.Get("label")
		return name
	})
	// 注册验证器的翻译器
	uni := ut.New(CN_ZH.New())
	trans, _ := uni.GetTranslator("zh")
	zhTran.RegisterDefaultTranslations(valid, trans)
	return &validate{
		validator: valid,
		trans:     trans,
	}
}

// 验证器结构体
type validate struct {
	validator *validator.Validate
	trans     ut.Translator
	errors    []error
}

// 注册一个验证器
// errTemplate: "{0}输入错误"
func (v *validate) RegisterValidation(tag string, fn validator.Func, errTemplate string, callValidationEvenIfNull ...bool) error {
	v.validator.RegisterTranslation(tag, v.trans, func(ut ut.Translator) error {
		return ut.Add(tag, errTemplate, true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T(tag, fe.Field())
		return t
	})
	return v.validator.RegisterValidation(tag, fn, callValidationEvenIfNull...)
}

// 设置结构体的 validate tag
func (v *validate) SetTagName(name string) *validate {
	v.validator.SetTagName(name)
	return v
}

// 设置集合验证器别名
// RegisterAlias("password", "required,min=6,max=18")
func (v *validate) RegisterAlias(alias, tags string) *validate {
	v.validator.RegisterAlias(alias, tags)
	return v
}

// 验证结构体
func (v *validate) Struct(data interface{}) *validate {
	err := v.validator.Struct(data)
	if err != nil {
		v.errors = make([]error, 0)
		for _, validErr := range err.(validator.ValidationErrors) {
			v.errors = append(v.errors, fmt.Errorf(validErr.Translate(v.trans)))
		}
	}
	return v
}

// 获取一个验证错误信息
func (v *validate) GetErrors() []error {
	return v.errors
}

// 获取一个验证错误信息
func (v *validate) GetError() error {
	if len(v.errors) == 0 {
		return nil
	}
	return v.errors[0]
}
