package main

import (
	"Douban-TOP/parse"
	"strconv"
	"strings"
)

var BaseUrl string = "https://movie.douban.com/top250?start="

func crawler() {
	var movies []parse.DoubanMovie
	for i := 0; i < 10; i++ {
		num := strconv.Itoa(i * 25)
		// 每页的URL
		url := strings.Join([]string{BaseUrl, num, "&filter="}, "")

		movies = append(movies, parse.ParseFilm(url)...)

	}

}

func main() {

	crawler()

}
