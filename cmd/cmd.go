package cmd

import (
	"flag"
	"github.com/garyburd/redigo/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"pea-web/api/model"
)

var (
	Conf   *Configs
	DB     *gorm.DB
	RDS    redis.Pool
	pagram = flag.String("c", "./pea.yaml", "请配置yaml文件路径")
)

type Configs struct {
	DevEnv     string `yaml:"DevEnv"`     //开发环境
	HostURL    string `yaml:"HostURL"`    //域名
	HostPort   string `yaml:"HostPort"`   //域名端口
	LoggerFile string `yaml:"LoggerFile"` //日志文件地址
	MysqlDB    string `yaml:"MysqlDB"`    //数据库地址
	RedisDB    string `yaml:"RedisDB"`    //redis连接
	AliyunOss  struct {
		Host         string `yaml:"Host"`
		Bucket       string `yaml:"Bucket"`
		Endpoint     string `yaml:"Endpoint"`
		AccessId     string `yaml:"AccessId"`
		AccessSecret string `yaml:"AccessSecret"`
	} `yaml:"AliyunOss"`
	SMTP struct {
		Host     string `yaml:"Host"`
		Port     string `yaml:"Port"`
		Username string `yaml:"Username"`
		Password string `yaml:"Password"`
		SSL      bool   `yaml:"SSL"`
	} `yaml:"SMTP"`
}

// 调试模式
func DebugStart() error {
	flag.Parse()
	err := initConfig(*pagram)
	if err != nil {
		return err
	}
	//初始化数据库
	err = initDB()
	if err != nil {
		return err
	}
	//初始化reids
	return nil
}

// 发布模式
func PublishStart() error {
	flag.Parse()
	err := initConfig(*pagram)
	if err != nil {
		return err
	}
	//初始化数据库
	err = initDB()
	if err != nil {
		return err
	}
	return nil
}

//初始化数据库
func initDB() error {
	db, err := gorm.Open("mysql", Conf.MysqlDB)
	if err != nil {
		return err
	}
	//连接池配置
	db.DB().SetMaxIdleConns(50)        //初始化数据库连接数
	db.DB().SetMaxOpenConns(70)        //额外增开20
	db.DB().SetConnMaxLifetime(60 * 1) //链接时长
	db.SingularTable(true)             //默认表名单数
	db.LogMode(true)
	DB = db
	return db.AutoMigrate(model.Model...).Error
}

//初始化redis
func initRedis() {
	RDS = redis.Pool{
		MaxIdle:     20,     //初始化最大连接数
		MaxActive:   50,     //最大打开连接数
		IdleTimeout: 60 * 1, //空闲3分钟回收
		Dial: func() (conn redis.Conn, e error) {
			return redis.Dial("tcp", Conf.RedisDB)
		},
	}
}

// 解析配置文件
func initConfig(f string) error {
	file, err := ioutil.ReadFile(f)
	if err != nil {
		return err
	}
	Conf = new(Configs)
	err = yaml.Unmarshal(file, Conf)
	if err != nil {
		return err
	}
	return nil
}
