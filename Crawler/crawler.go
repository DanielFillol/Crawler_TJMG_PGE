package Crawler

import (
	"github.com/tebeka/selenium"
	"log"
)

const InitWebSite = "https://pje.tjmg.jus.br/pje/Processo/ConsultaProcesso/listView.seam"

type Lawsuit struct {
	Cover LawsuitCover
	Poles Poles
}

func Craw(driver selenium.WebDriver, lawsuit string) (Lawsuit, error) {
	searchLink := InitWebSite
	htmlPgSrc, err := SearchLawsuit(driver, searchLink, lawsuit)
	if err != nil {
		log.Println(err)
		return Lawsuit{}, nil
	}

	cover, err := GetLawsuitCover(htmlPgSrc, lawsuit)
	if err != nil {
		log.Println(err)
		return Lawsuit{}, err
	}

	poles, err := GetLawsuitPoles(htmlPgSrc, lawsuit)
	if err != nil {
		log.Println(err)
		return Lawsuit{}, err
	}

	return Lawsuit{
		Cover: cover,
		Poles: poles,
	}, nil
}
