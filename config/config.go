package config

import (
	"bytes"
	"embed"
	"github.com/spf13/viper"
	"gorm.io/gorm/logger"
)

// ConfEnv env环境变量
var ConfEnv string

type (
	Conf struct {
		Server Server `yaml:"server"`
		Jwt    Jwt    `yaml:"jwt"`
		Log    Log    `yaml:"log"`
		Mysql  Mysql  `yaml:"mysql"`
		Redis  Redis  `yaml:"redis"`
		Amqp   Amqp   `yaml:"amqp"`
		Oss    Oss    `yaml:"oss"`
	}
	Server struct {
		Mode            string `yaml:"mode" default:"debug"`
		DefaultPageSize int    `yaml:"defaultPageSize" default:"10"`
		MaxPageSize     int    `yaml:"maxPageSize" default:"500"`
		FileUploadPath  string `yaml:"fileUploadPath"`
	}
	Jwt struct {
		TokenExpire int64  `yaml:"tokenExpire" default:"864000"`
		TokenKey    string `yaml:"tokenKey" default:"Authorization"`
		TokenIssuer string `yaml:"tokenIssuer" default:"gin-framework"`
		Secret      string `yaml:"secret"`
	}
	Log struct {
		Debug    bool   `yaml:"debug" default:"true"`
		FileName string `yaml:"fileName" default:"gin-framework"`
		DirPath  string `yaml:"dirPath" default:"runtime/logs"`
	}
	Mysql []struct {
		Host         string          `yaml:"host" default:"127.0.0.1"`
		Port         string          `yaml:"port" default:"3306"`
		User         string          `yaml:"user" default:"root"`
		Password     string          `yaml:"password" default:"123456"`
		DbName       string          `yaml:"dbname"`
		Prefix       string          `yaml:"prefix" default:""`
		MaxIdleConns int             `yaml:"maxIdleConns" default:"10"`
		MaxOpenConns int             `yaml:"maxOpenConns" default:"100"`
		MaxLifeTime  int             `yaml:"maxLifeTime" default:"60"`
		LogLevel     logger.LogLevel `yaml:"logLevel" default:"1"`
	}
	Redis struct {
		Host        string `yaml:"host" default:"127.0.0.1"`
		Port        string `yaml:"port" default:"6379"`
		Password    string `yaml:"password"`
		DbNum       int    `yaml:"dbNum" default:"0"`
		LoginPrefix string `yaml:"loginPrefix" default:"mqenergy_login_auth_"`
	}
	Amqp struct {
		Host     string `yaml:"host" default:"127.0.0.1"`
		Port     string `yaml:"port" default:"5672"`
		User     string `yaml:"user" default:"guest"`
		Password string `yaml:"password" default:""`
		Vhost    string `yaml:"vhost" default:""`
	}
	Oss struct {
		EndPoint        string `yaml:"endPoint" default:"https://oss-cn-shanghai.aliyuncs.com"`
		AccessKeyId     string `yaml:"accessKeyId"`
		AccessKeySecret string `yaml:"accessKeySecret"`
		BucketName      string `yaml:"bucketName"`
	}
)

//go:embed yaml
var yamlCfg embed.FS

// InitConfig 初始化配置
func InitConfig() (*Conf, error) {
	var cfg *Conf
	v := viper.New()
	v.SetConfigType("yaml")
	yamlConf, _ := yamlCfg.ReadFile("yaml/config." + ConfEnv + ".yaml")
	if err := v.ReadConfig(bytes.NewBuffer(yamlConf)); err != nil {
		return nil, err
	}
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
