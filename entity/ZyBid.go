package entity

type ZyBid struct {
	ID                string   `json:"id"`                      // 竞价ID
	AdverId           string   `json:"ader_id"`                 // 广告主 ID：(流量平台广告主送审返回的 ID),
	CreativeId        string   `json:"cid"`                     // 创意 ID：(流量平台创意送审返回的 ID)
	Price             int      `json:"price"`                   // 分/CPM 计，低于底价的将竞价失败
	Adm               string   `json:"adm"`                     // 非信息流广告填写物料访问地址；信息流广告 填写信息流模板相关 JSON 数据并 且 templateid 必填
	LandingPageUrl    string   `json:"durl"`                    // 到达页地址
	LandingPageAction int      `json:"adck"`                    // 到达页 durl 打开类型： 1：网页类型（默认） 2：下载类型
	DeepLinkUrl       string   `json:"deep_link,omitempty"`     // deeplink 地址：(当请求中 atype 等于 3 或 4 时此字段才有效，媒体会优先处理 deep_link, 当无法处理 deep_link 时转而处理 durl,如果 deep_link 可以落地则不会处理 durl，无论是 否传 deep_link 字段，都必须填写 durl 字段; 如果媒体不支持 deeplink, 将会直接执行 durl)
	Width             int      `json:"w,omitempty"`             // 物料宽
	Height            int      `json:"h,omitempty"`             // 物料高
	ImpUrls           []string `json:"nurl,omitempty"`          // 曝光监测地址，支持价格宏替换
	ClickUrls         []string `json:"curl,omitempty"`          // 点击监测上报地址
	DealId            string   `json:"dealid,omitempty"`        // PMP 售卖方式必须带上,并且必须跟广告请求 所带相等
	TemplateId        string   `json:"templateid,omitempty"`    // 原生广告必填（Request 对应的模板 ID)
	AppPackage        string   `json:"package_name,omitempty"`  // （安卓应用包名、ios 为 APP Store 中的 ID）：下载和 deeplink 类广告建议填写
	AppName           string   `json:"app_name,omitempty"`      // 应用名如：“锤子便签”；下载和 deeplink 类广告建议填写
	AppVersion        string   `json:"app_version,omitempty"`   // 应用版本号：下载和 deeplink 类广告填
	From              string   `json:"from,omitempty"`          // 广告来源回传给SDK使用
	CreativeIdType    int      `json:"creative_id_type"`        //
	BidType           int      `json:"bid_type,omitempty"`      // 品牌出价方式  10=cpt、11=cpm
	CampaignStartDate int64    `json:"campaign_date,omitempty"` // 品牌表示素材希望投放的日期，在开屏这种预加载时间较长的流量中使用
	CampaignEndDate   int64    `json:"expire_time,omitempty"`
	CptRatio          int      `json:"cpt_ratio,omitempty"`
	CardPlace         int      `json:"place,omitempty"` //针对卡片流
}
