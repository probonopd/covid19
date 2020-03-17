package main

import (
	"fmt"
	"github.com/antchfx/htmlquery"
	"github.com/antchfx/jsonquery"
	"strings"
)

// Get number of Corona infections per Landkreis in Baden-Württemberg
// from Schwäbische Zeitung. It may be interesting to compare those numbers
// over time with the figures from sozialministerium.baden-wuerttemberg.de

func main() {

	// Infogram ID: 1hxr4z1vydoq2yo
	// https://schwaebische.infogram.com/falle-corona-baden-wurttemberg-1hxr4z1vydoq2yo
	// or https://e.infogram.com/9f1a3619-b781-43e4-b5b7-7f3853e9d252?src=embed
	// These links are working as of 17.03.2020

	url := "https://e.infogram.com/9f1a3619-b781-43e4-b5b7-7f3853e9d252?src=embed"

	doc, err := htmlquery.LoadURL(url)
	if err != nil {
		fmt.Println(err)
	}

	script := htmlquery.FindOne(doc, "/html/body/script[1]")
	jsonString := strings.Replace(string(script.FirstChild.Data), "window.infographicData=", "", -1)
	jsonString = strings.Replace(string(jsonString), "}};", "}}", -1)
	// fmt.Println(jsonString)

	jdoc, err := jsonquery.Parse(strings.NewReader(jsonString))
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(jdoc)
	data := jsonquery.FindOne(jdoc,"//data[1]")
	// fmt.Println(data.FirstChild.FirstChild.FirstChild.FirstChild.Data) // Alb-Donau-Kreis
	interestingData := data.FirstChild
	// fmt.Println(interestingData.InnerText()) // Has all the data we ware interested in
	nodes := interestingData.ChildNodes()
	var output string
	for _, node := range nodes {
		// fmt.Println(node.InnerText())
		landkreis := node.FirstChild.InnerText()
		cases := node.FirstChild.NextSibling.InnerText()
		output = output + landkreis + ";" + cases + "\r\n"
	}
	fmt.Println(output)

	/*
	file, err := os.Create("corona.csv")
	defer file.Close()
	if err != nil {
		fmt.Println(err)
	} else {
		file.WriteString(output)
	}
	*/
}
