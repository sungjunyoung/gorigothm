package longest_common_prefix

import "testing"

func Test_longestCommonPrefix(t *testing.T) {
	tests := []struct {
		strs   []string
		expect string
	}{
		{
			strs:   []string{"flower", "flow", "flight"},
			expect: "fl",
		},
		{
			strs:   []string{"dog", "racecar", "car"},
			expect: "",
		},
	}

	for _, test := range tests {
		if longestCommonPrefix(test.strs) != test.expect {
			t.Fatal("failed")
		}
	}
}
