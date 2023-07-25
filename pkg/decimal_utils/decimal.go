package utils

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/serenefiregroup/ffa_server/pkg/log"
)

// RoundFloat64 保留 precision 位小数，并进行四舍五入
func RoundFloat64(f float64, precision int) float64 {
	if precision < 0 || precision > 5 {
		precision = 5
	}
	var result float64
	f1 := f * math.Pow10(precision)
	ff := math.Floor(f1)
	if ff+0.5 > f1 {
		result = ff / math.Pow10(precision)
	} else {
		result = (ff + 1) / math.Pow10(precision)
	}
	return result
}

// NumFormat 将数字金额转成字符串，带2位小数，按千位分
func NumFormat(num float64) string {
	str := strconv.FormatFloat(num, 'f', 2, 64)
	parts := strings.Split(str, ".")
	numStr := parts[0]
	length := len(numStr)
	if length < 4 {
		return str
	}
	count := (length - 1) / 3 //取于-有多少组三位数
	for i := 0; i < count; i++ {
		numStr = numStr[:length-(i+1)*3] + "," + numStr[length-(i+1)*3:]
	}
	numStr = fmt.Sprintf("%s.%s", numStr, parts[1])
	return numStr
}

// ConvertNumToCny 将数值转为中文大写
func ConvertNumToCny(num float64) string {
	strnum := strconv.FormatFloat(num*100, 'f', 0, 64)
	sliceUnit := []string{"仟", "佰", "拾", "亿", "仟", "佰", "拾", "万", "仟", "佰", "拾", "元", "角", "分"}
	// log.Println(sliceUnit[:len(sliceUnit)-2])
	s := sliceUnit[len(sliceUnit)-len(strnum) : len(sliceUnit)]
	upperDigitUnit := map[string]string{"0": "零", "1": "壹", "2": "贰", "3": "叁", "4": "肆", "5": "伍", "6": "陆", "7": "柒", "8": "捌", "9": "玖"}
	str := ""
	for k, v := range strnum[:] {
		str = str + upperDigitUnit[string(v)] + s[k]
	}
	reg, err := regexp.Compile(`零角零分$`)
	str = reg.ReplaceAllString(str, "整")

	reg, err = regexp.Compile(`零角`)
	str = reg.ReplaceAllString(str, "零")

	reg, err = regexp.Compile(`零分$`)
	str = reg.ReplaceAllString(str, "整")

	reg, err = regexp.Compile(`零[仟佰拾]`)
	str = reg.ReplaceAllString(str, "零")

	reg, err = regexp.Compile(`零{2,}`)
	str = reg.ReplaceAllString(str, "零")

	reg, err = regexp.Compile(`零亿`)
	str = reg.ReplaceAllString(str, "亿")

	reg, err = regexp.Compile(`零万`)
	str = reg.ReplaceAllString(str, "万")

	reg, err = regexp.Compile(`零*元`)
	str = reg.ReplaceAllString(str, "元")

	reg, err = regexp.Compile(`亿零{0, 3}万`)
	str = reg.ReplaceAllString(str, "^元")

	reg, err = regexp.Compile(`零元`)
	str = reg.ReplaceAllString(str, "零")
	if err != nil {
		log.Warn("", err)
	}
	return str
}
