package deposit

import (
	"fmt"
	"testing"

	"sber-test/internal/services/deposit/types"

	"github.com/stretchr/testify/assert"
)

func TestCalculateDeposit(t *testing.T) {

	s := New()
	type args struct {
		requestID string
		req       *types.DepositeRequest
	}

	testTable := []struct {
		args        args
		expected    []byte
		expectedErr error
	}{
		{
			args: args{
				requestID: "1",
				req: &types.DepositeRequest{
					Date:    "31.12.2023",
					Periods: 12,
					Amount:  20000,
					Rate:    7,
				},
			},
			expected:    []byte("{\"31.12.2023\": 20116.67,\"31.01.2024\": 20234.01,\"29.02.2024\": 20352.05,\"31.03.2024\": 20470.77,\"30.04.2024\": 20590.18,\"31.05.2024\": 20710.29,\"30.06.2024\": 20831.10,\"31.07.2024\": 20952.61,\"31.08.2024\": 21074.84,\"30.09.2024\": 21197.77,\"31.10.2024\": 21321.43,\"30.11.2024\": 21445.80}"),
			expectedErr: nil,
		},
		{
			args: args{
				requestID: "1",
				req: &types.DepositeRequest{
					Date:    "31.12.2023",
					Periods: 3,
					Amount:  2000000,
					Rate:    7,
				},
			},
			expected:    []byte("{\"31.12.2023\": 2011666.67,\"31.01.2024\": 2023401.39,\"29.02.2024\": 2035204.56}"),
			expectedErr: nil,
		},
		{
			args: args{
				requestID: "2",
				req: &types.DepositeRequest{
					Date:    "31/12/2023",
					Periods: 12,
					Amount:  20000,
					Rate:    7,
				},
			},
			expected:    nil,
			expectedErr: types.ErrValidation,
		},
		{
			args: args{
				requestID: "3",
				req: &types.DepositeRequest{
					Date:    "31.12.2023",
					Periods: 0,
					Amount:  20000,
					Rate:    7,
				},
			},
			expected:    nil,
			expectedErr: types.ErrValidation,
		},
		{
			args: args{
				requestID: "4",
				req: &types.DepositeRequest{
					Date:    "31.12.2023",
					Periods: 60,
					Amount:  20000,
					Rate:    7,
				},
			},
			expected:    nil,
			expectedErr: types.ErrValidation,
		},
		{
			args: args{
				requestID: "5",
				req: &types.DepositeRequest{
					Date:    "31.12.2023",
					Periods: 12,
					Amount:  9999,
					Rate:    7,
				},
			},
			expected:    nil,
			expectedErr: types.ErrValidation,
		},
		{
			args: args{
				requestID: "6",
				req: &types.DepositeRequest{
					Date:    "31.12.2023",
					Periods: 12,
					Amount:  3000000,
					Rate:    7,
				},
			},
			expected:    nil,
			expectedErr: types.ErrValidation,
		},
		{
			args: args{
				requestID: "7",
				req: &types.DepositeRequest{
					Date:    "31.12.2023",
					Periods: 12,
					Amount:  20000,
					Rate:    0,
				},
			},
			expected:    nil,
			expectedErr: types.ErrValidation,
		},
		{
			args: args{
				requestID: "5",
				req: &types.DepositeRequest{
					Date:    "31.12.2023",
					Periods: 12,
					Amount:  20000,
					Rate:    8,
				},
			},
			expected:    nil,
			expectedErr: types.ErrValidation,
		},
	}

	for _, tt := range testTable {

		got, err := s.CalculateDeposit(tt.args.requestID, tt.args.req)
		if err != nil {
			assert.Equal(t, tt.expectedErr, err, fmt.Sprintf("Incorrect result. Expected %v, got %v", tt.expectedErr, err))
		}
		assert.Equal(t, tt.expected, got, fmt.Sprintf("Incorrect result. Expected %v, got %v", tt.expected, got))
	}
}

func TestColculateAmount(t *testing.T) {

	testTable := []struct {
		amount   float64
		rate     float64
		expected float64
	}{
		{
			amount: 10000,
			rate:   6,
		},
		{
			amount: 20000,
			rate:   10,
		},
		{
			amount: 20133.33,
			rate:   8,
		},
	}

	for _, testCase := range testTable {
		result := calcAmount(testCase.amount, testCase.rate)
		assert.Equal(t, testCase.amount*(1+testCase.rate/12/100), result, fmt.Sprintf("Incorrect result. Expected %v, got %v", testCase.amount*(1+testCase.rate/12/100), result))
	}
}

func TestCalculateDepositFunc(t *testing.T) {

	testTable := []struct {
		req           *types.DepositeRequest
		expected      []byte
		expectedError error
	}{
		{
			req: &types.DepositeRequest{
				Date:    "31.01.2021",
				Periods: 1,
				Amount:  10000,
				Rate:    6,
			},
			expected: []byte("{\"31.01.2021\": 10050}"),
		},
		{
			req: &types.DepositeRequest{
				Date:    "31.12.2023",
				Periods: 12,
				Amount:  20000,
				Rate:    8,
			},
			expected: []byte("{\"31.12.2023\": 20133.33,\"31.01.2024\": 20267.56,\"29.02.2024\": 20402.67,\"31.03.2024\": 20538.69,\"30.04.2024\": 20675.62,\"31.05.2024\": 20813.45,\"30.06.2024\": 20952.21,\"31.07.2024\": 21091.89,\"31.08.2024\": 21232.50,\"30.09.2024\": 21374.05,\"31.10.2024\": 21516.55,\"30.11.2024\": 21659.99}"),
		},
	}

	for _, testCase := range testTable {
		result := calculateDeposit(testCase.req)

		assert.Equal(t, testCase.expected, result, fmt.Sprintf("Incorrect result. Expected %v, got %v", testCase.expected, result))
	}
}

func TestConvertFloatRes(t *testing.T) {

	testTable := []struct {
		req      float64
		expected string
	}{
		{
			req:      3000000,
			expected: "3000000",
		},
		{
			req:      3000000.11,
			expected: "3000000.11",
		},
	}

	for _, testCase := range testTable {
		result := convertFloatToRes(testCase.req)

		assert.Equal(t, testCase.expected, result, fmt.Sprintf("Incorrect result. Expected %v, got %v", testCase.expected, result))
	}
}
