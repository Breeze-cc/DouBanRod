package method

import (
	"douban/model"
	"fmt"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/utils"
)

func GETNewMovies(page *rod.Page) {
	// 存放电影信息的切片
	var movieDetails []model.Movie
	// 导航至豆瓣电影
	page.MustNavigate(`https://movie.douban.com/chart`).MustWaitStable()
	// 获取所有页面上的电影
	movies := page.MustElements(`#content > div > div.article > div > div > table`)
	/*
		由于需要点击超链接,每次返回后movies数据会刷新，因此需要在每次循环中重新获取
		否则会发生栈内存溢出panic
	*/
	for i := 0; i < len(movies); i++ {
		// 重新获取movies
		movies := page.MustElements(`#content > div > div.article > div > div > table`)
		// 指定一个movie
		movie := movies[i]
		// 点进该电影的页面
		movie.MustElement(`a`).MustClick()
		// 通过CollectMovieInfo方法获取电影详细信息,返回model.Movie类型
		movieDetail := CollectMovieInfo(page)
		// 将刚才得到的movie添加到切片中
		movieDetails = append(movieDetails, movieDetail)
		// 页面后退，回到排行榜页面
		page.MustNavigateBack()
		utils.Sleep(1)
	}

	for i, movie := range movieDetails {
		fmt.Println(i, movie)
	}

}
