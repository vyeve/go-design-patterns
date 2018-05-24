package template

import (
	"strings"
	"testing"
)

type TestStruct struct {
	TemplateImpl
}

func (m *TestStruct) Message() string {
	return "world"
}

func TestTemplate_ExecuteAlgorithm(t *testing.T) {
	t.Run("Using interfaces", func(t *testing.T) {
		s := &TestStruct{}
		res := s.ExecuteAlgorithm(s)
		expectedOrError(res, " world", t)
	})

	t.Run("Using anonymous functions", func(t *testing.T) {
		m := new(AnonymousTemplate)
		res := m.ExecuteAlgorithm(func() string {
			return "world"
		})
		expectedOrError(res, " world", t)
	})

	t.Run("Using anonymous functions adapted to an interface", func(t *testing.T) {
		messageRetriever := MessageRetrieverAdapter(func() string {
			return "world"
		})

		if messageRetriever == nil {
			t.Fatal("cannot continue with a nil MessageRetriever")
		}
		template := TestStruct{}
		res := template.ExecuteAlgorithm(messageRetriever)
		expectedOrError(res, "world", t)
	})
}

func expectedOrError(res string, expected string, t *testing.T) {
	if !strings.Contains(res, expected) {
		t.Errorf("expected string '%s' was not found on returned string\n", expected)
	}
}
