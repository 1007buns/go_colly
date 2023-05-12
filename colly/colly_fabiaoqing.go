// package main

// import (
// 	"fmt"
// 	"strconv"
// 	"time"

// 	"github.com/gocolly/colly"
// )

// func main() {

// 	userAgent := "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.0.0 Safari/537.36"

// 	c := colly.NewCollector(colly.UserAgent(userAgent), colly.DetectCharset())

// 	for i := 1; i <= 10; i++ {

// 		s := strconv.Itoa(i)

// 		c.OnHTML("a[href]", func(e *colly.HTMLElement) {
// 			e.Request.Visit(e.Attr("href"))
// 		})

// 		err := c.Visit("https://www.fabiaoqing.com/tag/detail/id/212/page/" + s + ".html")

// 		if err != nil {
// 			fmt.Println(err.Error())
// 		}

// 		fmt.Println("https://www.fabiaoqing.com/tag/detail/id/212/page/" + s + ".html")
// 		time.Sleep(300 * time.Millisecond)

// 	}

// }
