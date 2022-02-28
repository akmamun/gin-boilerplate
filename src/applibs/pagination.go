package applibs

import "strconv"

type Data struct {
	TotalData    int64
	FilteredData int64
	Data         interface{}
}

type Args struct {
	Sort   string
	Order  string
	Offset string
	Limit  string
	Search string
}

func Offset(offset string) int {
	offsetInt, err := strconv.Atoi(offset)
	if err != nil {
		offsetInt = 0
	}
	return offsetInt
}

// Limit returns the number of result for pagination
func Limit(limit string) int {
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		limitInt = 25
	}
	return limitInt
}
