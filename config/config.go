package config

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"gomodule/pkg/redis"
	"math/rand"
	"strings"
	"sync"
	"time"
)

var Cfg = &Configuration{
	Env: Develop,
}

type Env string

const (
	Production Env = "production"
	Develop    Env = "dev"
)

type Configuration struct {
	Port            int              `mapstructure:"port"`
	Pidfile         string           `mapstructure:"pidfile"`
	Env             Env              `mapstructure:"env"`
	DebugSwitch     bool             `mapstructure:"debug_switch"`
	LogDir          string           `mapstructure:"log_dir"`
	Domains         Domains          `mapstructure:"domains"`
	ControlCenterIp []string         `mapstructure:"control_center_ip"`
	Iplib           string           `mapstructure:"ip_lib"`
	Regionlib       string           `mapstructure:"region_lib"`
	Redis           redis.PoolOption `mapstructure:"redis"`

	Mysql *MySQLConf `mapstructure:"mysql"`

	//Kafka *queue.Options `mapstructure:"kafka"`

	Topics *Topics `mapstructure:"topics"`

	HttpClient HTTPClient `mapstructure:"http_client"`

	Apps *Apps `mapstructure:"apps"`

	DspTimeoutMs int `mapstructure:"dsp_timeout_ms"`

	LoaderIntervalSecond int `mapstructure:"loader_interval_second"`

	PrintSql              bool   `mapstructure:"print_sql"`
	UnionMonitorHttpProto string `mapstructure:"union_monitor_http_proto"`

	ZyncProxy *ZyncProxy `mapstructure:"zyns-proxy"`

	Metrics Metrics `mapstructure:"metrics"`

	Aerospike Aerospike `mapstructure:"aerospike"`

	TimeoutBreaker TimeoutBreaker `mapstructure:"timeout_breaker"`
}

type TimeoutBreaker struct {
	StopCounts int `mapstructure:"stop_counts"`

	DspTimeoutCounts int `mapstructure:"dsp_timeout_counts"`
	//MaxRequests      uint32 `mapstructure:"max_requests"`
	Interval uint `mapstructure:"interval"`
	//Timeout          uint   `mapstructure:"timeout"`
}

func (c *Configuration) Print() string {
	jsonData, _ := json.MarshalIndent(c, "", "    ")
	return string(jsonData)
}

type Topics struct {
	AppReq       string `mapstructure:"app_req"`
	DspReq       string `mapstructure:"dsp_req"`
	Monitor      string `mapstructure:"monitor"`
	EventStatus  string `mapstructure:"event_status"`
	PreloadTrack string `mapstructure:"preload_track"`
	AdxCreative  string `mapstructure:"adx_creative"`
	AdCfg        string `mapstructure:"ad_cfg"`
	FilterErr    string `mapstructure:"filter_err"`
	PvReq        string `mapstructure:"pv_req"`
	ActivityReq  string `mapstructure:"ac_req"`
	ReportErr    string `mapstructure:"report_err"`
}

type Domains struct {
	AdsTrack   string `mapstructure:"ads_track"`
	UnionTrack string `mapstructure:"union_track"`
}

type Apps struct {
	ZhangyueMediaIds []string `mapstructure:"zhangyue_media_ids"`
	CsjAcceptId      string   `mapstructure:"csj_accept_id"`
	GdtAcceptId      string   `mapstructure:"gdt_accept_id"`
	SigmobAcceptId   string   `mapstructure:"sigmob_accept_id"`
	KuaishouAcceptId string   `mapstructure:"kuaishou_accept_id"`
	BaiduAcceptId    string   `mapstructure:"baidu_accept_id"`
	NotFilterPids    []string `mapstructure:"not_filter_pids"`
	ChuanYueAcceptId []string `mapstructure:"chuanyue_accept_id"`
	SouGouAcceptId   string   `mapstructure:"sougou_accept_id"`
}

