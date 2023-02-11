package main

import (
	"fmt"
	"log"
	"net/http"
	//"os"
	"flag"
	"strconv"
	"strings"
	"github.com/PuerkitoBio/goquery"
	"regexp"
	"unicode/utf8"
	//"image"
)

func ratingColor(rating int,s string) string {
	if rating>=2800{
		return "\033[38;5;196m"+s+"\033[m"
	} else if rating>=2400{
		return "\033[38;5;208m"+s+"\033[m"
	} else if rating>=2000{
		return "\033[38;5;226m"+s+"\033[m"
	} else if rating>=1600{
		return "\033[38;5;27m"+s+"\033[m"
	} else if rating>=1200{
		return "\033[38;5;123m"+s+"\033[m"
	} else if rating>=800{
		return "\033[38;5;76m"+s+"\033[m"
	} else if rating>=400{
		return "\033[38;5;94m"+s+"\033[m"
	} else{
		return "\033[38;5;248m"+s+"\033[m"
	}
}

func max(x int,y int) int {
	if x>y{
		return x
	} else{
		return y
	}
}

func paddingSpace(s string,n int) string{
	return s+strings.Repeat(" ",n-utf8.RuneCountInString(s))
}

func main(){
	var(
		//u=flag.String("u","tourist","user name flag")
		t=flag.String("t","a","rating type flag")
	)
	flag.Parse()
	if len(flag.Args())==0{
		fmt.Println("fatal error : no input user name")
		return
	}
	webPage:="https://atcoder.jp/users/"+flag.Arg(0)
	if *t=="h"{
		webPage+="?contestType=heuristic"
	}
	resp,err:=http.Get(webPage)
	if err!=nil{
		fmt.Println("failed to get html")
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200{
		fmt.Println("failed to fetch data")
		return
	}

	doc,err:=goquery.NewDocumentFromReader(resp.Body)
	if err!=nil{
		log.Printf("failed to load html:%s",err)
	}

	user:=flag.Arg(0)
	country:=""
	birth:=""
	twitter:=""
	affiliation:=""
	rank:=""
	rating:=""

	/*
	imgUrl,exists:=doc.Find(".avatar").Attr("src")
	if exists{
		if imgUrl=="//img.atcoder.jp/assets/icon/avatar.png"{imgUrl="https:"+imgUrl}
		//fmt.Println(imgUrl)
		imgResp,err:=http.Get(imgUrl)
		if err!=nil{
			fmt.Println("failed to get icon-img")
		}
		defer imgResp.Body.Close()
		
	}
	*/
	maxLen:=0
	doc.Find("div .dl-table tr").Each(func(i int,s *goquery.Selection){
			title:=s.Find("th").Text()
			item:=strings.TrimSpace(s.Find("td").Text())
			item=strings.Replace(item,"\t","",-1)
			item=strings.Replace(item,"\n","",-1)
			if title=="Country/Region"{
				country=item
				maxLen=max(maxLen,utf8.RuneCountInString(title))
			} else if title=="Birth Year"{
				birth=item
				maxLen=max(maxLen,utf8.RuneCountInString(title))
			} else if title=="Twitter ID"{
				twitter=item
				maxLen=max(maxLen,utf8.RuneCountInString(title))
			} else if title=="Affiliation"{
				affiliation=item
				maxLen=max(maxLen,utf8.RuneCountInString(title))
			} else if title=="Rank"{
				rank=item
				maxLen=max(maxLen,utf8.RuneCountInString(title))
			} else if title=="Rating"{
				rex:=regexp.MustCompile("[0-9]+")
				item=rex.FindString(item)
				rating=item
				maxLen=max(maxLen,utf8.RuneCountInString(title))
			}
	})
	if *t=="a"{
		fmt.Println("[Algorithm]")
	} else{
		fmt.Println("[Heuristic]")
	}
	if rating!=""{
		rate,err:=strconv.Atoi(rating)
		if err!=nil{
			log.Fatalf("failed to string to int:%s",err)
		}
		fmt.Println(paddingSpace("User",maxLen)+" : "+ratingColor(rate,user))
	} else{
		fmt.Println(paddingSpace("User",maxLen)+" : "+user)
	}
	if country!=""{
		fmt.Println(paddingSpace("Country/Region",maxLen)+" : "+country)
	}
	if birth!=""{
		fmt.Println(paddingSpace("Birth Year",maxLen)+" : "+birth)
	}
	if twitter!=""{
		fmt.Println(paddingSpace("Twitter ID",maxLen)+" : "+twitter)
	}
	if affiliation!=""{
		fmt.Println(paddingSpace("Affiliation",maxLen)+" : "+affiliation)
	}
	if rank!=""{
		fmt.Println(paddingSpace("Rank",maxLen)+" : "+rank)
	}
	if rating!=""{
		rate,err:=strconv.Atoi(rating)
		if err!=nil{
			log.Fatalf("failed to string to int:%s",err)
		}
		fmt.Println(paddingSpace("Rating",maxLen)+" : "+ratingColor(rate,rating))
	}
}