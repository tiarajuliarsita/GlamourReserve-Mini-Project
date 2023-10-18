package helpers

import (
	"encoding/hex"
	"math/rand"
	"time"
)

var charset = []byte("1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ")

func CreateInvoice(date time.Time) string {

	dateString := date.Format("2006-01-02")
	rand.Seed(time.Now().Unix())
	number := make([]byte, 8)
	for i := range number {
		number[i] = charset[rand.Intn(len(charset))]
	}

	hexString := hex.EncodeToString(number)
	return dateString + hexString
}

func SumTotal(listPrice []int)int{
	var total int
	for _, price := range listPrice {
		total += price
	}
	return total
}