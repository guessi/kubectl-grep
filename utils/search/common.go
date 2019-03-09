package search

func getAge(duration float64) (int, string) {
	var age int
	var unit string

	if duration >= 86400 {
		age = int(duration) / 86400
		unit = "d"
	} else if duration > 3600 {
		age = int(duration) / 3600
		unit = "h"
	} else if duration > 60 {
		age = int(duration) / 60
		unit = "m"
	} else {
		age = int(duration)
		unit = "s"
	}

	return age, unit
}
