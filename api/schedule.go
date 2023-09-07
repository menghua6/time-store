package api

import (
	"strconv"
	"strings"
	"time"
)

const showDay = 20

func Schedule() (string, error) {
	timeCard := [24][showDay]int{}
	err := timeCardInit(&timeCard)
	if err != nil {
		return "", err
	}
	timeLine := ""
	for i := 0; i < 50; i++ {
		if i == 0 {
			timeLine += "  |"
			t := time.Now()
			for j := 0; j < showDay; j++ {
				date := t.Add(time.Hour * 24 * time.Duration(j)).String()[5:10]
				timeLine += " " + date + " |"
			}
			timeLine += " \n"
		} else if i%2 == 1 {
			timeLine += "  "
			for j := 0; j < showDay; j++ {
				timeLine += " -------"
			}
			timeLine += " \n"
		} else {
			hour := i/2 - 1
			timeLine += hourToString(hour) + "|"
			for j := 0; j < showDay; j++ {
				if timeCard[hour][j] == 0 {
					timeLine += "       |"
				} else if timeCard[hour][j] < 0 {
					timeLine += "   *   |"
				} else {
					num := strconv.Itoa(timeCard[hour][j])
					if len(num) > 7 {
						num = "inf"
					}
					timeLine += num
					for k := 0; k < 7-len(num); k++ {
						timeLine += " "
					}
					timeLine += "|"
				}
			}
			timeLine += " \n"
		}
	}
	return timeLine, nil
}

func timeCardInit(timeCard *[24][showDay]int) error {
	//work-day
	specialWorkDayMap := make(map[time.Time]int)
	specialWorkDayMap[time.Date(2023, 10, 7, 0, 0, 0, 0, time.Local)] = 1
	specialWorkDayMap[time.Date(2023, 10, 8, 0, 0, 0, 0, time.Local)] = 1

	specialRestDayMap := make(map[time.Time]int)
	specialRestDayMap[time.Date(2023, 9, 29, 0, 0, 0, 0, time.Local)] = 1
	specialRestDayMap[time.Date(2023, 10, 02, 0, 0, 0, 0, time.Local)] = 1
	specialRestDayMap[time.Date(2023, 10, 03, 0, 0, 0, 0, time.Local)] = 1
	specialRestDayMap[time.Date(2023, 10, 04, 0, 0, 0, 0, time.Local)] = 1
	specialRestDayMap[time.Date(2023, 10, 05, 0, 0, 0, 0, time.Local)] = 1
	specialRestDayMap[time.Date(2023, 10, 06, 0, 0, 0, 0, time.Local)] = 1

	partTimeJobMap := make(map[time.Time]string)
	partTimeJobMap[time.Date(2023, 9, 10, 0, 0, 0, 0, time.Local)] = "10-12,14-18"

	now := time.Now()
	now = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	for i := 0; i < showDay; i++ {
		date := now.Add(time.Hour * 24 * time.Duration(i))
		if date.Weekday() != time.Saturday && date.Weekday() != time.Sunday {
			if _, ok := specialRestDayMap[date]; !ok {
				for j := 10; j < 19; j++ {
					timeCard[j][i] = -1
				}
			}
		} else {
			if _, ok := specialWorkDayMap[date]; ok {
				for j := 10; j < 19; j++ {
					timeCard[j][i] = -1
				}
			}
		}
		if val, ok := partTimeJobMap[date]; ok {
			partTimeJobs := strings.Split(val, ",")
			for j := 0; j < len(partTimeJobs); j++ {
				se := strings.Split(partTimeJobs[j], "-")
				s, _ := strconv.Atoi(se[0])
				e, _ := strconv.Atoi(se[1])
				for k := s; k < e; k++ {
					timeCard[k][i] = -1
				}
			}
		}
	}
	return nil
}

func hourToString(hour int) string {
	if hour == 0 {
		return "00"
	} else if hour == 1 {
		return "01"
	} else if hour == 2 {
		return "02"
	} else if hour == 3 {
		return "03"
	} else if hour == 4 {
		return "04"
	} else if hour == 5 {
		return "05"
	} else if hour == 6 {
		return "06"
	} else if hour == 7 {
		return "07"
	} else if hour == 8 {
		return "08"
	} else if hour == 9 {
		return "09"
	} else if hour == 10 {
		return "10"
	} else if hour == 11 {
		return "11"
	} else if hour == 12 {
		return "12"
	} else if hour == 13 {
		return "13"
	} else if hour == 14 {
		return "14"
	} else if hour == 15 {
		return "15"
	} else if hour == 16 {
		return "16"
	} else if hour == 17 {
		return "17"
	} else if hour == 18 {
		return "18"
	} else if hour == 19 {
		return "19"
	} else if hour == 20 {
		return "20"
	} else if hour == 21 {
		return "21"
	} else if hour == 22 {
		return "22"
	} else if hour == 23 {
		return "23"
	}
	return ""
}
