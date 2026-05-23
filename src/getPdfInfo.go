package src

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pdfcpu/pdfcpu/pkg/api"
)

// PdfInfo holds basic metadata about a PDF file.
type PdfInfo struct {
	FileName  string `json:"fileName"`
	SizeBytes int64  `json:"sizeBytes"`
	PageCount int    `json:"pageCount"`
	DPI       int    `json:"dpi"`
	OCR       bool   `json:"ocr"`
}

// GetPdfInfo returns metadata for the given PDF located in the "900-Eingang" folder of workDir.
func getPdfInfo(workDir, fileName string) (*PdfInfo, error) {
	if workDir == "" {
		return nil, fmt.Errorf("workDir is empty")
	}
	if fileName == "" {
		return nil, fmt.Errorf("fileName is empty")
	}
	filePath := filepath.Join(workDir, "900-Eingang", fileName)

	// File size
	info, err := os.Stat(filePath)
	if err != nil {
		return nil, err
	}
	sizeBytes := info.Size()

	// Page count using pdfcpu (if it fails we still return size information)
	pageCount := 0
	ctx, err := api.ReadContextFile(filePath)
	if err == nil && ctx != nil {
		pageCount = ctx.PageCount
	}

	// DPI and OCR detection are not implemented yet – placeholders
	return &PdfInfo{
		FileName:  fileName,
		SizeBytes: sizeBytes,
		PageCount: pageCount,
		DPI:       0,
		OCR:       false,
	}, nil
}
