package src

import (
	"os/exec"
	"strconv"
	"strings"

	"github.com/pdfcpu/pdfcpu/pkg/api"
)

func extractTextViaPdftotext(path string) (string, error) {
	cmd := exec.Command("pdftotext", path, "-")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

// FindDate sucht nach dem ersten erkennbaren Datum im Text und gibt Jahr, Monat, Tag als String zurück.
func FindDate(text string) (year, month, day string) {
	lower := strings.ToLower(text)

	for _, pat := range datePatterns {
		matches := pat.FindStringSubmatch(lower)
		if len(matches) < 4 {
			continue
		}

		g1, g2, g3 := matches[1], matches[2], matches[3]

		if m, ok := monthNames[g2]; ok && len(g2) > 2 {
			if len(g1) == 4 {
				// YYYY MONTH DD
				y, err1 := strconv.Atoi(g1)
				d, err2 := strconv.Atoi(g3)
				if err1 == nil && err2 == nil && y >= 2000 && y <= 2099 && d >= 1 && d <= 31 {
					return strconv.Itoa(y), strconv.Itoa(m), strconv.Itoa(d)
				}
			} else {
				// DD MONTH YYYY
				d, err1 := strconv.Atoi(g1)
				y, err2 := strconv.Atoi(g3)
				if err1 == nil && err2 == nil && y >= 2000 && y <= 2099 && d >= 1 && d <= 31 {
					return strconv.Itoa(y), strconv.Itoa(m), strconv.Itoa(d)
				}
			}
		}

		if len(g1) == 4 {
			y, _ := strconv.Atoi(g1)
			m, _ := strconv.Atoi(g2)
			d, _ := strconv.Atoi(g3)
			if y >= 2000 && y <= 2099 && m >= 1 && m <= 12 && d >= 1 && d <= 31 {
				return strconv.Itoa(y), strconv.Itoa(m), strconv.Itoa(d)
			}
		}

		if len(g3) == 4 {
			d, _ := strconv.Atoi(g1)
			m, _ := strconv.Atoi(g2)
			y, _ := strconv.Atoi(g3)
			if y >= 2000 && y <= 2099 && m >= 1 && m <= 12 && d >= 1 && d <= 31 {
				return strconv.Itoa(y), strconv.Itoa(m), strconv.Itoa(d)
			}
		}

		if m, ok := monthNames[g1]; ok {
			d, err1 := strconv.Atoi(g2)
			y, err2 := strconv.Atoi(g3)
			if err1 == nil && err2 == nil && y >= 2000 && y <= 2099 && d >= 1 && d <= 31 {
				return strconv.Itoa(y), strconv.Itoa(m), strconv.Itoa(d)
			}
		}
	}

	return "", "", ""
}

func findDate(text string) (year, month, day int) {
	y, m, d := FindDate(text)
	if y == "" {
		return 0, 0, 0
	}
	yi, _ := strconv.Atoi(y)
	mi, _ := strconv.Atoi(m)
	di, _ := strconv.Atoi(d)
	return yi, mi, di
}

func findCorrespondent(text string, correspondents []string) string {
	lower := strings.ToLower(text)
	for _, c := range correspondents {
		if strings.Contains(lower, strings.ToLower(c)) {
			return c
		}
	}
	return ""
}

func analyzePDF(path string, correspondents []string) (*Analysis, error) {
	text, err := extractTextViaPdftotext(path)
	if err != nil {
		text = ""
	}

	year, month, day := findDate(text)
	corr := findCorrespondent(text, correspondents)

	return &Analysis{
		DateYear:      year,
		DateMonth:     month,
		DateDay:       day,
		Correspondent: corr,
	}, nil
}

func writePDFTags(path string, tags []string) error {
	tagStr := strings.Join(tags, ", ")
	return api.AddPropertiesFile(path, path, map[string]string{"lowsky-pdf": tagStr}, nil)
}
