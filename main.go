package main

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func main() {
	inputs := []decimal.Decimal{
		decimal.NewFromFloat(1234),
		decimal.NewFromFloat(33333.75),
	}
	for _, input := range inputs {
		fmt.Println(input)

		// convert decimals to thai text (bath) and print the result here
		fmt.Println(convertToThaiBath(input))

	}
}

func convertToThaiBath(amount decimal.Decimal) string {
	intPart := amount.Truncate(0)
	decimalPart := amount.Sub(intPart).Mul(decimal.NewFromInt(100)).Round(0)

	bahtText := NumberToThaiText(intPart.IntPart()) + "บาท"

	if decimalPart.IsZero() {
		bahtText += "ถ้วน"
	} else {
		bahtText += NumberToThaiText(decimalPart.IntPart()) + "สตางค์"
	}

	return bahtText

}

var thaiNums = []string{"ศูนย์", "หนึ่ง", "สอง", "สาม", "สี่", "ห้า", "หก", "เจ็ด", "แปด", "เก้า"}
var thaiUnits = []string{"", "สิบ", "ร้อย", "พัน", "หมื่น", "แสน", "ล้าน"}

func NumberToThaiText(n int64) string {
	if n == 0 {
		return thaiNums[0]
	}

	result := ""
	numStr := fmt.Sprintf("%d", n)
	length := len(numStr)

	for i, r := range numStr {
		digit := r - '0'
		pos := length - i - 1

		if digit == 0 {
			continue
		}

		if pos == 1 && digit == 2 {
			result += "ยี่"
		} else if pos == 1 && digit == 1 {
			result += ""
		} else if pos == 0 && digit == 1 && length > 1 {
			result += "เอ็ด"
		} else {
			result += thaiNums[digit]
		}

		result += thaiUnits[pos]
	}

	return result
}
