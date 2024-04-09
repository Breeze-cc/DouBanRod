package method

import (
	"douban/model"
	"fmt"
	"github.com/go-rod/rod"
)

func GETNewMovies(page *rod.Page) {
	var movieDetails []model.Movie
	page.MustNavigate(`https://movie.douban.com/chart`).MustWaitStable()
	movies := page.MustElements(`#content > div > div.article > div > div > table`)
	for _, movie := range movies {
		movie.MustElement(fmt.Sprintf(`tbody > tr > td:nth-child(2) > div > a`)).MustClick()
		movieDetail := CollectMovieInfo(page)
		movieDetails = append(movieDetails, movieDetail)
	}
	fmt.Printf("%v", movieDetails)
	//writeToExcel(movieDetails)
}
