package main

func parseRecords(records [][]string) []problem {
	ret := make([]problem, len(records))
	for i, r := range records {
		ret[i] = problem{
			q: r[0],
			a: r[1],
		}
	}
	return ret
}
