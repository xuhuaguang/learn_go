package entity

type Accept struct {
	Id          int    `gorm:"column:id" json:"id"`                     // 接收方id
	Name        string `gorm:"column:name" json:"name"`                 // 接收方名称
	ServiceType string `gorm:"column:service_type" json:"service_type"` // 业务类型（1：品牌 2：自营DSP 3：第三方DSP 4；联盟）
	Type        string `gorm:"column:type" json:"type"`                 // 类型（1:ssp 2:adx 3:dsp）
	AppFields   string `gorm:"column:app_fields" json:"app_fields"`     // 应用对接必须字段
	Url         string `gorm:"column:url" json:"url"`                   // url地址
	PriceKey    string `gorm:"column:price_key" json:"price_key"`       // 成交价密钥
	Token       string `gorm:"column:token" json:"token"`               // 对接token
	QpsLimit    int    `gorm:"column:qps_limit" json:"qps_limit"`       // 请求次数限制（次/每秒）
	Code        string `gorm:"column:code" json:"code"`                 // 接收方码
	LogoUrl     string `gorm:"column:logo_url" json:"logo_url"`         // 接收方码
}
