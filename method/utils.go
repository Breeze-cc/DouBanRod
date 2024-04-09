package method

import (
	"douban/model"
	"fmt"
	"github.com/go-rod/rod"
	"github.com/xuri/excelize/v2"
	"reflect"
	"strings"
)

func writeToExcel(data []model.Movie) {
	// 创建文件
	file := excelize.NewFile()
	sheetName := "Sheet1"
	// 创建工作表
	index, _ := file.NewSheet(sheetName)

	// 设置标题行
	titles := []string{"名称", "上映时间", "评分", "海报"}
	for i, title := range titles {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		file.SetCellValue("Sheet1", cell, title)
	}
	// 遍历电影切片，写入每部电影的数据
	for i, movie := range data {
		v := reflect.ValueOf(movie)
		for j := 0; j < v.NumField(); j++ {
			cell, _ := excelize.CoordinatesToCellName(j+1, i+2) // 行号从2开始，因为第1行已经是列标题了
			file.SetCellValue(sheetName, cell, v.Field(j).Interface())
		}
	}
	// 设置默认打开的工作表
	file.SetActiveSheet(index)
	// 保存文件
	if err := file.SaveAs("resource/movies.xlsx"); err != nil {
		println(err.Error())
	}
}

// 采集电影详细信息
func CollectMovieInfo(page *rod.Page) model.Movie {
	name := page.MustElement(`span[property="v:itemreviewed"]`).MustText()
	director := page.MustElement(`a[rel="v:directedBy"]`).MustText()
	screenWriter := page.MustElement(`#info > span:nth-child(3) > span.attrs > a`).MustText()
	actors := []string{}
	for _, actor := range page.MustElements(`a[rel="v:starring"]`) {
		actors = append(actors, strings.TrimSuffix(actor.MustText(), "/"))
	}
	genre := page.MustElement(`span[property="v:genre"]`).MustText()
	// 地区和语言没有单独的标签，排列在span标签后面
	//region := strings.TrimSpace(page.MustElement(`#info > span:nth-child(10)`).MustParent().MustElementX(`./text()`).MustText())
	//language := strings.TrimSpace(page.MustElement(`#info > span:nth-child(12)`).MustParent().MustElementX(`./text()`).MustText())
	var releaseDate []string
	for _, element := range page.MustElements(`span[property="v:initialReleaseDate"]`) {
		releaseDate = append(releaseDate, element.MustText())
	}
	runtime := page.MustElement(`span[property="v:runtime"]`).MustText()
	//alias := strings.TrimSpace(page.MustElement(`#info > span:nth-child(21)`).MustParent().MustElementX(`./text()`).MustText())
	//IMDB := strings.TrimSpace(page.MustElement(`#info > span:nth-child(22)`).MustParent().MustElementX(`./text()`).MustText())
	rate := page.MustElement(`strong[property="v:average"]`).MustText()
	movie := model.Movie{
		MovieName:    name,
		ReleaseDate:  strings.Join(releaseDate, ";"),
		Rate:         rate,
		CoverPath:    "",
		Director:     director,
		ScreenWriter: screenWriter,
		Actors:       actors,
		Type:         genre,
		//Region:       region,
		//Language:     language,
		Length: runtime,
		//Alias:        alias,
		//IMDB:         IMDB,
	}
	v := reflect.ValueOf(movie)
	for i := 0; i < v.NumField(); i++ {
		fmt.Printf("%s %s\n", v.Type().Field(i).Name, v.Field(i).Interface())
	}
	return movie
}
