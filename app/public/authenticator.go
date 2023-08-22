package public

import (
	"errors"
	"log"
	"reflect"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	ruleError      string = "为必填字段"
	regexpError    string = "校验错误: "
	emailRegexp    string = `^([\w\.\_\-]{2,10})@(\w{1,}).([a-z]{2,4})$`
	emailPrompt    string = ""
	accountRegexp  string = `^[a-zA-Z][a-zA-Z0-9_]{4,15}$`
	accountPrompt  string = "字母开头，允许5-16字节，允许字母数字下划线"
	passwordRegexp string = `/^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[$@$!%*?&.])[A-Za-z\d$@$!%*?&.]{6, 20}/`
	passwordPrompt string = `最少8个最多16个字符，至少1个大写字母，1个小写字母，1个数字和1个特殊字符`
)

var RegexpMap = map[string]RegexpInfo{
	"email":    {regexp.MustCompile(emailRegexp), ""},              // 邮箱认证规则
	"account":  {regexp.MustCompile(accountRegexp), accountPrompt}, // 账户认证规则，字母开头，允许5-16字节，允许字母数字下划线
	"password": {regexp.MustCompile(passwordPrompt), passwordPrompt},
}

type RegexpInfo struct {
	Regexp *regexp.Regexp
	Prompt string
}

type TestRule struct {
	Name string `rule:"notnull" label:"名字" regexp:"^([\\w\\.\\_\\-]{2,10})@(\\w{1,}).([a-z]{2,4})$"`
	Age  string `rule:"notnull" label:"年龄" regexp:"email"`
}

func Authenticator(ctx *gin.Context, params any) (err error) {
	if err := ctx.ShouldBind(params); err != nil {
		log.Panicln(err)
	}
	err = ParserStruct(params)
	if err != nil {
		return err
	}
	return nil
}

func ParserStruct(params any) error {
	v := reflect.ValueOf(params)
	if v.Kind() != reflect.Ptr {
		log.Println("请传入一个指针")
		return errors.New("程序错误")
	}
	elem := v.Elem()
	var buillder strings.Builder
	// 循环校验每一个字段
	for i := 0; i < elem.NumField(); i++ {
		LegalVerification(elem, i, &buillder)
	}
	// 调用 Check
	check := v.MethodByName("Check")
	if check.IsValid() {
		check.Call([]reflect.Value{reflect.ValueOf(&buillder)})
	}
	if buillder.Len() == 0 {
		return nil
	}
	return errors.New(buillder.String())
}

// 校验字段是否符合规则
func LegalVerification(v reflect.Value, i int, buillder *strings.Builder) {
	tag := v.Type().Field(i).Tag
	rule := tag.Get("rule")
	label := tag.Get("label")
	jsonTag := tag.Get("json")
	if rule != "" && rule == "notnull" && v.Field(i).IsZero() {
		if label == "" && jsonTag != "-" {
			buillder.WriteString(jsonTag + ruleError + EndMark)
		} else {
			buillder.WriteString(label + ruleError + EndMark)
		}
	} else {
		// 在内容不为空的情况下进行内容合法化
		regexpRule := tag.Get("regexp")
		if regexpRule != "" {
			regexpInfo, ok := RegexpMap[regexpRule]
			if ok {
				if !regexpInfo.Regexp.MatchString(v.Field(i).String()) {
					if label == "" && jsonTag != "-" {
						buillder.WriteString(jsonTag + regexpError + regexpInfo.Prompt + EndMark)
					} else {
						buillder.WriteString(label + regexpError + regexpInfo.Prompt + EndMark)
					}
				}
			} else {
				reg := regexp.MustCompile(regexpRule)
				if !reg.MatchString(v.Field(i).String()) {
					if label == "" && jsonTag != "-" {
						buillder.WriteString(jsonTag + regexpError + EndMark)
					} else {
						buillder.WriteString(label + regexpError + EndMark)
					}
				}
			}
		}
	}
}
