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

	
}
