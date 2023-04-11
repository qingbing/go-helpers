package validator

import (
	"log"
	"testing"

	"github.com/go-playground/validator/v10"
)

type User struct {
	Contact         string         `validate:"customValidate" label:"联系"`
	Name            string         `validate:"required,alphaunicode"`
	Age             uint8          `validate:"gte=10,lte=30" label11:"年龄"`
	Phone           string         `validate:"required,e164"`
	Email           string         `validate:"required,email"`
	FavouriteColor1 string         `validate:"iscolor"`
	FavouriteColor2 string         `validate:"hexcolor|rgb|rgba|hsl|hsla"`
	Address         *Address       `validate:"required"`
	ContactUser     []*ContactUser `validate:"required,gte=1,dive"` // dive 表示验证深入下一层
	hobby           []string       `validate:"required,gte=2,dive,required,gte=2,alphaunicode"`
}
type Address struct {
	Province string `validate:"required" label11:"省份"`
	City     string `validate:"required"`
}
type ContactUser struct {
	Name    string   `validate:"required,alphaunicode" label11:"联系人姓名"`
	Age     uint8    `validate:"gte=20,lte=30"`
	Phone   string   `validate:"required_without_all=Email Address,omitempty,e164"`
	Email   string   `validate:"required_without_all=Phone Address,omitempty,email"`
	Address *Address `validate:"required_without_all=Phone Email"`
}

func customValidate(fl validator.FieldLevel) bool {
	xx := fl.Field().String()
	if len(xx) >= 4 {
		return true
	}
	return false
}

func Test_ValidStruct(t *testing.T) {
	address := &Address{
		Province: "",
		City:     "成都",
	}
	contactUser1 := &ContactUser{
		Name:    "",
		Age:     30,
		Phone:   "+8613800138000",
		Email:   "dd@sdf.com",
		Address: address,
	}
	contactUser2 := &ContactUser{
		Name:    "",
		Age:     30,
		Phone:   "+8613800138000",
		Email:   "dd@sdf.com",
		Address: address,
	}
	user := &User{
		// Contact:         "xx",
		Name:            "",
		Age:             9,
		Phone:           "+8613800138000",
		Email:           "ddd@dd.com",
		FavouriteColor1: "#ffff",
		FavouriteColor2: "rgb(255,255,255)",
		Address:         address,
		ContactUser:     []*ContactUser{contactUser1, contactUser2},
		hobby:           []string{"足球", "篮球"},
	}

	validate := New()
	validate.RegisterValidation("customValidate", customValidate, "{0}输入了沙内容")
	validate.Struct(user)
	log.Println(validate.GetError())
	log.Println(validate.GetErrors())
}
