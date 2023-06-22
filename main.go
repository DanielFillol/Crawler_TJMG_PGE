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
	LawsuitNumberMP string
	NameLawsuitMP   string
	DocumentNumber  []string
}

const (
	Login    = "YOUR_LOGIN"
	Password = "YOUR_PASSWORD"
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

		lawsuit, err := Crawler.Craw(driver, lawsuitNumber.LawsuitNumberMP)
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
	{LawsuitNumberMP: "5000919-84.2023.8.13.0251", NameLawsuitMP: "ALEX DA SILVA SANTOS", DocumentNumber: []string{"70575309601"}},
	{LawsuitNumberMP: "5004828-03.2018.8.13.0707", NameLawsuitMP: "BRUNO SANTOS FREITAS", DocumentNumber: []string{"4031451636"}},
	{LawsuitNumberMP: "5002468-27.2020.8.13.0707", NameLawsuitMP: "BRUNO SANTOS FREITAS", DocumentNumber: []string{"4031451636"}},
	{LawsuitNumberMP: "5001901-93.2020.8.13.0707", NameLawsuitMP: "BRUNO SANTOS FREITAS", DocumentNumber: []string{"4031451636"}},
	{LawsuitNumberMP: "5001479-21.2020.8.13.0707", NameLawsuitMP: "BRUNO SANTOS FREITAS", DocumentNumber: []string{"4031451636"}},
	{LawsuitNumberMP: "5000841-85.2020.8.13.0707", NameLawsuitMP: "BRUNO SANTOS FREITAS", DocumentNumber: []string{"4031451636"}},
	{LawsuitNumberMP: "5008127-51.2019.8.13.0707", NameLawsuitMP: "BRUNO SANTOS FREITAS", DocumentNumber: []string{"4031451636"}},
	{LawsuitNumberMP: "5000417-38.2019.8.13.0720", NameLawsuitMP: "BRUNO SANTOS FREITAS", DocumentNumber: []string{"4031451636"}},
	{LawsuitNumberMP: "5007077-53.2020.8.13.0707", NameLawsuitMP: "BRUNO SANTOS FREITAS", DocumentNumber: []string{"4031451636"}},
	{LawsuitNumberMP: "5000469-05.2021.8.13.0707", NameLawsuitMP: "BRUNO SANTOS FREITAS", DocumentNumber: []string{"4031451636"}},
	{LawsuitNumberMP: "5001117-82.2021.8.13.0707", NameLawsuitMP: "BRUNO SANTOS FREITAS", DocumentNumber: []string{"4031451636"}},
	{LawsuitNumberMP: "0279391-45.2013.8.13.0707", NameLawsuitMP: "BRUNO SANTOS FREITAS", DocumentNumber: []string{"4031451636"}},
	{LawsuitNumberMP: "5004435-39.2022.8.13.0707", NameLawsuitMP: "BRUNO SANTOS FREITAS", DocumentNumber: []string{"4031451636"}},
	{LawsuitNumberMP: "5009809-36.2022.8.13.0707", NameLawsuitMP: "BRUNO SANTOS FREITAS", DocumentNumber: []string{"4031451636"}},
	{LawsuitNumberMP: "5001135-06.2021.8.13.0707", NameLawsuitMP: "BRUNO SANTOS FREITAS", DocumentNumber: []string{"4031451636"}},
	{LawsuitNumberMP: "5000671-84.2018.8.13.0707", NameLawsuitMP: "BRUNO SANTOS FREITAS", DocumentNumber: []string{"4031451636"}},
	{LawsuitNumberMP: "5000684-44.2018.8.13.0720", NameLawsuitMP: "BRUNO SANTOS FREITAS", DocumentNumber: []string{"4031451636"}},
	{LawsuitNumberMP: "0469039-38.2009.8.13.0431", NameLawsuitMP: "CARLOS ANTONIO LUCIO", DocumentNumber: []string{"4730558642"}},
	{LawsuitNumberMP: "0198949-76.2014.8.13.0701", NameLawsuitMP: "CARLOS ANTONIO LUCIO", DocumentNumber: []string{"4730558642"}},
	{LawsuitNumberMP: "5000667-18.2019.8.13.0382", NameLawsuitMP: "CARLOS ANTONIO LUCIO", DocumentNumber: []string{"4730558642"}},
	{LawsuitNumberMP: "5000027-41.2023.8.13.0134", NameLawsuitMP: "CARLOS EDUARDO DA SILVA PEREIRA", DocumentNumber: []string{"492732647"}},
	{LawsuitNumberMP: "0062622-75.2015.8.13.0512", NameLawsuitMP: "CLESIO MATOS DE JESUS", DocumentNumber: []string{"10910442606"}},
	{LawsuitNumberMP: "0072003-67.2019.8.13.0480", NameLawsuitMP: "CLESIO MATOS DE JESUS", DocumentNumber: []string{"10910442606"}},
	{LawsuitNumberMP: "5007926-61.2023.8.13.0079", NameLawsuitMP: "DANIEL SOUZA DE ALBUQUERQUE", DocumentNumber: []string{"8576636662"}},
	{LawsuitNumberMP: "5008819-74.2022.8.13.0471", NameLawsuitMP: "DIEGO DA SILVA FARIA", DocumentNumber: []string{"7594368606"}},
	{LawsuitNumberMP: "5000267-89.2022.8.13.0740", NameLawsuitMP: "DIEGO DA SILVA FARIA", DocumentNumber: []string{"7594368606"}},
	{LawsuitNumberMP: "5008072-38.2022.8.13.0338", NameLawsuitMP: "DIEGO DA SILVA FARIA", DocumentNumber: []string{"7594368606"}},
	{LawsuitNumberMP: "0068717-87.2016.8.13.0027", NameLawsuitMP: "DIEGO DA SILVA FARIA", DocumentNumber: []string{"7594368606"}},
	{LawsuitNumberMP: "5000708-43.2018.8.13.0471", NameLawsuitMP: "DIEGO DA SILVA FARIA", DocumentNumber: []string{"7594368606"}},
	{LawsuitNumberMP: "0022364-43.2020.8.13.0384", NameLawsuitMP: "EDMAR MARCELINO DA SILVA", DocumentNumber: []string{"5842262681"}},
	{LawsuitNumberMP: "2721317-16.2013.8.13.0024", NameLawsuitMP: "EDVALDO PINHEIRO CORREA", DocumentNumber: []string{"24727369672"}},
	{LawsuitNumberMP: "0515453-15.2011.8.13.0079", NameLawsuitMP: "EDVALDO PINHEIRO CORREA", DocumentNumber: []string{"24727369672"}},
	{LawsuitNumberMP: "0027850-22.2018.8.13.0079", NameLawsuitMP: "EDVALDO PINHEIRO CORREA", DocumentNumber: []string{"24727369672"}},
	{LawsuitNumberMP: "5000688-77.2018.8.13.0301", NameLawsuitMP: "ELIABE ALVES", DocumentNumber: []string{"5037325676"}},
	{LawsuitNumberMP: "5004780-71.2023.8.13.0027", NameLawsuitMP: "ELIABE ALVES", DocumentNumber: []string{"5037325676"}},
	{LawsuitNumberMP: "5030903-43.2022.8.13.0027", NameLawsuitMP: "FABRICIO HENRIQUE COSTA", DocumentNumber: []string{"8941128641"}},
	{LawsuitNumberMP: "5030373-39.2022.8.13.0027", NameLawsuitMP: "FABRICIO HENRIQUE COSTA", DocumentNumber: []string{"8941128641"}},
	{LawsuitNumberMP: "5011761-91.2022.8.13.0079", NameLawsuitMP: "FABRICIO HENRIQUE COSTA", DocumentNumber: []string{"8941128641"}},
	{LawsuitNumberMP: "5038338-14.2019.8.13.0079", NameLawsuitMP: "FABRICIO HENRIQUE COSTA", DocumentNumber: []string{"8941128641"}},
	{LawsuitNumberMP: "5037324-92.2019.8.13.0079", NameLawsuitMP: "FABRICIO HENRIQUE COSTA", DocumentNumber: []string{"8941128641"}},
	{LawsuitNumberMP: "5037331-84.2019.8.13.0079", NameLawsuitMP: "FABRICIO HENRIQUE COSTA", DocumentNumber: []string{"8941128641"}},
	{LawsuitNumberMP: "5001491-40.2023.8.13.0251", NameLawsuitMP: "FLAVIO ERNANDES TAVARES GOMES DA SILVA", DocumentNumber: []string{"16699588419"}},
	{LawsuitNumberMP: "5004411-21.2022.8.13.0251", NameLawsuitMP: "FLAVIO ERNANDES TAVARES GOMES DA SILVA", DocumentNumber: []string{"16699588419"}},
	{LawsuitNumberMP: "5002130-30.2019.8.13.0338", NameLawsuitMP: "FRANCISCO DE ALMEIDA CAMARGOS", DocumentNumber: []string{"26513625653"}},
	{LawsuitNumberMP: "5009466-80.2022.8.13.0338", NameLawsuitMP: "FRANCISCO DE ALMEIDA CAMARGOS", DocumentNumber: []string{"26513625653"}},
	{LawsuitNumberMP: "5008193-66.2022.8.13.0338", NameLawsuitMP: "FRANCISCO DE ALMEIDA CAMARGOS", DocumentNumber: []string{"26513625653"}},
	{LawsuitNumberMP: "5007434-05.2022.8.13.0338", NameLawsuitMP: "FRANCISCO DE ALMEIDA CAMARGOS", DocumentNumber: []string{"26513625653"}},
	{LawsuitNumberMP: "5003336-41.2023.8.13.0079", NameLawsuitMP: "FRANCISCO LUIZ DOS SANTOS", DocumentNumber: []string{"46172416634"}},
	{LawsuitNumberMP: "0127479-35.2021.8.13.0702", NameLawsuitMP: "FRANCISCO LUIZ DOS SANTOS", DocumentNumber: []string{"46172416634"}},
	{LawsuitNumberMP: "5002127-71.2019.8.13.0016", NameLawsuitMP: "GUILHERME MOREIRA SILVA", DocumentNumber: []string{"13274012661"}},
	{LawsuitNumberMP: "0023287-09.2020.8.13.0016", NameLawsuitMP: "GUILHERME MOREIRA SILVA", DocumentNumber: []string{"13274012661"}},
	{LawsuitNumberMP: "5001995-34.2022.8.13.0331", NameLawsuitMP: "GUSTAVO MONTEIRO DOS SANTOS", DocumentNumber: []string{"14172314678"}},
	{LawsuitNumberMP: "5001778-65.2018.8.13.0384", NameLawsuitMP: "ITAMAR GOMES DE OLIVEIRA", DocumentNumber: []string{"8122294600"}},
	{LawsuitNumberMP: "5003176-08.2022.8.13.0384", NameLawsuitMP: "ITAMAR GOMES DE OLIVEIRA", DocumentNumber: []string{"8122294600"}},
	{LawsuitNumberMP: "5005381-81.2022.8.13.0134", NameLawsuitMP: "ITAMAR GOMES DE OLIVEIRA", DocumentNumber: []string{"8122294600"}},
	{LawsuitNumberMP: "5000725-73.2023.8.13.0384", NameLawsuitMP: "ITAMAR GOMES DE OLIVEIRA", DocumentNumber: []string{"8122294600"}},
	{LawsuitNumberMP: "0049271-59.2020.8.13.0027", NameLawsuitMP: "JACKSON JUNIO FERREIRA RIBEIRO", DocumentNumber: []string{"15027974612"}},
	{LawsuitNumberMP: "0008528-49.2021.8.13.0713", NameLawsuitMP: "JOHNY DA SILVA BARBOSA", DocumentNumber: []string{"12635643644"}},
	{LawsuitNumberMP: "0012330-72.2021.8.13.0384", NameLawsuitMP: "JOSE EDUARDO RODRIGUES", DocumentNumber: []string{"7452547640"}},
	{LawsuitNumberMP: "5000484-70.2021.8.13.0384", NameLawsuitMP: "JOSE EDUARDO RODRIGUES", DocumentNumber: []string{"7452547640"}},
	{LawsuitNumberMP: "5001975-42.2023.8.13.0223", NameLawsuitMP: "JULIANO MARCAL BUENO", DocumentNumber: []string{"7481329613"}},
	{LawsuitNumberMP: "5015916-93.2022.8.13.0223", NameLawsuitMP: "JULIANO MARCAL BUENO", DocumentNumber: []string{"7481329613"}},
	{LawsuitNumberMP: "5000249-77.2016.8.13.0223", NameLawsuitMP: "JULIANO MARCAL BUENO", DocumentNumber: []string{"7481329613"}},
	{LawsuitNumberMP: "0019404-32.2019.8.13.0261", NameLawsuitMP: "JULIO CESAR SOARES SILVA", DocumentNumber: []string{"13083622694"}},
	{LawsuitNumberMP: "0214202-84.2017.8.13.0027", NameLawsuitMP: "LAZARO GARCIA FILHO", DocumentNumber: []string{"10907564615"}},
	{LawsuitNumberMP: "0064042-52.2014.8.13.0027", NameLawsuitMP: "LAZARO GARCIA FILHO", DocumentNumber: []string{"10907564615"}},
	{LawsuitNumberMP: "0056977-26.2002.8.13.0027", NameLawsuitMP: "LAZARO GARCIA FILHO", DocumentNumber: []string{"10907564615"}},
	{LawsuitNumberMP: "0053465-06.2000.8.13.0027", NameLawsuitMP: "LAZARO GARCIA FILHO", DocumentNumber: []string{"10907564615"}},
	{LawsuitNumberMP: "0011203-85.1993.8.13.0027", NameLawsuitMP: "LAZARO GARCIA FILHO", DocumentNumber: []string{"10907564615"}},
	{LawsuitNumberMP: "5020716-78.2019.8.13.0027", NameLawsuitMP: "LAZARO GARCIA FILHO", DocumentNumber: []string{"10907564615"}},
	{LawsuitNumberMP: "0021666-47.1997.8.13.0027", NameLawsuitMP: "LAZARO GARCIA FILHO", DocumentNumber: []string{"10907564615"}},
	{LawsuitNumberMP: "5020483-18.2018.8.13.0027", NameLawsuitMP: "LAZARO GARCIA FILHO", DocumentNumber: []string{"10907564615"}},
	{LawsuitNumberMP: "0003468-72.2021.8.13.0074", NameLawsuitMP: "LEANDRO PASCOAL MENDES", DocumentNumber: []string{"3675972682"}},
	{LawsuitNumberMP: "5002332-84.2020.8.13.0301", NameLawsuitMP: "LEANDRO PASCOAL MENDES", DocumentNumber: []string{"3675972682"}},
	{LawsuitNumberMP: "5000103-63.2016.8.13.0699", NameLawsuitMP: "LUCAS DA SILVA SANTOS", DocumentNumber: []string{"8632759663"}},
	{LawsuitNumberMP: "0049879-20.2016.8.13.0699", NameLawsuitMP: "LUCAS DA SILVA SANTOS", DocumentNumber: []string{"8632759663"}},
	{LawsuitNumberMP: "5012087-80.2022.8.13.0134", NameLawsuitMP: "LUCIANO RAMOS DE ARAUJO", DocumentNumber: []string{"252089693"}},
	{LawsuitNumberMP: "5001438-36.2023.8.13.0194", NameLawsuitMP: "LUCIANO RAMOS DE ARAUJO", DocumentNumber: []string{"252089693"}},
	{LawsuitNumberMP: "5000682-27.2023.8.13.0194", NameLawsuitMP: "LUCIANO RAMOS DE ARAUJO", DocumentNumber: []string{"252089693"}},
	{LawsuitNumberMP: "0032720-27.2014.8.13.0346", NameLawsuitMP: "LUIZ ALBERTO ARAUJO MARQUES", DocumentNumber: []string{"6584886603"}},
	{LawsuitNumberMP: "5027103-84.2018.8.13.0079", NameLawsuitMP: "LUIZ ALBERTO ARAUJO MARQUES", DocumentNumber: []string{"6584886603"}},
	{LawsuitNumberMP: "5017927-42.2022.8.13.0079", NameLawsuitMP: "LUIZ ALBERTO ARAUJO MARQUES", DocumentNumber: []string{"6584886603"}},
	{LawsuitNumberMP: "5008013-17.2023.8.13.0079", NameLawsuitMP: "LUIZ ALBERTO ARAUJO MARQUES", DocumentNumber: []string{"6584886603"}},
	{LawsuitNumberMP: "0015471-64.2021.8.13.0040", NameLawsuitMP: "MARCIO LUIZ MACHADO", DocumentNumber: []string{"58597344687"}},
	{LawsuitNumberMP: "5001344-83.2020.8.13.0456", NameLawsuitMP: "MARLON CASTRO MARTINS", DocumentNumber: []string{"7602190602"}},
	{LawsuitNumberMP: "5000942-02.2020.8.13.0456", NameLawsuitMP: "MARLON CASTRO MARTINS", DocumentNumber: []string{"7602190602"}},
	{LawsuitNumberMP: "5005255-35.2022.8.13.0456", NameLawsuitMP: "MARLON CASTRO MARTINS", DocumentNumber: []string{"7602190602"}},
	{LawsuitNumberMP: "5004782-78.2022.8.13.0026", NameLawsuitMP: "MATEUS ADRIANO RODRIGUES", DocumentNumber: []string{"15360226676"}},
	{LawsuitNumberMP: "5000108-24.2023.8.13.0349", NameLawsuitMP: "MATEUS ADRIANO RODRIGUES", DocumentNumber: []string{"15360226676"}},
	{LawsuitNumberMP: "0001885-56.2020.8.13.0084", NameLawsuitMP: "MATEUS ADRIANO RODRIGUES", DocumentNumber: []string{"15360226676"}},
	{LawsuitNumberMP: "5001950-71.2023.8.13.0209", NameLawsuitMP: "MATHEUS FONSECA DE OLIVEIRA", DocumentNumber: []string{"12646752679"}},
	{LawsuitNumberMP: "5018800-76.2023.8.13.0024", NameLawsuitMP: "MATHEUS FONSECA DE OLIVEIRA", DocumentNumber: []string{"12646752679"}},
	{LawsuitNumberMP: "0011723-44.2022.8.13.0701", NameLawsuitMP: "MATHEUS RIBEIRO SILVA", DocumentNumber: []string{"13024057641"}},
	{LawsuitNumberMP: "5002576-36.2023.8.13.0518", NameLawsuitMP: "MICHEL RIBEIRO", DocumentNumber: []string{"10332963683"}},
	{LawsuitNumberMP: "5008933-37.2020.8.13.0518", NameLawsuitMP: "MICHEL RIBEIRO", DocumentNumber: []string{"10332963683"}},
	{LawsuitNumberMP: "5006904-94.2022.8.13.0016", NameLawsuitMP: "MICHEL RIBEIRO", DocumentNumber: []string{"10332963683"}},
	{LawsuitNumberMP: "5003337-29.2022.8.13.0251", NameLawsuitMP: "PATRICK FERNANDES DA SILVA", DocumentNumber: []string{"10633488623"}},
	{LawsuitNumberMP: "0016075-73.2018.8.13.0155", NameLawsuitMP: "PATRICK FERNANDES DA SILVA", DocumentNumber: []string{"10633488623"}},
	{LawsuitNumberMP: "5000924-09.2023.8.13.0251", NameLawsuitMP: "PATRICK FERNANDES DA SILVA", DocumentNumber: []string{"10633488623"}},
	{LawsuitNumberMP: "5002745-79.2022.8.13.0058", NameLawsuitMP: "PAULO CESAR PEREIRA BRANDAO", DocumentNumber: []string{"8004937616"}},
	{LawsuitNumberMP: "0037425-86.2017.8.13.0209", NameLawsuitMP: "PAULO CESAR PEREIRA BRANDAO", DocumentNumber: []string{"8004937616"}},
	{LawsuitNumberMP: "0004360-45.2019.8.13.0140", NameLawsuitMP: "PEDRO AUGUSTO FERREIRA", DocumentNumber: []string{"13038683698"}},
	{LawsuitNumberMP: "5002412-04.2022.8.13.0002", NameLawsuitMP: "PEDRO AUGUSTO FERREIRA", DocumentNumber: []string{"13038683698"}},
	{LawsuitNumberMP: "5000409-42.2023.8.13.0002", NameLawsuitMP: "PEDRO AUGUSTO FERREIRA", DocumentNumber: []string{"13038683698"}},
	{LawsuitNumberMP: "5000328-93.2023.8.13.0002", NameLawsuitMP: "PEDRO AUGUSTO FERREIRA", DocumentNumber: []string{"13038683698"}},
	{LawsuitNumberMP: "0023646-69.2018.8.13.0002", NameLawsuitMP: "PEDRO AUGUSTO FERREIRA", DocumentNumber: []string{"13038683698"}},
	{LawsuitNumberMP: "5052929-73.2022.8.13.0079", NameLawsuitMP: "ROGERIO FERREIRA", DocumentNumber: []string{"91053757620"}},
	{LawsuitNumberMP: "0278209-02.2012.8.13.0079", NameLawsuitMP: "ROGERIO FERREIRA", DocumentNumber: []string{"91053757620"}},
	{LawsuitNumberMP: "0735709-87.2014.8.13.0079", NameLawsuitMP: "ROGERIO FERREIRA", DocumentNumber: []string{"91053757620"}},
	{LawsuitNumberMP: "6016047-42.2015.8.13.0079", NameLawsuitMP: "ROGERIO FERREIRA", DocumentNumber: []string{"91053757620"}},
	{LawsuitNumberMP: "5000922-07.2022.8.13.0079", NameLawsuitMP: "ROGERIO FERREIRA", DocumentNumber: []string{"91053757620"}},
	{LawsuitNumberMP: "0125629-46.2015.8.13.0672", NameLawsuitMP: "ROGERIO FERREIRA", DocumentNumber: []string{"91053757620"}},
	{LawsuitNumberMP: "0005249-24.2012.8.13.0696", NameLawsuitMP: "SEBASTIAO PAULO DA SILVA JUNIOR", DocumentNumber: []string{"10144418622"}},
	{LawsuitNumberMP: "0021557-28.2018.8.13.0696", NameLawsuitMP: "THIAGO ALVES DA SILVA", DocumentNumber: []string{"10834747669"}},
	{LawsuitNumberMP: "0059655-75.2020.8.13.0223", NameLawsuitMP: "TULIO CESAR DE OLIVEIRA", DocumentNumber: []string{"6960573602"}},
	{LawsuitNumberMP: "5023731-38.2017.8.13.0702", NameLawsuitMP: "WEBERTON RESENDE SANTOS DE BRITO", DocumentNumber: []string{"9400063679"}},
	{LawsuitNumberMP: "5023663-88.2017.8.13.0702", NameLawsuitMP: "WEBERTON RESENDE SANTOS DE BRITO", DocumentNumber: []string{"9400063679"}},
	{LawsuitNumberMP: "0764170-38.2017.8.13.0702", NameLawsuitMP: "WEBERTON RESENDE SANTOS DE BRITO", DocumentNumber: []string{"9400063679"}},
	{LawsuitNumberMP: "5014586-55.2017.8.13.0702", NameLawsuitMP: "WEBERTON RESENDE SANTOS DE BRITO", DocumentNumber: []string{"9400063679"}},
	{LawsuitNumberMP: "5008789-76.2022.8.13.0394", NameLawsuitMP: "WEBERTON RESENDE SANTOS DE BRITO", DocumentNumber: []string{"9400063679"}},
	{LawsuitNumberMP: "5012740-90.2023.8.13.0702", NameLawsuitMP: "WELLINGTON ALMEIDA RODRIGUES", DocumentNumber: []string{"11424462630"}},
	{LawsuitNumberMP: "0050964-35.2017.8.13.0431", NameLawsuitMP: "WELLINGTON ALMEIDA RODRIGUES", DocumentNumber: []string{"11424462630"}},
	{LawsuitNumberMP: "0026808-26.2020.8.13.0027", NameLawsuitMP: "WELLINGTON FERNANDES DAS NEVES", DocumentNumber: []string{"12069750671"}},
	{LawsuitNumberMP: "0000534-78.2018.8.13.0517", NameLawsuitMP: "WILIAN SABINO DOS SANTOS", DocumentNumber: []string{"11641847689"}},
	{LawsuitNumberMP: "5000587-37.2019.8.13.0517", NameLawsuitMP: "WILIAN SABINO DOS SANTOS", DocumentNumber: []string{"11641847689"}},
	{LawsuitNumberMP: "0007974-91.2019.8.13.0517", NameLawsuitMP: "WILIAN SABINO DOS SANTOS", DocumentNumber: []string{"11641847689"}},
	{LawsuitNumberMP: "0116196-75.2020.8.13.0079", NameLawsuitMP: "WILLIAM APARECIDO MARTINS", DocumentNumber: []string{"89383729600"}},
}
