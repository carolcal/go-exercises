package expenses

import "fmt"

type Record struct {
	Day      int
	Amount   float64
	Category string
}

type DaysPeriod struct {
	From int
	To   int
}

func Filter(in []Record, predicate func(Record) bool) []Record {
	var filtered []Record
    for _, v := range in {
        if predicate(v) {
            filtered = append(filtered, v)
        }
    }
    return filtered
}

func ByDaysPeriod(p DaysPeriod) func(Record) bool {
    return func(r Record) bool {
    	return r.Day >= p.From && r.Day <= p.To
    }
}

func ByCategory(c string) func(Record) bool {
    return func(r Record) bool {
    	return r.Category == c 
    }
}

func TotalByPeriod(in []Record, p DaysPeriod) float64 {
    var sum float64 = 0
	filtered := Filter(in, ByDaysPeriod(p))
    for _, f := range filtered {
        sum = sum + f.Amount
    }
    return sum
}

type InvalidCategoryError struct {
    category string
}

func (e *InvalidCategoryError) Error() string {
    return fmt.Sprintf("unknown category %s", e.category)
}

func CategoryExpenses(in []Record, p DaysPeriod, c string) (float64, error) {
	var sum float64 = 0
    filteredByPeriod := Filter(in, ByDaysPeriod(p))
    if len(filteredByPeriod) == 0 {
        return 0, nil
    }
    filteredByCategory := Filter(filteredByPeriod, ByCategory(c))
    if len(filteredByCategory) == 0 {
        return 0, &InvalidCategoryError{category: c}
    }
    for _, f := range filteredByCategory {
        sum = sum + f.Amount
    }
    return sum, nil
}
