package acuser

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"github.com/cotora/ac-profile/util"
	"unicode/utf8"
	//"image"
)

type ACuser struct {
	name        string
	country     string
	birth       string
	twitter     string
	affiliation string
	rank        string
	rating      string
	highest     string
	matches     string
	competed    string
	maxLen      int
	hflag       bool
}

func (self *ACuser)Init(name string, hflag bool) error {

	self.maxLen = 0
	self.hflag = hflag

	webPage := "https://atcoder.jp/users/" + name

	if hflag {
		webPage += "?contestType=heuristic"
	}

	resp, err := http.Get(webPage)

	if err != nil {
		return fmt.Errorf("[Error] : failed to get information")
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		if resp.StatusCode == 404 {
			return fmt.Errorf("[Error] : user not found")
		} else {
			return fmt.Errorf("[Error] : failed to fetch data")
		}
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return fmt.Errorf("[Error] : failed to load html")
	}

	self.name = doc.Find("div h3 .username span").Text()
	rex:=regexp.MustCompile("[0-9]+")

	//get profile information
	doc.Find("div .dl-table tr").Each(func(i int, s *goquery.Selection) {
		title := s.Find("th").Text()
		title = strings.TrimSpace(title)
		item := strings.TrimSpace(s.Find("td").Text())
		item = strings.Replace(item, "\t", "", -1)
		item = strings.Replace(item, "\n", "", -1)
		switch title{
			case "Country/Region":
				self.country = item
				self.maxLen = util.Max(self.maxLen, utf8.RuneCountInString(title))
			case "Birth Year":
				self.birth = item
				self.maxLen = util.Max(self.maxLen, utf8.RuneCountInString(title))
			case "Twitter ID":
				self.twitter = item
				self.maxLen = util.Max(self.maxLen, utf8.RuneCountInString(title))					
			case "Affiliation":
				self.affiliation = item
				self.maxLen = util.Max(self.maxLen, utf8.RuneCountInString(title))
			case "Rank":
				self.rank = item
				self.maxLen = util.Max(self.maxLen, utf8.RuneCountInString(title))
			case "Rating":
				item = rex.FindString(item)
				self.rating = item
				self.maxLen = util.Max(self.maxLen, utf8.RuneCountInString(title))
			case "Highest Rating":
				self.highest = rex.FindString(item)
				self.maxLen = util.Max(self.maxLen, utf8.RuneCountInString(title))
			case "Rated Matches":
				self.matches = item
				self.maxLen = util.Max(self.maxLen, utf8.RuneCountInString(title))
			case "Last Competed":
				self.competed = item
				self.maxLen = util.Max(self.maxLen, utf8.RuneCountInString(title))
		}
	})
	return nil
}

func (self ACuser)printItems(title string,item string){
	if item!=""{
		fmt.Println(util.PaddingSpace(title, self.maxLen) + " : " + item)
	}
}


func (self ACuser) PrintInformation() {

	if !self.hflag {
		fmt.Println("[Algorithm]")
	} else {
		fmt.Println("[Heuristic]")
	}

	var rate int

	if self.rating != "" {
		var err error
		rate, err = strconv.Atoi(self.rating)
		if err != nil {
			log.Fatalf("failed to string to int:%s", err)
		}
		self.printItems("User",util.RatingColor(rate, self.name))
	} else {
		self.printItems("User",self.name)
	}

	self.printItems("Country/Region",self.country)
	self.printItems("Birth Year",self.birth)
	self.printItems("Twitter ID",self.twitter)
	self.printItems("Affiliation",self.affiliation)
	self.printItems("Rank",self.rank)
	self.printItems("Rating",util.RatingColor(rate, self.rating))

	if self.highest != "" {
		hrate, err := strconv.Atoi(self.highest)
		if err != nil {
			log.Fatalf("failed to string to int:%s", err)
		}
		self.printItems("Highest Rating",util.RatingColor(hrate, self.highest))
	}

	self.printItems("Rated Matches",self.matches)
	self.printItems("Last Competed",self.competed)
}