package entity

type RequestInfo struct {
	Date      int64 `json:"date,omitempty"`       // 日期
	Hour      int   `json:"hour,omitempty"`       // 小时
	EventTime int64 `json:"event_time,omitempty"` // 事件时间

	ID      string `json:"uid"`                // 请求ID
	Type    string `json:"type"`               // 日志类型 api
	ReqTime string `json:"req_time"`           // 请求时间 秒
	Pid     string `json:"pid"`                // 广告位ID
	PidSign []int  `json:"pid_sign,omitempty"` // 加密的pid对应策略集ID
	MediaID string `json:"media_id"`           // 广告位所属媒体，关联Redis

	AppId      string `json:"app_id"`      // 应用包名
	AppPackage string `json:"app_package"` // 应用包名
	AppName    string `json:"app_name"`    // 应用名
	AppVersion string `json:"app_ver"`     // 应用名
	Priority   int    `json:"priority"`

	GeoLat            string `json:"geo_lat"`            // GPS纬度（-90-90）
	GeoLon            string `json:"geo_lon"`            // GPS经度（-180-180）
	DeviceImei        string `json:"device_imei"`        // IMEI号
	DeviceAdid        string `json:"device_adid"`        // 安卓为android id, ios则为idfa
	DeviceOaid        string `json:"device_oaid"`        // 安卓为android oaid
	DeviceIdfa        string `json:"device_idfa"`        // ios则为idfa
	DeviceOpenudid    string `json:"device_openudid"`    // 苹果设备唯一标识号
	DeviceIdfv        string `json:"device_idfv"`        // ios idfv
	DeviceDpi         string `json:"device_dpi"`         // 屏幕dpi
	DeviceMac         string `json:"device_mac"`         // MAC 地址
	DeviceOsVersion   string `json:"device_type_os"`     // 操作系统版本
	DeviceType        string `json:"device_type"`        // 设备类型（-1:未知, 0:phone, 1:pad, 2:pc, 3:tv 4:wap）
	DeviceBrand       string `json:"device_brand"`       // 设备品牌、生产厂商（"HUAWEI"、"samsung"、"Apple"、"Xiaomi"等）
	DeviceBrandVer    string `json:"device_brand_ver"`   // 设备品牌、生产厂商（"HUAWEI"、"samsung"、"Apple"、"Xiaomi"等）
	DeviceModel       string `json:"device_model"`       // 手机型号（"MHA-AL00"、"SM-G9280"、"iPhone8"、"MIX 2S"等）
	DeviceHeight      string `json:"device_height"`      // 屏幕宽度
	DeviceWidth       string `json:"device_width"`       // 屏幕高度
	DeviceImsi        string `json:"device_imsi"`        // 网络运营商代码取值
	DeviceNetwork     string `json:"device_network"`     // 网络类型：（1：WIFI，2：4G，3：3G，4：2G，5：其他,6：5G）
	DeviceOs          string `json:"device_os"`          // Android/iOS/WP/Others 字符串，注意大小写
	DeviceDensity     string `json:"device_density"`     // 屏幕密度
	DeviceIP          string `json:"ip"`                 // 客户端 ip（必须是外网可访问IP）
	DeviceUa          string `json:"ua"`                 // User-Agent(须进行一次urlencode)必须是标准Webview UA而非自定义UA
	DeviceOrientation string `json:"device_orientation"` // 横竖屏“1” 横屏 “0” 竖屏
	DeviceLan         string `json:"device_lan"`         // 目前使用的语言国家 zh-CN
	VerIos            string `json:"ver"`                // 版本
	IsMobile          string `json:"is_mobile"`          // 是否为移动端 默认 1
	DeviceIsroot      string `json:"device_isroot"`      // Android 设备是否 ROOT。1--是, 0--否/ 未知(默认)
	RegionID          string `json:"region_id"`          // 对应的IP库代码
	Size              string `json:"size"`               // 平台配置广告位尺寸 width x height
	DeviceID          string `json:"device_id"`          // 同 device_adid
	JailBreak         string `json:"jailBreak"`          // 越狱
	UserId            string `json:"usr"`                // i账号
	Secure            string `json:"secure"`             // 是否https 0/1
	ExtData           string `json:"ext_data"`
	Status            string `json:"status"`
	Province          string `json:"province"` // 省
	City              string `json:"province"` // 市
	ChKey             string `json:"ch_key"`   //定向频道key

	ChannelId    string `json:"channel_id"`    //渠道号
	InnerVersion string `json:"inner_version"` //客户端内部版本号
	VersionCode  string `json:"version_code"`  //version_code
	AppPlatform  string `json:"app_platform"`  //平台号
	Network      string `json:"network"`       //客户端网络类型
	OsVer        string `json:"os_ver"`        //操作系统版本
	PluginName   string `json:"plugin_name"`   //广告插件名称
	PluginVer    string `json:"plugin_ver"`    //广告插件版本
	BiNetWork    string `json:"bi_network"`    //网络类型
	HostVersion  string `json:"host_version"`  //
	SdkVersion   string `json:"sdk_version"`   //
	BookId       string `json:"book_id"`       //计算广告-增加上下文特征
	BookName     string `json:"book_name"`     //计算广告-增加上下文特征
	PageNum      int    `json:"page_num,omitempty"`
}
