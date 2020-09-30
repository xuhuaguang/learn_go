package learn_go

import (
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"gomodule/config"
	"gomodule/utils"
	"log"
	"math/rand"
	"os"
	"testing"
	"time"
)

//测试加载配置文件
func TestConfigYaml(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	pflag.String("conf", "./app.yaml", "set configuration `file`")
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
	return nil
}

func TestConfigJson(t *testing.T) {
	conf := viper.New()
	conf.AddConfigPath(".")
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
