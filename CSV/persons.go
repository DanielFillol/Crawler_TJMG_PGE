package CSV

import (
	"Crawler_TJMG_PGE/Crawler"
	"encoding/csv"
)

const fileNameP = "Persons"

func WritePersons(lawsuits []Crawler.Lawsuit) error {
	var rows [][]string

	rows = tablePersonsRows(lawsuits)

	cf, err := createFile(folderName + "/" + fileNameP + ".csv")
	if err != nil {
		return err
	}

	w := csv.NewWriter(cf)

	err = w.WriteAll(rows)
	if err != nil {
		return err
	}

	return nil
}

func tablePersonsRows(lawsuits []Crawler.Lawsuit) [][]string {
	var prs [][]string
	prs = append(prs, []string{"Processo", "Polo", "Nome"})

	for _, lawsuit := range lawsuits {
		for _, person1 := range lawsuit.Poles.ActivePole {
			prs = append(prs, []string{lawsuit.Poles.LawsuitNumber, "ativo", person1})
		}
		for _, person2 := range lawsuit.Poles.PassivePole {
			prs = append(prs, []string{lawsuit.Poles.LawsuitNumber, "passivo", person2})
		}
		for _, person3 := range lawsuit.Poles.OtherPole {
			prs = append(prs, []string{lawsuit.Poles.LawsuitNumber, "outros", person3})
		}
	}

	return prs
}
