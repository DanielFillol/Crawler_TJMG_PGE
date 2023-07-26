package main

import (
	"Crawler_TJMG_PGE/CSV"
	"Crawler_TJMG_PGE/Crawler"
	"fmt"
	"log"
	"strconv"
	"time"
)

type SearchLawsuits struct {
	LawsuitNumber  string
	NameLawsuit    string
	DocumentNumber []string
}

const (
	Login    = "YOUR_LOGIN"
	Password = "YOUR_LOGIN"
)

func main() {
	start1 := time.Now()

	driver, err := Crawler.SeleniumWebDriver()
	if err != nil {
		fmt.Println(err)
	}

	//defer driver.Close()

	err = Crawler.Login(driver, Login, Password)
	if err != nil {
		fmt.Println(err)
	}

	var suits []Crawler.Lawsuit
	for i, lawsuitNumber := range lawsuitNumbers {
		start2 := time.Now()

		lawsuit, err := Crawler.Craw(driver, lawsuitNumber.LawsuitNumber)
		if err != nil {
			log.Println(err)
		}
		suits = append(suits, lawsuit)

		t1 := time.Since(start1).String()
		t2 := time.Since(start2).String()
		m := int(time.Since(start1).Seconds()) / (i + 1)
		r := int(time.Since(start1).Seconds()) % (i + 1)
		md := strconv.Itoa(m) + "." + strconv.Itoa(r)
		fmt.Printf("processado %v | tempo: %v%v | total: %v%v | média: %vs \n", i+1, t2[0:4], t2[len(t2)-1:], t1[0:4], t1[len(t1)-1:], md)
	}

	err = CSV.WriteCSV(suits)
	if err != nil {
		fmt.Println(err)
	}
}

var lawsuitNumbers = []SearchLawsuits{
	{LawsuitNumber: "5004331-30.2020.8.13.0024", NameLawsuit: "", DocumentNumber: []string{"073.496.146-40", "07349614640"}},
}
