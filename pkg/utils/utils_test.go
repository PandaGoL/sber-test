package utils

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRoundFloat(t *testing.T) {

	testTable := []struct {
		val       float64
		precision uint
		expected  float64
	}{
		{
			val:       1.111111,
			precision: 2,
			expected:  1.11,
		},
		{
			val:       3.1455555,
			precision: 2,
			expected:  3.15,
		},
		{
			val:       3.104444,
			precision: 2,
			expected:  3.10,
		},
	}

	for _, testCase := range testTable {
		result := RoundFloat(testCase.val, testCase.precision)
		assert.Equal(t, testCase.expected, result, fmt.Sprintf("Incorrect result. Expected %v, got %v", testCase.expected, result))
	}

}

func TestAddMonth(t *testing.T) {

	layout := "02.01.2006"

	testDate1, _ := time.Parse(layout, "31.12.2021")
	testDate2, _ := time.Parse(layout, "14.01.2021")
	testDate3, _ := time.Parse(layout, "01.06.2021")

	testResp1, _ := time.Parse(layout, "28.02.2022")
	testResp2, _ := time.Parse(layout, "14.02.2021")
	testResp3, _ := time.Parse(layout, "01.06.2022")

	testTable := []struct {
		date     time.Time
		m        int
		expected time.Time
	}{
		{
			date:     testDate1,
			m:        2,
			expected: testResp1,
		},
		{
			date:     testDate2,
			m:        1,
			expected: testResp2,
		},
		{
			date:     testDate3,
			m:        12,
			expected: testResp3,
		},
	}

	for _, testCase := range testTable {
		result := AddMonth(testCase.date, testCase.m)
		assert.Equal(t, testCase.expected, result, fmt.Sprintf("Incorrect result. Expected %v, got %v", testCase.expected, result))
	}

}
