package method

import (
	"douban/model"
	"fmt"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/utils"
)

func GETNewMovies(page *rod.Page) {
	var movieDetails []model.Movie
	page.MustNavigate(`https://movie.douban.com/chart`).MustWaitStable()
	movies := page.MustElements(`#content > div > div.article > div > div > table`)

	for i := 0; i < len(movies); i++ {
		movies := page.MustElements(`#content > div > div.article > div > div > table`)
		movie := movies[i]
		movie.MustElement(`a`).MustClick()
		movieDetail := CollectMovieInfo(page)
		movieDetails = append(movieDetails, movieDetail)
		page.MustNavigateBack()
		utils.Sleep(1)
	}

	for i, movie := range movieDetails {
		fmt.Println(i, movie)
	}

}
