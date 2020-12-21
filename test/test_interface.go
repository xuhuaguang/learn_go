package test

import "fmt"

type SpecialBizCodeName string

const (
	PotentSCREEN     SpecialBizCodeName = "POTENT_SCREEN"
	SelectCardsVideo SpecialBizCodeName = "SELECT_CARDS_VIDEO"
)

type item struct {
	Title     string `json:"title,omitempty"`
	Content   string `json:"content,omitempty"`
	Icon      string `json:"icon,omitempty"`
	IconTitle string `json:"icon_title,omitempty"`
	FromLogo  string `json:"from_logo,omitempty"`
	From      string `json:"from,omitempty"`
}

//特殊广告位模板
type templates struct {
	PotentScreen *PotentScreen `json:"potentScreen,omitempty"`
	CardsVideo   *CardsVideo   `json:"cardsVideo,omitempty"`
}

type CardsVideo struct {
	Title    string `json:"title,omitempty"`
	Content  string `json:"content,omitempty"`
	VideoUrl string `json:"videoUrl,omitempty"`
	PicUrl   string `json:"picUrl,omitempty"`
	Mark     string `json:"mark,omitempty"`
}

type PotentScreen struct {
	VideoUrl   string `json:"videoUrl,omitempty"`
	VideoCover string `json:"videoCover,omitempty"`
	PicUrl     string `json:"picUrl,omitempty"`
	Mark       string `json:"mark,omitempty"`
}

//模板函数
type buildTemplatesParser func(response *item, templateId string) *templates

//--NO3
var buildTemplatesParserMap = map[SpecialBizCodeName]buildTemplatesParser{
	PotentSCREEN:     buildPopupIconBookstore,
	SelectCardsVideo: buildSelectCardsVideo,
}

//--NO4
func buildPopupIconBookstore(req *item, templateId string) *templates {
	//TODO 具体处理逻辑省略
	store := &PotentScreen{
		//
	}
	return &templates{PotentScreen: store}
}

//--NO4
func buildSelectCardsVideo(req *item, templateId string) *templates {
	//TODO 具体处理逻辑省略
	cardsVideo := &CardsVideo{
		VideoUrl: "",
		PicUrl:   "",
		Mark:     "",
		Title:    "",
		Content:  "",
	}
	return &templates{CardsVideo: cardsVideo}
}

//--NO1
func main() {
	item := &item{
		Title:     "",
		Content:   "",
		Icon:      "",
		IconTitle: "",
		FromLogo:  "",
		From:      "",
	}
	buildTemplates := getBuildTemplates(item, "1-1", PotentSCREEN)
	fmt.Println(buildTemplates)
}

//--NO2
func getBuildTemplates(req *item, templateId string, bizCode SpecialBizCodeName) *templates {
	if parser, ok := buildTemplatesParserMap[bizCode]; ok {
		return parser(req, templateId)
	}
	return nil
}
