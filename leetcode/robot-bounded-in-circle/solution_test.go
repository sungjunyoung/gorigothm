package robot_bounded_in_circle

import "testing"

func Test_isRobotBounded(t *testing.T) {
	tests := []struct {
		instructions string
		expect       bool
	}{
		{
			instructions: "GGLLGG",
			expect:       true,
		},
		{
			instructions: "GG",
			expect:       false,
		},
		{
			instructions: "GL",
			expect:       true,
		},
		{
			instructions: "GLRLLGLL",
			expect:       true,
		},
	}

	for _, test := range tests {
		if isRobotBounded(test.instructions) != test.expect {
			t.Fatal("failed")
		}
	}
}
