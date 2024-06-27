package topwebs

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"strings"
)

type Site struct {
	Position   string `json:"id"`
	Website    string `json:"name"`
	Change     string `json:"change"`
	Visits     string `json:"visits"`
	AvgPages   string `json:"avg_pages"`
	BounceRate string `json:"bounce_rate"`
}

//	create site objects, fill their data, store in Sites.
//
//	to create site objects:
//	iterate through the scanner, in the order of:
//		id, name, change, visits, avg_pages, bounce_rate
//

var jsonBody = ""
var Sites []Site
var finalJson = ""

func main() {
	fmt.Println(topTen())
}

func topTen() string {
	//scrape the website for the correct data
	ScrapeUrl := "https://www.semrush.com/website/top/"

	c := colly.NewCollector(colly.AllowedDomains("www.semrush.com", "semrush.com"))

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("on request Visiting", r.URL.String(), "\n")
	})

	c.OnHTML("table.table_table__Rggo8 td", func(h *colly.HTMLElement) {
		//add data to string or structure.
		updateJson(h.Text)
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)

		fmt.Println(ParseBody(jsonBody))
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Response from ", r.Request.URL.String())
	})

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Error encountered: ", err.Error())
	})

	c.Visit(ScrapeUrl)
	fmt.Println(finalJson, "ALRIGHT. ")
	fmt.Println("yup that's it.")
	return finalJson
}

func updateJson(line string) {
	jsonBody += line
	jsonBody += "\n"
}

func ParseBody(body string) string {
	scanner := bufio.NewScanner(strings.NewReader(body))

	objCount := 0
	propCount := 0
	currentObj := 0
	tempObj := Site{} //update temp obj and whenever objCount iterates, save and set to a new object.
	for scanner.Scan() {
		//keep track of what object we are creating: objCount
		//keep track of what properties we are assigning values to: propCount
		//How?
		//everytime propCount == 6, objCount++
		//everytime propCount == 6, propCount = 0

		if propCount == 6 {
			objCount++
			propCount = 1
		} else {
			propCount++
		}

		//if current object isn't
		if currentObj != objCount {
			//move to a different object to populate with data
			//do something then set to object count
			if objCount < 11 {
				Sites = append(Sites, tempObj)
			}
			currentObj = objCount
		}

		switch propCount {
		case 1:
			//assign the id to the id.id
			tempObj.Position = scanner.Text()
		case 2:
			tempObj.Website = scanner.Text()
		case 3:
			tempObj.Change = scanner.Text()
		case 4:
			tempObj.Visits = scanner.Text()
		case 5:
			tempObj.AvgPages = scanner.Text()
		case 6:
			tempObj.BounceRate = scanner.Text()
		default:

		}

	}

	fmt.Println("-------------------------------------S    I    T     E    S----------------------------------")

	j, _ := json.MarshalIndent(Sites, "", "  ")
	//log.Println(string(j))
	//fmt.Println(string(j))

	if err := scanner.Err(); err != nil {
		fmt.Printf("error occurred: %v\n", err)
	}
	finalJson = string(j)
	return string(j)
}

func toptenwebs() string {
	return finalJson
}
