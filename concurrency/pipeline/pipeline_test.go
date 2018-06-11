package pipeline

import "testing"

func TestLaunchPipeline(t *testing.T) {
	tableTest := [][]int{
		{3, 14},
		{4, 30},
		{5, 55},
	}

	for _, test := range tableTest {
		if res := LaunchPipeline(test[0]); res != test[1] {
			t.Fatal()
		} else {
			t.Logf("%d == %d", res, test[1])
		}
	}
}
