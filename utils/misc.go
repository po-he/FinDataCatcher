package utils

import(
	"strings"
	"strconv"
	"time"
	"regexp"
	"fmt"
)

func StoInt(v string) int {
	v = strings.Trim(v, " ")
	n, err := strconv.Atoi(v)
	if err != nil {
		return 0
	}
	return n
}


func StoI32(v string) int32 {
	v = strings.Trim(v, " ")
	n, err := strconv.Atoi(v)
	if err != nil {
		return 0
	}
	return int32(n)
}

func StoI64(v string) int64 {
	v = strings.Trim(v, " ")
	n, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return 0
	}
	return n
}

func Stof64(v string) float64 {
	v = strings.Trim(v, " ")
	if val, err := strconv.ParseFloat(v, 64); err == nil {
		return val
	}

	return 0
}

func DateString2LocalNanoUnixTime(s string, stringFormatParser func (string) []string) int64 {
	slices := stringFormatParser(s)
	if slices == nil || len(slices) != 7 {
		return -1
	}

	year := StoInt(slices[1])
	month := StoInt(slices[2])
	day := StoInt(slices[3])
	hour := StoInt(slices[4])
	min := StoInt(slices[5])
	sec := StoInt(slices[6])
	loc, _ := time.LoadLocation("Local") 
	t := time.Date(year, time.Month(month), day, hour, min, sec, 0, loc)
	return int64(t.UnixNano())
}

/*
输入格式："YY-MM-DD H:M:S"
返回：毫秒級本地時間戳
*/
func NormalDate2LocalUnixTimeMs(s string) int64 {
	parser := func(str string) []string {
				re := regexp.MustCompile(`([\d]+)-([\d]+)-([\d]+) ([\d]+):([\d]+):([\d]+)`)
				return re.FindStringSubmatch(s)
			  }

	return DateString2LocalNanoUnixTime(s, parser) / 1000000
}

/*
输入格式："YY-MM-DD H:M:S"
返回：秒級本地時間戳
*/
func NormalDate2LocalUnixTimeSec(s string) int64 {
	parser := func(str string) []string {
				re := regexp.MustCompile(`([\d]+)-([\d]+)-([\d]+) ([\d]+):([\d]+):([\d]+)`)
				return re.FindStringSubmatch(s)
			  }

	return DateString2LocalNanoUnixTime(s, parser) / 1000000000
}

/*
*/
func CurrentLocalDate() string {
	t := time.Now().Local()
	return strings.Split(t.Format("2006-01-02 15:04:05"), " ")[0]
}

func CurrentLocalTime() string {
	t := time.Now().Local()
	return t.Format("2006-01-02 15:04:05")
}

/*
时间戳转年月日
*/
func Ms2LocalTime(ms int64) time.Time {
	sec := ms / 1e3
	nsec := (ms % 1e3) * 1e6
	return time.Unix(sec, nsec).Local()
}

func Ms2LocalDateString(ms int64) string {
	tm := Ms2LocalTime(ms)
	year, mon, day := tm.Local().Date()
	hour, min, sec := tm.Local().Clock()
	return fmt.Sprintf("%4d-%02d-%02d %02d:%02d:%02d", year, mon, day, hour, min, sec)
}

func Sec2LocalTime(sec int64) time.Time {
	return Ms2LocalTime(sec * 1000)
}

func Sec2LocalDataString(sec int64) string {
	return Ms2LocalDateString(sec * 1000)
}

