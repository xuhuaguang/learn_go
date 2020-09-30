package learn_go

import (
	"crypto/tls"
	"fmt"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"learn_go/config"
	"learn_go/pkg/izap"
	"learn_go/pkg/mysql"
	"learn_go/pkg/redis"
	"learn_go/utils"
	"log"
	"math/rand"
	"net/http"
	"os"
	"testing"
	"time"
)

//测试加载配置文件
func TestConfigYaml(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	pflag.String("conf", "../app.yaml", "set configuration `file`")
	pflag.Parse()
	_ = viper.BindPFlags(pflag.CommandLine)
	configFile := viper.GetString("conf")
	if utils.IsEmpty(configFile) {
		pflag.Usage()
		log.Println("can't not found config file")
		os.Exit(1)
	}
	cfg, err := config.ReadConfig(configFile)
	if err != nil {
		log.Fatalf("Configuration could not be loaded or did not pass validation: %v", err)
	}
	if err := serve(cfg); err != nil {
		log.Fatalf("ssp-server failed: %v", err)
	}
}

func serve(cfg *config.Configuration) error {
	izap.Init()
	izap.Get().Info("start server")
	izap.Get().Info("config content \n", cfg.Print(), "\n")
	if cfg.Env == config.Production {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.New()
	router.Use(ginzap.RecoveryWithZap(izap.GetError().Desugar(), true))
	// 初始化依赖
	dspHttpClient := createHttpClient(cfg, cfg.DspTimeoutMs)
	fmt.Println(dspHttpClient)
	connPool, err := redis.NewPool(&cfg.Redis)
	if err != nil {
		izap.Get().Errorf("init redis pool error %s", err.Error())
		os.Exit(10)
	}
	fmt.Println(connPool)
	mysqlDB := mysql.OpenDB(cfg.Mysql, cfg.PrintSql)
	fmt.Println(mysqlDB)

	router.GET("/ping", ping)
	izap.Get().Infof("ssp server start on port %d", cfg.Port)
	err = router.Run(fmt.Sprintf(":%d", cfg.Port))
	return err
}

func TestConfigJson(t *testing.T) {
	conf := viper.New()
	conf.AddConfigPath("../")
	conf.SetConfigName("config_json")
	conf.SetConfigType("json")
	if err := conf.ReadInConfig(); err != nil {
		panic(err)
	}
	fmt.Println(conf.GetString("appId"))
	fmt.Println(conf.GetString("secret"))
	fmt.Println(conf.GetString("host.address"))
	fmt.Println(conf.GetString("host.port"))

	//直接反序列化为Struct
	var configJson config.JsonConf
	if err := conf.Unmarshal(&configJson); err != nil {
		fmt.Println(err)
	}

	fmt.Println(configJson.Host)
	fmt.Println(configJson.AppId)
	fmt.Println(configJson.Secret)
}

func createHttpClient(cfg *config.Configuration, timeout int) *http.Client {
	return &http.Client{
		Timeout: time.Millisecond * time.Duration(timeout),
		Transport: &http.Transport{
			MaxIdleConns:        cfg.HttpClient.MaxIdleConns,
			MaxIdleConnsPerHost: cfg.HttpClient.MaxIdleConnsPerHost,
			IdleConnTimeout:     time.Duration(cfg.HttpClient.IdleConnTimeout) * time.Second,
			TLSClientConfig:     &tls.Config{InsecureSkipVerify: true}, // disable verify
		},
	}
}

func ping(c *gin.Context) {
	usr := c.Query("usr")
	md5 := utils.Md5(usr)
	c.String(http.StatusOK, "i号MD5后尾号的值为: ==> "+md5[len(md5)-1:])
}
