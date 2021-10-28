package etl

import "regexp"

var (
	regexList = map[string]*regexp.Regexp{
		"date":      regexp.MustCompile(`\d\d/\d\d/\d\d\d\d`),
		"time":      regexp.MustCompile(`\d\d:\d\d:\d\d`),
		"place":     regexp.MustCompile(`INAOE|UPTLAX`),
		"lightning": regexp.MustCompile(`Lightning Detected`),
		"distance":  regexp.MustCompile(`at\s\d\d\skm|at\s\d\skm`),
	}
)
