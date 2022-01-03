package week

var week = []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

func dayOfTheWeek(day int, month int, year int) string {
	years := year - 1968
	days := 1461 * (years / 4)
	years %= 4
	for i := 1; i <= years; i += 1 {
		if (year-i)%4 == 0 {
			days += 366
		} else {
			days += 365
		}
	}
	for i := 1; i < month; i += 1 {
		if i < 8 {
			if i == 2 {
				if year%4 == 0 && year != 2100 {
					days += 29
				} else {
					days += 28
				}
			} else {
				if i&1 == 1 {
					days += 31
				} else {
					days += 30
				}
			}
		} else {
			if i&1 == 1 {
				days += 30
			} else {
				days += 31
			}
		}
	}
	days += day
	return week[days%7]
}
