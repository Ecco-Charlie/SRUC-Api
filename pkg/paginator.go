package pkg

import (
	"net/url"
	"strconv"
)

func GeneratePaginator(total int, actual *int) *[]string {
	can := (total / 11) + 1
	if *actual > can {
		*actual = 1
	}
	const middle = 5
	var p []string

	if can < middle+4 {
		for i := range can {
			p = append(p, strconv.Itoa(i+1))
		}
		return &p
	}

	p = append(p, "1")
	p = append(p, "2")

	if *actual > 5 {
		p = append(p, "...")
	}

	s := *actual - 2
	e := *actual + 2

	if s < 3 {
		s = 3
		e = s + 4
	}

	if e > can-2 {
		e = can - 2
		s = e - 4
	}

	if s < 3 {
		s = 3
	}

	for i := s; i <= e; i++ {
		p = append(p, strconv.Itoa(i))
	}

	if e < can-2 {
		p = append(p, "...")
	}

	p = append(p, strconv.Itoa(can-1))
	p = append(p, strconv.Itoa(can))

	return &p
}

func GetCurrentPage(params *url.Values) int {
	v, err := strconv.Atoi(params.Get("p"))
	if err != nil {
		return 1
	}

	if params.Has("pp") {
		v -= 1
	}

	if params.Has("pn") {
		v += 1
	}

	return v
}
