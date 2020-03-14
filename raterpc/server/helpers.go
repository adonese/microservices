package main

import (
	"net/http"
	"strings"

	"golang.org/x/net/html"
	"log"
)

//extract extracts links of provided URL
func extract(domain string) []string {
	var links []string

	res, err := http.Get(domain)
	if err != nil {
		return nil
	}
	defer res.Body.Close()
	doc, err := html.Parse(res.Body)
	if err != nil {
		return nil
	}
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "strong" {
			// for _, a := range n.Attr {
			// 	if a.Key == "span" {
			// 		fmt.Printf("The value we found is: %v", a)
			// 	}
			// }

			// if n.FirstChild.Data == "سعر الدولار الأمريكي" {
			// 	fmt.Printf("The value we want is: %v\n%v", n.FirstChild.NextSibling.Data, n.FirstChild.Data)
			// }
			links = append(links, n.FirstChild.Data)
			// fmt.Printf("The values are: %#v\n", n.FirstChild.Data)
			// for _, s := range n.Attr {
			// 	fmt.Printf("The value is :%v\n", s.Val)
			// }

		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	//
	return links
}

func getUSD(links []string) (bool, string) {
	for i, v := range links {
		log.Printf("the current string is: %v\n", v)
		if v == "الدولار الامريكي" || strings.Contains(v, "دولار") || strings.Contains(v, "دولار&nbsp;الامريكي") {
			usd := strings.Split(links[i+1], " ")
			// log.Printf("the usd from getUSD is: %v\n", usd)
			return true, usd[0]
		}
	}
	log.Printf("why are not we here?\n")
	return false, ""
}
