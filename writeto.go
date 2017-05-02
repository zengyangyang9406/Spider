package main

import (
	"io/ioutil"
	"log"

	"github.com/PuerkitoBio/goquery"
)

type Qiushi struct {
	name    string
	content string
}

var content string

func Writetofile(data []Qiushi) string {
	for _, state := range data {

		content = content + state.name + state.content

	}
	s := []byte(content)
	err := ioutil.WriteFile("/Users/RNTD009/Desktop/data.txt", s, 0644)

	if err != nil {
		panic(err)
		return "no"
	}
	return "yes"
}

func GetJokes() []Qiushi {
	doc, err := goquery.NewDocument("http://www.qiushibaike.com")
	if err != nil {
		log.Fatal(err)
	}

	// doc.Find("#content-left .article .author a img").Each(func(i int, s *goquery.Selection) {
	// 	qiushi.img = s.Attr("src")
	// })

	var slice []string
	var slice1 []string
	var slice2 []Qiushi

	doc.Find("#content-left .article .author a h2").Each(func(m int, d *goquery.Selection) {

		slice = append(slice, d.Text())

	})

	doc.Find(".content").Each(func(h int, f *goquery.Selection) {

		slice1 = append(slice1, f.Text())

	})

	for i, state := range slice1 {

		slice2 = append(slice2, Qiushi{state, slice1[i]})

	}

	return slice2
}
