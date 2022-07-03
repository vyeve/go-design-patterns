package singleton

import "testing"

func TestGetInstance(t *testing.T) {
	counter1 := GetInstance()
	if counter1 == nil {
		// Test of acceptance criteria 1 failed
		t.Error("expected pointer to Singleton after calling GetInstance(), not nil")
	}

	expectedCounter := counter1
	expectedCount := 1
	currentCount := counter1.AddOne()
	if currentCount != expectedCount {
		t.Errorf("After calling for the first time to count, the count must be %d but it is %d\n",
			expectedCount, currentCount)
	}

	counter2 := GetInstance()
	if counter2 != expectedCounter {
		// Test 2 failed
		t.Error("Expected same instance in counter2 but it got a different instance")
	}

	counter2.AddOne()
	currentCount = counter2.GetCount()
	expectedCount++
	if currentCount != expectedCount {
		t.Errorf("After calling 'AddOne' using the second counter, the current count must be %d but was %d\n",
			expectedCount, currentCount)
	}
}
