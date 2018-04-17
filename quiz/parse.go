package main

import "strings"

func parseRecords(records [][]string) []problem {
	ret := make([]problem, len(records))
	for i, r := range records {
		ret[i] = problem{
			q: r[0],
			a: strings.ToLower(strings.TrimSpace(r[1])),
		}
	}
	return ret
}
