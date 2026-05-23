package src

import "regexp"

var monthNames = map[string]int{
	"januar": 1, "jan": 1, "jänner": 1, "jäner": 1,
	"februar": 2, "feb": 2, "feber": 2,
	"märz":   3, "maerz": 3,
	"april":  4,
	"mai":    5,
	"juni":   6, "jun": 6,
	"juli":   7, "jul": 7,
	"august": 8,
	"september": 9,
	"oktober": 10, "okt": 10,
	"november": 11,
	"dezember": 12, "dez": 12,

	"january":  1,
	"february": 2,
	"march":    3, "mar": 3,
	"may":   5,
	"june":  6,
	"july":  7,
	"aug":   8,
	"sep":   9,
	"october": 10, "oct": 10,
	"nov":   11,
	"december": 12, "dec": 12,
}

var datePatterns = []*regexp.Regexp{
	regexp.MustCompile(`(\d{4})\s*-\s*(\d{1,2})\s*-\s*(\d{1,2})`),
	regexp.MustCompile(`(\d{1,2})\s*\.\s*(\d{1,2})\s*\.\s*(\d{4})`),
	regexp.MustCompile(`(\d{1,2})\.?\s+([A-Za-zäöüßÄÖÜéêè]+)\s+(\d{4})`),
	regexp.MustCompile(`(\d{4})\s+([A-Za-zäöüßÄÖÜéêè]+)\s+(\d{1,2})`),
	regexp.MustCompile(`([A-Za-zäöüßÄÖÜéêè]+)\s+(\d{1,2})\,?\s*(\d{4})`),
}
