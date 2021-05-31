package fpe

import (
	"math/big"
	"reflect"
	"testing"
)

func TestFactor(t *testing.T) {
	tests := []struct {
		name    string
		args    int64
		wantA   int64
		wantB   int64
		wantErr error
	}{
		{"success - 10", 10, 5, 2, nil},
		{"success - 12", 12, 6, 2, nil},
		{"success - 21", 21, 7, 3, nil},
		{"success - 10001", 10001, 137, 73, nil},
		{"success - 9999999", 9999999, 13947, 717, nil},
		{"failure - modulus is prime", 17, 17, 1, nil},
		{"failure - modulus is negative", -1, 0, 0, ErrNegativeArgs},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotA, gotB, err := factor(big.NewInt(tt.args))
			if tt.wantErr != nil {
				if err != tt.wantErr {
					t.Errorf("Factor() err = %v, want %v", err, tt.wantErr)
				}
				return
			}
			if !reflect.DeepEqual(gotA, big.NewInt(tt.wantA)) {
				t.Errorf("Factor() gotA = %v, want %v", gotA, tt.wantA)
			}
			if !reflect.DeepEqual(gotB, big.NewInt(tt.wantB)) {
				t.Errorf("Factor() gotB = %v, want %v", gotB, tt.wantB)
			}
		})
	}
}
