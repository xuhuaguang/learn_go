package models

const (
	FixedMinute int = 5  //固定值
	MaxMinute   int = 25 //时间范围：最大值  --70%的用户
	MinMinute   int = 0  //时间范围：最小值  --70%的用户

	MaxMinuteUp int = 35 //时间范围：最大值  --30%的用户
	MinMinuteUp int = 15 //时间范围：最小值  --30%的用户

	FreqSidTime  string = "%d_time"
	FreqSidCount string = "%d_count"
	FreqUsrTanx  string = "tanx_screen"

	MinRandRange    int = 0   //最小随机范围值
	MaxRandRange    int = 100 //最大随机范围值
	MiddleRandRange int = 70  //分割线数值

	CPM = 1000

	// 基本全额单位 最小单位是分 二价+1分结算
	// 一元=10角，一元=100分，一元=1000厘，一元=10000毫
	// 为了不存小数据，程序内部及报表数据都存的分*CPM
	// 100 元/cpm	 0.100/个
	// 1元/cpm       0.001/个
	// 0.1+0.01元 /cpm    0.00011/个 ×　　Ssp_Yuan = 11 Ssp_Fen
	Ssp_Yuan int64 = 100 * CPM
	Ssp_Jiao       = 10 * CPM
	Ssp_Fen        = 1 * CPM

	Ssp_OldFen         = 100
	Ssp_OldYuan  int64 = 10000
	OldPriceDate       = "2020-03-28"

	RMB_Yuan_TO_Fen = 100

	URL_SIGN_TOKEN = "ZY"

	// yyyy-MM-dd HH:mm:ss
	Format_yyyy_MM_dd_HH_mm_ss = "2006-01-02 15:04:05"
	Format_HH_mm_ss            = "15:04:05"
	DateFormat_yyyy_MM_dd      = "2006-01-02"
	DateFormat_yyyyMMdd        = "20060102"

	SecondMs               = 1000
	Start_Default_HH_mm_ss = "00:00:00"
	End_Default_HH_mm_ss   = "23:59:59"
)
