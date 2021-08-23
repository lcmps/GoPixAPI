// Copyright (c) 2021 Jonnas Fonini

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
package app

import (
	"errors"
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"unicode/utf8"

	"github.com/lcmps/gopix/models"
	"github.com/r10r/crc16"
)

type intMap map[int]interface{}

func GeneratePaste(amt float64, name, city, descpt, txID, pixKey string) (string, error) {
	opts := models.PixOpts{
		Name:          name,
		City:          city,
		Description:   descpt,
		TransactionID: txID,
		Amount:        amt,
		Key:           pixKey,
	}

	paste, err := Pix(opts)
	if err != nil {
		return paste, err
	}
	return paste, nil
}

// Pix generates a Copy and Paste Pix code
func Pix(options models.PixOpts) (string, error) {
	if err := validateData(options); err != nil {
		return "", err
	}
	data := buildDataMap(options)
	str := parseData(data)
	// Add the CRC at the end
	str += "6304"
	crc, err := calculateCRC16(str)
	if err != nil {
		return "", err
	}
	str += crc
	return str, nil
}

func validateData(options models.PixOpts) error {
	if options.Key == "" {
		return errors.New("key must not be empty")
	}
	if options.Name == "" {
		return errors.New("name must not be empty")
	}
	if options.City == "" {
		return errors.New("city must not be empty")
	}
	if utf8.RuneCountInString(options.Name) > 25 {
		return errors.New("name must be at least 25 characters long")
	}
	if utf8.RuneCountInString(options.City) > 15 {
		return errors.New("city must be at least 15 characters long")
	}
	return nil
}

func buildDataMap(options models.PixOpts) intMap {
	data := make(intMap)
	// Payload Format Indicator
	data[0] = "01"
	// Merchant Account Information
	data[26] = intMap{0: "BR.GOV.BCB.PIX", 1: options.Key, 2: options.Description}
	// Merchant Category Code
	data[52] = "0000"
	// Transaction Currency - Brazilian Real - ISO4217
	data[53] = "986"
	// Transaction Amount
	data[54] = options.Amount
	// Country Code - ISO3166-1 alpha 2
	data[58] = "BR"
	// Merchant Name. 25 characters maximum
	data[59] = options.Name
	// Merchant City. 15 characters maximum
	data[60] = options.City
	// Transaction ID
	data[62] = intMap{5: "***", 50: intMap{0: "BR.GOV.BCB.BRCODE", 1: "1.0.0"}}
	if options.TransactionID != "" {
		data[62].(intMap)[5] = options.TransactionID
	}
	return data
}

func parseData(data intMap) string {
	var str string

	keys := sortKeys(data)
	for _, k := range keys {
		v := reflect.ValueOf(data[k])
		switch v.Kind() {
		case reflect.String:
			value := data[k].(string)
			str += fmt.Sprintf("%02d%02d%s", k, len(value), value)
		case reflect.Float64:
			value := strconv.FormatFloat(v.Float(), 'f', 2, 64)
			str += fmt.Sprintf("%02d%02d%s", k, len(value), value)
		case reflect.Map:
			// If the element is another map, do a recursive call
			content := parseData(data[k].(intMap))
			str += fmt.Sprintf("%02d%02d%s", k, len(content), content)
		}
	}
	return str
}

func sortKeys(data intMap) []int {
	keys := make([]int, len(data))
	i := 0

	for k := range data {
		keys[i] = k
		i++
	}
	sort.Ints(keys)
	return keys
}

func calculateCRC16(str string) (string, error) {
	table := crc16.MakeTable(crc16.CRC16_CCITT_FALSE)
	h := crc16.New(table)
	_, err := h.Write([]byte(str))
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%04X", h.Sum16()), nil
}
