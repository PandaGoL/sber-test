package options

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	testTable := []struct {
		val         string
		expected    *Options
		errExpected string
	}{
		{
			val: "api-sber-for-test",
			expected: &Options{
				APIAddr: ":8000",
			},
		},
		{
			val:         "",
			errExpected: "configName is empty",
		},
	}

	for _, testCase := range testTable {
		result, err := LoadConfig(testCase.val)
		if err != nil {
			assert.Equal(t, testCase.expected, result, fmt.Sprintf("Incorrect result. Expected %v, got %v", testCase.expected, result))
		}
		assert.Equal(t, testCase.expected, result, fmt.Sprintf("Incorrect result. Expected %v, got %v", testCase.expected, result))
	}

}
