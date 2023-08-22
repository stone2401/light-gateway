// 此文件为测试文件，无意义，仅作为权限认证测试使用
package config

import "sync"

type empty struct{}
type Authority map[string]empty

var sonce sync.Once
var admin = Authority{
	"/api/v1/admin_login/login": {},
}

var tourist = Authority{
	"/api/v1/ping": {},
}

var app = Authority{
	"/api/v1/app/app_add": {},
}

var authoritys = map[string][]Authority{
	"admin":   {admin, app, tourist},
	"app":     {app, tourist},
	"tourist": {tourist},
}

var temp = map[string]Authority{}

func GetAuthoritys() map[string]Authority {
	sonce.Do(func() {
		for key, value := range authoritys {
			temp[key] = make(Authority)
			for _, vv := range value {
				for kkk, vvv := range vv {
					temp[key][kkk] = vvv
				}
			}
		}
	})
	return temp
}
