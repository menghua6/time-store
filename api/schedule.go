package api

import (
	"strconv"
	"time"
)

func Schedule() (string, error) {
	timeCard := [24][14]int{}
	timeCardInit(timeCard)
	timeLine := ""
	for i := 0; i < 50; i++ {
		if i == 0 {
			timeLine += "  |"
			t := time.Now()
			for j := 0; j < 14; j++ {
				date := t.Add(time.Hour * 24 * time.Duration(j)).String()[5:10]
				timeLine += " " + date + " |"
			}
			timeLine += " \n"
		} else if i%2 == 1 {
			timeLine += "   ------- ------- ------- ------- ------- ------- ------- ------- ------- ------- ------- ------- ------- ------- \n"
		} else {
			hour := i/2 - 1
			timeLine += hourToString(hour) + "|"
			for j:=0;j<14;j++{
				if timeCard[hour][j] == 0{
					timeLine += "       |"
				}else if timeCard[hour][j] < 0{
					timeLine += "   *   |"
				}else {
					num := strconv.Itoa(timeCard[hour][j])
					if len(num) > 7{
						num = "inf"
					}
					timeLine += num
					for k:=0;k<7-len(num);k++{
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

func timeCardInit(timeCard [24][14]int) {
	
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
