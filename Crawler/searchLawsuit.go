package Crawler

import (
	"errors"
	"github.com/Darklabel91/CNJ_Validate/CNJ"
	"github.com/antchfx/htmlquery"
	"github.com/tebeka/selenium"
	"golang.org/x/net/html"
	"log"
	"strings"
	"time"
)

const (
	xpathSearchBtt   = "//*[@id=\"fPP:searchProcessos\"]"
	xpathNumber1     = "//*[@id=\"fPP:numeroProcesso:numeroSequencial\"]"
	xpathNumber2     = "//*[@id=\"fPP:numeroProcesso:numeroDigitoVerificador\"]"
	xpathNumber3     = "//*[@id=\"fPP:numeroProcesso:Ano\"]"
	xpathNumber4     = "//*[@id=\"fPP:numeroProcesso:NumeroOrgaoJustica\"]"
	xpathReturn      = "//*[@id=\"fPP:processosTable:tb\"]/tr"
	xpathLawsuitLink = "/html/body/div[5]/div/div/div/div[2]/form/div[2]/div/table/tbody/tr/td[2]/a"
)

type lawsuitNumber struct {
	number1 string
	number2 string
	number3 string
	number4 string
}

func SearchLawsuit(driver selenium.WebDriver, searchLink string, lawsuit string) (*html.Node, error) {
	url, err := driver.CurrentURL()
	if err != nil {
		return nil, errors.New("error getting current url")
	}

	if url != searchLink {
		err := driver.Get(searchLink)
		if err != nil {
			return nil, errors.New("url unavailable")
		}
	}

	n1, err := driver.FindElement(selenium.ByXPATH, xpathNumber1)
	if err != nil {
		return nil, errors.New("could not find xpathNumber1")
	}

	n2, err := driver.FindElement(selenium.ByXPATH, xpathNumber2)
	if err != nil {
		return nil, errors.New("could not find xpathNumber2")
	}

	n3, err := driver.FindElement(selenium.ByXPATH, xpathNumber3)
	if err != nil {
		return nil, errors.New("could not find xpathNumber3")
	}

	n4, err := driver.FindElement(selenium.ByXPATH, xpathNumber4)
	if err != nil {
		return nil, errors.New("could not find xpathNumber4")
	}

	number, err := formatNumber(lawsuit)
	if err != nil {
		log.Println(err, number)
		return nil, err
	}

	err = n1.Clear()
	err = n2.Clear()
	err = n3.Clear()
	err = n4.Clear()
	err = n1.SendKeys(number.number1)
	err = n2.SendKeys(number.number2)
	err = n3.SendKeys(number.number3)
	err = n4.SendKeys(number.number4)
	if err != nil {
		return nil, errors.New("could not send lawsuit number as search parameter")
	}

	btt, err := driver.FindElement(selenium.ByXPATH, xpathSearchBtt)
	if err != nil {
		return nil, errors.New("could not find xpathSearchBtt")
	}

	err = btt.Click()
	if err != nil {
		return nil, errors.New("could not click on search button")
	}

	//there is a delay of 1.5s to load the page
	time.Sleep(1500 * time.Millisecond)

	exist, err := existLawsuit(driver, xpathReturn)
	if err != nil {
		return nil, errors.New("err:" + err.Error())
	}
	if exist != true {
		log.Println("could not find lawsuit")
		return nil, errors.New("could not find lawsuit")
	}

	lawsuitLink, err := driver.FindElement(selenium.ByXPATH, xpathLawsuitLink)
	if err != nil {
		return nil, errors.New("could not find  lawsuit link")
	}

	//click on lawsuit
	err = lawsuitLink.Click()
	if err != nil {
		return nil, errors.New("could not click on lawsuit link")
	}

	//there is a delay of 1.5s to load the page
	time.Sleep(1500 * time.Millisecond)
	driver.AcceptAlert()
	//there is a delay of 0.5s to load the page
	time.Sleep(500 * time.Millisecond)

	wd, err := driver.WindowHandles()
	if err != nil {
		return nil, errors.New("could not get window handles")
	}

	err = driver.SwitchWindow(wd[1])
	if err != nil {
		return nil, errors.New("could not switch window")
	}

	pageSource, err := driver.PageSource()
	if err != nil {
		return nil, errors.New("could not get page source")
	}

	htmlPgSrc, err := htmlquery.Parse(strings.NewReader(pageSource))
	if err != nil {
		return nil, errors.New("could not convert string to node html")
	}

	err = driver.SwitchWindow(wd[0])
	if err != nil {
		return nil, errors.New("could not switch window")
	}

	err = driver.CloseWindow(wd[1])
	if err != nil {
		return nil, errors.New("could not close window")
	}

	return htmlPgSrc, nil
}

func formatNumber(lawsuit string) (lawsuitNumber, error) {
	decomposedCNJ, _ := CNJ.DecomposeCNJ(lawsuit)

	return lawsuitNumber{
		number1: decomposedCNJ.LawsuitNumber,
		number2: decomposedCNJ.VerifyingDigit,
		number3: decomposedCNJ.ProtocolYear,
		number4: decomposedCNJ.SourceUnit,
	}, nil

}

func existLawsuit(driver selenium.WebDriver, xpathReturn string) (bool, error) {
	elements, err := driver.FindElements(selenium.ByXPATH, xpathReturn)
	if err != nil {
		return false, err
	}

	if len(elements) != 0 {
		return true, nil
	}

	return false, nil
}
