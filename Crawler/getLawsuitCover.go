package Crawler

import (
	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
	"strconv"
	"strings"
)

type LawsuitCover struct {
	LawsuitNumber   string
	Class           string
	Subject         string
	LawsuitRefer    string
	Jurisdiction    string
	Aut             string
	LastMovement    string
	Value           string
	Secret          string
	FreeJustice     string
	LegalProtection string
	Priority        string
	District        string
	JudgePosition   string
	LegalCompetence string
}

const (
	titleClass           = "Classe judicial"
	titleSubject         = "Assunto"
	titleLawsuitRefer    = "Processo referência"
	titleJurisdiction    = "Jurisdição"
	titleAut             = "Autuação"
	titleLastMovement    = "Última distribuição"
	titleValue           = "Valor da causa"
	titleSecret          = "Segredo de justiça?"
	titleFreeJustice     = "Justiça gratuita?"
	titleLegalProtection = "Tutela/liminar?"
	titlePriority        = "Prioridade?"
	titleDistrict        = "Órgão julgador"
	titleJudgePosition   = "Cargo judicial"
	titleLegalCompetence = "Competência"
)

func GetLawsuitCover(htmlPgSrc *html.Node, lawsuitNumber string) (LawsuitCover, error) {
	var class string
	var subject string
	var lawsuitRefer string
	var jurisdiction string
	var aut string
	var lastMovement string
	var value string
	var secret string
	var freeJustice string
	var legalProtection string
	var priority string
	var district string
	var judgePosition string
	var legalCompetence string

	//*[@id="maisDetalhes"]/div[1]/dl/dd/text()
	totalNodes := htmlquery.Find(htmlPgSrc, "//*[@id=\"maisDetalhes\"]/dl/dt")
	for i, node := range totalNodes {
		title := htmlquery.InnerText(node)
		information := strings.ReplaceAll(strings.TrimSpace(htmlquery.InnerText(htmlquery.FindOne(htmlPgSrc, "//*[@id=\"maisDetalhes\"]/dl/dd["+strconv.Itoa(i+1)+"]/text()"))), "\n", "")

		if title == titleClass {
			class = information
		} else if title == titleSubject {
			nodes := htmlquery.Find(htmlPgSrc, "//*[@id=\"maisDetalhes\"]/dl/dd["+strconv.Itoa(i+1)+"]/ul/li")
			if len(nodes) != 0 {
				information2 := strings.ReplaceAll(strings.TrimSpace(htmlquery.InnerText(nodes[0])), "\n", "")
				subject = information2
			}
		} else if title == titleLawsuitRefer {
			lawsuitRefer = information
		} else if title == titleJurisdiction {
			jurisdiction = information
		} else if title == titleAut {
			aut = information
		} else if title == titleLastMovement {
			lastMovement = information
		} else if title == titleValue {
			value = information
		} else if title == titleSecret {
			secret = information
		} else if title == titleFreeJustice {
			freeJustice = information
		} else if title == titleLegalProtection {
			legalProtection = information
		} else if title == titlePriority {
			priority = information
		}
	}

	totalNodes2 := htmlquery.Find(htmlPgSrc, "//*[@id=\"maisDetalhes\"]/div")
	for _, node2 := range totalNodes2 {
		title := htmlquery.InnerText(node2)
		information := strings.ReplaceAll(strings.TrimSpace(htmlquery.InnerText(htmlquery.FindOne(node2, "/dl/dd/text()"))), "\n", "")

		if title == titleDistrict {
			district = information
		} else if title == titleJudgePosition {
			judgePosition = information
		} else if title == titleLegalCompetence {
			legalCompetence = information
		}
	}

	return LawsuitCover{
		LawsuitNumber:   lawsuitNumber,
		Class:           class,
		Subject:         subject,
		LawsuitRefer:    lawsuitRefer,
		Jurisdiction:    jurisdiction,
		Aut:             aut,
		LastMovement:    lastMovement,
		Value:           value,
		Secret:          secret,
		FreeJustice:     freeJustice,
		LegalProtection: legalProtection,
		Priority:        priority,
		District:        district,
		JudgePosition:   judgePosition,
		LegalCompetence: legalCompetence,
	}, nil

}