type MySQLConf struct {
	Host     string `mapstructure:"host"`
	Port     uint   `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Name     string `mapstructure:"name"`
	Password string `mapstructure:"password"`
}

func SetupViper(v *viper.Viper, filename string) {
	if filename != "" {
		v.SetConfigType("yaml")
		v.SetConfigFile(filename)
		v.AddConfigPath(".")
		v.AddConfigPath("/etc/config")
	}
	v.SetDefault("host", "")
	v.SetDefault("port", 8080)
	v.SetDefault("admin_port", 6060)
	v.SetDefault("instance", 1)
	v.SetDefault("env", "dev")
	v.SetDefault("http_client.max_idle_connections", 400)
	v.SetDefault("http_client.max_idle_connections_per_host", 10)
	v.SetDefault("http_client.idle_connection_timeout_seconds", 60)
	v.SetDefault("inner_api_timeout_ms", 60*time.Second.Milliseconds())
	v.SetDefault("dsp_timeout_ms", 1*time.Second.Milliseconds())
	v.SetDefault("loader_interval_second", 60)
	v.SetDefault("external_api_urls.timeout_ms", 1000)
	v.SetDefault("metrics.port", 0)
	v.SetDefault("metrics.namespace", "ssp")
	v.SetDefault("metrics.subsystem", "")
	v.SetDefault("union_monitor_http_proto", "https")
	v.SetDefault("metrics.timeout_ms", 10000)
	v.SetDefault("timeout_breaker.dsp_timeout_counts", 50)
	v.SetDefault("timeout_breaker.interval", 20)
	v.SetDefault("timeout_breaker.timeout", 20)
	v.SetDefault("timeout_breaker.stop_counts", 200)
	// Set environment variable support:
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	//v.SetEnvPrefix("SSP")
	v.AutomaticEnv()
	v.ReadInConfig()
}

func ReadConfig(configFile string) (*Configuration, error) {
	v := viper.New()
	SetupViper(v, configFile)
	var c Configuration
	if err := v.Unmarshal(&c); err != nil {
		return nil, fmt.Errorf("viper failed to unmarshal app config: %v", err)
	}
	Cfg = &c
	return &c, nil
}

type HTTPClient struct {
	MaxIdleConns        int `mapstructure:"max_idle_connections"`
	MaxIdleConnsPerHost int `mapstructure:"max_idle_connections_per_host"`
	IdleConnTimeout     int `mapstructure:"idle_connection_timeout_seconds"`
}

type ZyncProxy struct {
	GlobalSwitchApiUrl     string   `mapstructure:"global_switch_url"`         // 九号任务全局开关
	UserFeatureApiUrl      string   `mapstructure:"user_feature_url"`          // 用户特征接口
	VideoExchangeUrl       string   `mapstructure:"video_exchange_url"`        // 从UC获取激励视频信息
	VideoExchangeReportUrl string   `mapstructure:"video_exchange_report_url"` // 上报UC激励视频信息
	WheelDrawUrl           string   `mapstructure:"wheel_draw_url"`            // 大转盘抽奖UC
	SendAmountUrl          string   `mapstructure:"send_amount_url"`           // 通过平台获取数量
	SignInfoUrl            string   `mapstructure:"sign_info_url"`             // 获取签到信息
	Nodes                  []string `mapstructure:"nodes"`                     // 代理节点
	TimeoutMs              int      `mapstructure:"timeout_ms"`
	badNodes               sync.Map
}

func (z *ZyncProxy) GetNode() string {
	if len(z.Nodes) == 1 {
		return z.Nodes[0]
	} else {
		i := rand.Intn(len(z.Nodes))
		return z.Nodes[i]
	}
}

type Metrics struct {
	Port          int    `mapstructure:"port"`
	Namespace     string `mapstructure:"namespace"`
	Subsystem     string `mapstructure:"subsystem"`
	TimeoutMillis int    `mapstructure:"timeout_ms"`
}

type Aerospike struct {
	DefaultTTL          int      `mapstructure:"default_ttl_seconds"`
	Servers             []string `mapstructure:"servers"`
	Port                int      `mapstructure:"port"`
	Username            string   `mapstructure:"username"`
	Password            string   `mapstructure:"password"`
	WriteTimeout        int      `mapstructure:"write_timeout"`
	ReadTimeout         int      `mapstructure:"read_timeout"`
	ConnectionTimeout   int      `mapstructure:"connection_timeout"`
	ConnectionQueueSize int      `mapstructure:"connection_queue_size"`

	UserProfileDB string `mapstructure:"user_profile_db"`
	DspDBName     string `mapstructure:"dsp_db"`
}

//---------------------------------------------json---------------------------------
//定义config结构体
type JsonConf struct {
	AppId  string
	Secret string
	Host   Host
}

//json中的嵌套对应结构体的嵌套
type Host struct {
	Address string
	Port    int
}
