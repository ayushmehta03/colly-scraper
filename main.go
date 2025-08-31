package main

import(
	"fmt"
	"log"
	"strings"
	"encoding/json"
	"flag"
	"github.com/gocolly/colly"
)
type star struct{
	Name string
	Photo string
	JobTitle string
	BirthDate string
	Bio string
	TopMovies []movie

}

type movie struct{
	title string
	Year string



}

func main(){
	month :=flag.Int("month",1,"Month to fetch birthdays for")
	day :=flag.Int("day",1,"day to fetch birthdays for")
	flag.Parse()
	crawl(*month,*day)
}


func crawl(month int , day int){
	c:=colly.NewCollector(
		colly.AllowedDomains("imdb.com","www.imdb.com")
	)

	infoCollector:=c.Clone()


	c.OnHTML(".mode-detail",func(e *colly.HTMLElement)){
		profileUrl:=e.ChildAttr("div.lister-item-image > a","href")
		profileUrl=e.Request.AbsoluteURL(profileUrl)
		infoCollector.Visit(profileUrl)
	}

	c.OnHTML("a.lister-page-next",func(e *colly.HTMLElement){
		nextPage:=e.Request.AbsoluteURL(e.Attr("href"))
		c.Visit(nextPage)
	})

	infoCollector.OnHTML("#")


	startUrl:=fmt.Sprintf("https://www.imdb.com/search/name/?birth_monthday=%d-%d",month,day)
	c.Visit(startUrl)




}
