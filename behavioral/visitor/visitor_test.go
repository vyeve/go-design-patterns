package visitor

import "testing"

type testHelper struct {
	Received string
}

func (h *testHelper) Write(p []byte) (int, error) {
	h.Received = string(p)
	return len(p), nil
}

func Test_Overall(t *testing.T) {
	helper := &testHelper{}
	visitor := &MessageVisitor{}

	t.Run("MessageA test", func(t *testing.T) {
		msg := MessageA{
			Msg:    "Hello World",
			Output: helper,
		}

		msg.Accept(visitor)
		msg.Print()

		expected := "A: Hello World (Visited A)"

		if helper.Received != expected {
			t.Errorf("expected result incorrect. %s != %s", helper.Received, expected)
		}
	})

	t.Run("MessageB test", func(t *testing.T) {
		msg := MessageB{
			Msg:    "Hello World",
			Output: helper,
		}

		msg.Accept(visitor)
		msg.Print()

		expected := "B: Hello World (Visited B)"

		if helper.Received != expected {
			t.Errorf("expected result incorrect. %s != %s", helper.Received, expected)
		}
	})
}
