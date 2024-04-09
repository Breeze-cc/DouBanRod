package method

import (
	"douban/model"
	"github.com/go-rod/rod"
	"strings"
	"time"
)

// CollectMovieInfo 采集电影详细信息
func CollectMovieInfo(page *rod.Page) model.Movie {
	var alias, IMDB string
	name := page.MustElement(`span[property="v:itemreviewed"]`).MustText()
	director := page.MustElement(`a[rel="v:directedBy"]`).MustText()
	screenWriter := page.MustElement(`#info > span:nth-child(3) > span.attrs > a`).MustText()
	var actors []string
	for _, actor := range page.MustElements(`a[rel="v:starring"]`) {
		actors = append(actors, strings.TrimSuffix(actor.MustText(), "/"))
	}
	genre := page.MustElement(`span[property="v:genre"]`).MustText()
	// 地区和语言没有单独的标签，排列在span标签后面，使用xPath查找
	span := page.MustElementR("span", "制片国家/地区:")
	region := span.MustElementX(`following-sibling::text()[1]`).MustText()
	span = page.MustElementR("span", "语言:")
	language := span.MustElementX(`following-sibling::text()[1]`).MustText()
	var releaseDate []string
	for _, element := range page.MustElements(`span[property="v:initialReleaseDate"]`) {
		releaseDate = append(releaseDate, element.MustText())
	}
	runtime := page.MustElement(`span[property="v:runtime"]`).MustText()
	span, err := page.Timeout(time.Second).ElementR("span", "又名:")
	if err != nil {
		alias = ""
	} else {
		alias = span.MustElementX(`following-sibling::text()[1]`).MustText()
	}
	span = page.MustElementR("span", "IMDb:")
	x, err := page.Timeout(time.Second).ElementX(`"//a[contains(@href, 'imdb')]"`)
	if err != nil {
		IMDB = span.MustElementX(`following-sibling::text()[1]`).MustText()
	} else {
		IMDB = x.MustText()
	}

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
		Region:       region,
		Language:     language,
		Length:       runtime,
		Alias:        alias,
		IMDB:         IMDB,
	}
	return movie
}
