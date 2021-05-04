package camelcase

import "testing"

func Test_camelcase(t *testing.T) {
	tests := []struct {
		s      string
		expect int32
	}{
		{
			s:      "saveChangesInTheEditor",
			expect: 5,
		},
		{
			s:      "thisIsTheTestCaseHoHo",
			expect: 7,
		},
	}

	for i, test := range tests {
		if camelcase(test.s) != test.expect {
			t.Errorf("test %d failed", i)
		}
	}
}
