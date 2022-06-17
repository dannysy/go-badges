package internal

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

const coverageUrl = `https://img.shields.io/badge/Coverage-%s25-%s`

func UpsertCoverageBadge(coveragePath, readmePath string) error {
	bytes, err := ioutil.ReadFile(coveragePath)
	if err != nil {
		return fmt.Errorf("can't read coverage test output file %w", err)
	}
	fields := strings.Fields(string(bytes))
	percentage := fields[len(fields)-1]
	val, err := strconv.ParseFloat(percentage[:len(percentage)-1], 32)
	if err != nil {
		return fmt.Errorf("can't parse coverage percent %w", err)
	}
	badge := fmt.Sprintf(coverageUrl, percentage, getCoverageColor(val))
	badgeLink := fmt.Sprintf("![GoCoverage](%s)", badge)
	readmeBytes, err := ioutil.ReadFile(readmePath)
	if err != nil {
		return fmt.Errorf("can't open readme file %w", err)
	}
	r, err := regexp.Compile(`!\[GoCoverage]\(.*\)`)
	if err != nil {
		return fmt.Errorf("regex compiling error %w", err)
	}
	var readmeOut string
	if r.MatchString(string(readmeBytes)) {
		readmeOut = r.ReplaceAllString(string(readmeBytes), badgeLink)
	} else {
		readmeOut = badgeLink + "\n" + string(readmeBytes)
	}
	err = ioutil.WriteFile(readmePath, []byte(readmeOut), 0644)
	if err != nil {
		return fmt.Errorf("can't write to readme file %w", err)
	}
	return nil
}

func getCoverageColor(percent float64) string {
	switch {
	case percent >= 90:
		return "brightgreen"
	case percent >= 70:
		return "green"
	case percent >= 60:
		return "yellowgreen"
	case percent >= 50:
		return "yellow"
	case percent >= 40:
		return "orange"
	default:
		return "red"
	}
}
