package public

import (
	"regexp"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestAuthenticator(t *testing.T) {
	type args struct {
		ctx    *gin.Context
		params any
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Authenticator(tt.args.ctx, tt.args.params); (err != nil) != tt.wantErr {
				t.Errorf("Authenticator() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestParser(t *testing.T) {
	// test := &TestRule{Name: "2919390584@qq.com", Age: "2919390584@qq.com"}
	// err := ParserStruct(test)
	// t.Log(err)
	// t.Log(RegexpMap["password"].Regexp.MatchString(test.Age))
	// reg := regexp.MustCompile("^([\\w\\.\\_\\-]{2,10})@(\\w{1,}).([a-z]{2,4})$")
	// t.Log(reg.MatchString(test.Name))
	matched, _ := regexp.MatchString("[0-3]", "0")
	t.Log(matched)
}
