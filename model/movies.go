package model

type Movie struct {
	MovieName    string   // 影片名称
	ReleaseDate  string   // 影片上映时间
	Rate         string   // 评分
	CoverPath    string   // 封面路径
	Director     string   //导演
	ScreenWriter string   //编剧
	Actors       []string // 主演
	Type         string   // 类型
	Region       string   //制片国家/地区
	Language     string   //语言
	Length       string   //片长
	Alias        string   // 别名
	IMDB         string   // IMDB编号
}
