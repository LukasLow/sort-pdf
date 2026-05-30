package src

import (
	"os/exec"
	"strings"
)

func extractTextViaPdftotext(path string) (string, error) {
	cmd := exec.Command("pdftotext", path, "-")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
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

	corr := findCorrespondent(text, correspondents)

	return &Analysis{
		Correspondent: corr,
	}, nil
}


