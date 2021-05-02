package maximum_units_on_a_truck

import "testing"

func Test_maximumUnits(t *testing.T) {
	tests := []struct {
		boxTypes  [][]int
		truckSize int
		expect    int
	}{
		{
			boxTypes:  [][]int{{5, 10}, {2, 5}, {4, 7}, {3, 9}},
			truckSize: 10,
			expect:    91,
		},
		{
			boxTypes:  [][]int{{1, 3}, {2, 2}, {3, 1}},
			truckSize: 4,
			expect:    8,
		},
	}

	for i, test := range tests {
		if maximumUnits(test.boxTypes, test.truckSize) != test.expect {
			t.Errorf("test %d failed", i)
		}
	}
}
