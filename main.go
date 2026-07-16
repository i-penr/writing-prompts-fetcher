package main

import (
	"os"
	"fmt"
	"time"
	"strings"
	"github.com/gocolly/colly"
)

const SOURCE_URL = "https://www.reddit.com/r/WritingPrompts/best";

func main() {
	found := false;
	targetPath := os.Args[1:][0]
	fileName := time.Now().Format("02-01-2006")

	f, f_err := os.Create(fmt.Sprintf("%s/%s.md", targetPath, fileName))

	if f_err != nil {
		panic(f_err)
	}

	c := colly.NewCollector(
		colly.AllowedDomains("www.reddit.com"),
	);

	c.OnHTML("article", func(e *colly.HTMLElement) {
		var wpURL string = e.ChildAttr("a.absolute", "href")
		var wp string = e.ChildText("a.absolute")

		if (wp != "" && !found) {
			found = true

			wp = "> " + wpURL + "\n\n```\n" + strings.Trim(strings.Replace(wp, `[WP]`, "", -1), " ") + "\n```"

			f.WriteString(wp)
			fmt.Printf("File created successfully at %s/%s.md\n", targetPath, fileName)
			f.Close()
		}
		
	})

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong: ", err)
	})
	
	c_err := c.Visit(SOURCE_URL)

	if c_err != nil {
		panic(c_err)
	}
}
