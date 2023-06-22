package Crawler

import (
	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
	"strings"
)

type Poles struct {
	LawsuitNumber string
	ActivePole    []string
	PassivePole   []string
	OtherPole     []string
}

func GetLawsuitPoles(htmlPgSrc *html.Node, lawsuitNumber string) (Poles, error) {
	var activePole []string
	var passivePole []string
	var otherPole []string

	totalActive := htmlquery.Find(htmlPgSrc, "//*[@id=\"poloAtivo\"]/table/tbody/tr/td")
	for _, node1 := range totalActive {
		activePole = append(activePole, strings.ReplaceAll(strings.TrimSpace(htmlquery.InnerText(htmlquery.FindOne(node1, "/span/span"))), "\n", ""))
	}

	totalPassive := htmlquery.Find(htmlPgSrc, "//*[@id=\"poloPassivo\"]/table/tbody/tr/td")
	for _, node2 := range totalPassive {
		passivePole = append(passivePole, strings.ReplaceAll(strings.TrimSpace(htmlquery.InnerText(htmlquery.FindOne(node2, "/span/span"))), "\n", ""))
	}

	totalOther := htmlquery.Find(htmlPgSrc, "//*[@id=\"outrosInteressados\"]/table/tbody/tr/td")
	for _, node3 := range totalOther {
		otherPole = append(otherPole, strings.ReplaceAll(strings.TrimSpace(htmlquery.InnerText(htmlquery.FindOne(node3, "/span/span"))), "\n", ""))
	}

	return Poles{
		LawsuitNumber: lawsuitNumber,
		ActivePole:    activePole,
		PassivePole:   passivePole,
		OtherPole:     otherPole,
	}, nil

}
