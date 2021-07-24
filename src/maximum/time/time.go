package time

func maximumTime(time string) string {
	hour := ""
	minute := ""
	if time[0] == '?' {
		if time[1] == '?' {
			hour = "23"
		} else {
			if time[1] > '3' {
				hour = "1" + string(time[1])
			} else {
				hour = "2" + string(time[1])
			}
		}
	} else {
		if time[1] == '?' {
			if time[0] == '2' {
				hour = "23"
			} else {
				hour = string(time[0]) + "9"
			}
		} else {
			hour = time[:2]
		}
	}

	if time[3] == '?' {
		if time[4] == '?' {
			minute = "59"
		} else {
			minute = "5" + string(time[4])
		}
	} else {
		if time[4] == '?' {
			minute = string(time[3]) + "9"
		} else {
			minute = time[3:5]
		}
	}
	return hour + ":" + minute
}
