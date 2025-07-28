package main

import (
	"fmt"
	"strconv"

	"github.com/shopspring/decimal"
)

func main() {
	inputs := []decimal.Decimal{
		decimal.NewFromFloat(1234),
		decimal.NewFromFloat(33333.75),
	}
	for _, input := range inputs {
		fmt.Println(input)
		fmt.Println(convertToThaiBath(input))
	}
}

func convertToThaiBath(amount decimal.Decimal) string {
	amount = amount.Round(2)
	intPart := amount.Truncate(0)
	decimalValue := amount.Sub(intPart)

	bahtText := NumberToThaiText(intPart.IntPart()) + "บาท"

	if decimalValue.IsZero() {
		bahtText += "ถ้วน"
	} else {
		decimalPart := decimalValue.Mul(decimal.NewFromInt(100)).Round(0)
		if decimalPart.IsZero() {
			bahtText += "ถ้วน"
		} else {
			bahtText += NumberToThaiText(decimalPart.IntPart()) + "สตางค์"
		}
	}

	return bahtText
}

var thaiNums = []string{"ศูนย์", "หนึ่ง", "สอง", "สาม", "สี่", "ห้า", "หก", "เจ็ด", "แปด", "เก้า"}
var thaiUnits = []string{"", "สิบ", "ร้อย", "พัน", "หมื่น", "แสน"}
var thaiBigUnits = []string{"", "ล้าน", "ล้านล้าน"}

func NumberToThaiText(n int64) string {
	if n == 0 {
		return thaiNums[0]
	}

	s := strconv.FormatInt(n, 10)
	length := len(s)
	result := ""

	for i := 0; i < length; i += 6 {
		chunkStart := length - (i + 6)
		if chunkStart < 0 {
			chunkStart = 0
		}
		chunk := s[chunkStart : length-i]
		chunkVal, _ := strconv.ParseInt(chunk, 10, 64)

		if chunkVal == 0 {
			continue
		}

		chunkText := convertSixDigitChunk(chunkVal)
		if i > 0 {
			result = chunkText + thaiBigUnits[i/6] + result
		} else {
			result = chunkText + result
		}
	}

	return result
}
func convertSixDigitChunk(n int64) string {
	if n == 0 {
		return ""
	}

	s := strconv.FormatInt(n, 10)
	length := len(s)
	chunkResult := ""

	for i := 0; i < length; i++ {
		digitChar := s[i]
		digit := int(digitChar - '0')
		pos := length - 1 - i

		if digit == 0 && pos != 0 {
			continue
		}

		if digit == 0 && pos == 0 {
			continue
		}

		numText := thaiNums[digit]
		unitText := thaiUnits[pos]

		if pos == 0 && digit == 1 {
			if length > 1 {
				if length >= 2 && s[length-2] == '1' {
					numText = "เอ็ด"
				} else {
					numText = "เอ็ด"
				}
			}
		}

		if pos == 1 && digit == 2 {
			numText = "ยี่"
		}

		if pos == 1 && digit == 1 {
			numText = ""
		}
		chunkResult += numText + unitText
	}
	return chunkResult
}
