package config

import (
	"fmt"
	"log"
	"os"
	"path"
	"runtime"
	"time"

	"gopkg.in/yaml.v3"
)

var Config *Configs
var Mode bool = true
var pwd string

func Init() {
	// 读取 配置文件
	var b []byte
	_, filename, _, _ := runtime.Caller(0)
	pwd = path.Dir(filename)
	s := os.Getenv("GIN_MODE")
	if s == "dev" {
		readConfigFile(&b, pwd, "../conf/config.dev.yaml")
	} else if s == "prod" {
		Mode = false
		readConfigFile(&b, pwd, "../conf/config.prod.yaml")
	} else {
		Mode = false
		readConfigFile(&b, pwd, "../conf/config.prod.yaml")
	}
	err2 := yaml.Unmarshal(b, &Config)
	if err2 != nil {
		log.Panicln("[!] 配置文件存在错误，请重新检查配置文件")
	}
	// 设置 log 输出
	log.SetFlags(log.LstdFlags | log.Llongfile)
	// log.SetOutput(GenLogFilename("logs"))
}

func readConfigFile(b *[]byte, pwd string, file string) {
	var err error
	*b, err = os.ReadFile(path.Join(pwd, file))
	if err != nil {
		log.Panicln("[!] 配置文件不存在，请查看conf文件夹是否存在配置文件")
	}
}

// config 结构体
type Configs struct {
	Mysql      `yaml:"Mysql"`
	Redis      `yaml:"Redis"`
	JWT        `yaml:"Jwt"`
	Cluster    `yaml:"Cluster"`
	DriverName string `yaml:"DriverName"`
}

// mysql 配置结构体
type Mysql struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Database string `yaml:"database"`
}

type Redis struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	Database int    `yaml:"database"`
}

// JWT 认证配置
type JWT struct {
	Key     string `json:"key" yaml:"key"`       // 加密key
	Issuer  string `json:"issuer" yaml:"issuer"` // 签发人
	TimeOut int    `json:"time_out" yaml:"time_out"`
}

type Cluster struct {
	IP      string `json:"ip" yaml:"ip"`
	Port    int    `json:"port" yaml:"port"`
	SSLPort int    `json:"ssl_port" yaml:"ssl_port"`
}

// 获取 database uri
func (m *Mysql) GetDatabaseConfig() string {
	return fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8",
		m.Username, m.Password, m.Host, m.Port, m.Database,
	)
}

// 获取 redis uri
func (r *Redis) GetRedisConfig() string {
	return fmt.Sprintf("%s:%d", r.Host, r.Port)
}

func GenLogFilename(class string) (logfile *os.File) {
	file := "../log/" + class + "." + time.Now().Format("200601021504") + ".log"
	logfile, err := os.OpenFile(path.Join(pwd, file), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Panicln(err)
	}
	return
}
