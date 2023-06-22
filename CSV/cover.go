package CSV

import (
	"Crawler_TJMG_PGE/Crawler"
	"encoding/csv"
)

const fileNameC = "Covers"

func WriteCovers(lawsuits []Crawler.Lawsuit) error {
	var rows [][]string

	rows = append(rows, generateCoverHeaders())

	for _, lawsuit := range lawsuits {
		rows = append(rows, tableCoverRows(lawsuit))
	}

	cf, err := createFile(folderName + "/" + fileNameC + ".csv")
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

func generateCoverHeaders() []string {
	return []string{
		"Número do Processo",
		"Classe judicial",
		"Assunto",
		"Processo referência",
		"Jurisdição",
		"Autuação",
		"Última distribuição",
		"Valor da causa",
		"Segredo de justiça?",
		"Justiça gratuita?",
		"Tutela/liminar?",
		"Prioridade?",
		"Órgão julgador",
		"Cargo judicial",
		"Competência",
	}
}

func tableCoverRows(results Crawler.Lawsuit) []string {
	return []string{
		results.Cover.LawsuitNumber,
		results.Cover.Class,
		results.Cover.Subject,
		results.Cover.LawsuitRefer,
		results.Cover.Jurisdiction,
		results.Cover.Aut,
		results.Cover.LastMovement,
		results.Cover.Value,
		results.Cover.Secret,
		results.Cover.FreeJustice,
		results.Cover.LegalProtection,
		results.Cover.Priority,
		results.Cover.District,
		results.Cover.JudgePosition,
		results.Cover.LegalCompetence,
	}
}
