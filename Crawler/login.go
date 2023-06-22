package Crawler

import (
	"errors"
	"github.com/tebeka/selenium"
	"time"
)

const (
	SiteLogin = "https://pje.tjmg.jus.br/pje/login.seam"
	Frame     = "/html/body/div[3]/iframe"
	XpathUser = "/html/body/div/div/div/div/div[1]/div/form/div/div/div[2]/div[1]/div/input"
	XpathPass = "/html/body/div/div/div/div/div[1]/div/form/div/div/div[2]/div[2]/div/input"
	XpathBtt  = "/html/body/div/div/div/div/div[1]/div/form/div/div/div[2]/div[3]/div/input[1]"
)

func Login(driver selenium.WebDriver, login string, password string) error {
	err := driver.Get(SiteLogin)
	if err != nil {
		return errors.New("url unavailable")
	}

	waitXpath(driver, Frame)

	err = driver.SwitchFrame("ssoFrame")
	if err != nil {
		return errors.New("frame not found")
	}

	userName, err := driver.FindElement(selenium.ByXPATH, XpathUser)
	if err != nil {
		return errors.New("xpathUser not found")
	}

	psw, err := driver.FindElement(selenium.ByXPATH, XpathPass)
	if err != nil {
		return errors.New("xpathPass not found")
	}

	btt, err := driver.FindElement(selenium.ByXPATH, XpathBtt)
	if err != nil {
		return errors.New("xpathBtt not found")
	}

	err = userName.SendKeys(login)
	if err != nil {
		return errors.New("could not send login parameter")
	}

	err = psw.SendKeys(password)
	if err != nil {
		return errors.New("could not send password parameter")
	}

	err = btt.Click()
	if err != nil {
		return errors.New("could not click on login button")
	}

	time.Sleep(2 * time.Second)
	err = driver.Refresh()
	if err != nil {
		return errors.New("could not refresh page")
	}
	time.Sleep(2 * time.Second)

	return nil
}

func waitXpath(driver selenium.WebDriver, XPATH string) {
	for {
		elem, _ := driver.FindElements(selenium.ByXPATH, XPATH)
		if len(elem) > 0 {
			break
		}
	}
}
