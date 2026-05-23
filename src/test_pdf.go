package src

import (
	"fmt"
	"os"
	"path/filepath"
)

func createTestPDF(path string) error {
	text := `BT
/F1 12 Tf
50 700 Td
(15. Mai 2024 - Rechnung) Tj
0 -20 Td
(Datum: 2024-03-12) Tj
0 -20 Td
(Klemper - Beispiel GmbH) Tj
0 -20 Td
(28. Dezember 2023) Tj
0 -20 Td
(2022 Januar 05 - Lieferung) Tj
0 -20 Td
(March 1, 2024) Tj
0 -20 Td
(15.03.2024 - Zahlungseingang) Tj
ET`

	contentStr := text
	obj4 := fmt.Sprintf("4 0 obj\n<< /Length %d >>\nstream\n%s\nendstream\nendobj", len(contentStr), contentStr)

	objects := []string{
		"%PDF-1.4",
		"1 0 obj\n<< /Type /Catalog /Pages 2 0 R >>\nendobj",
		"2 0 obj\n<< /Type /Pages /Kids [3 0 R] /Count 1 >>\nendobj",
		"3 0 obj\n<< /Type /Page /Parent 2 0 R /MediaBox [0 0 612 792] /Contents 4 0 R /Resources << /Font << /F1 5 0 R >> >> >>\nendobj",
		obj4,
		"5 0 obj\n<< /Type /Font /Subtype /Type1 /BaseFont /Helvetica >>\nendobj",
	}

	var buf []byte
	var xrefOffsets []int64

	for _, obj := range objects {
		xrefOffsets = append(xrefOffsets, int64(len(buf)))
		buf = append(buf, []byte(obj+"\n")...)
	}

	xrefOffset := len(buf)
	xref := "xref\n"
	xref += fmt.Sprintf("0 %d\n", len(objects)+1)
	xref += "0000000000 65535 f \n"
	for _, o := range xrefOffsets {
		xref += fmt.Sprintf("%010d 00000 n \n", o)
	}

	buf = append(buf, []byte(xref)...)
	buf = append(buf, []byte(fmt.Sprintf("trailer\n<< /Size %d /Root 1 0 R >>\n", len(objects)+1))...)
	buf = append(buf, []byte("startxref\n")...)
	buf = append(buf, []byte(fmt.Sprintf("%d\n", xrefOffset))...)
	buf = append(buf, []byte("%%EOF")...)

	return os.WriteFile(path, buf, 0644)
}

// CreateTestPDF erstellt eine Test-PDF mit verschiedenen Datumsformaten
func (b *WorkspaceBridge) CreateTestPDF() (string, error) {
	if b.currentWorkDir == "" {
		return "", fmt.Errorf("kein Arbeitsverzeichnis geladen")
	}
	path := filepath.Join(b.currentWorkDir, "900-Eingang", "test_datumserkennung.pdf")
	err := createTestPDF(path)
	if err != nil {
		return "", err
	}
	return "test_datumserkennung.pdf", nil
}
