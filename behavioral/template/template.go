/*
The Template design pattern lets the user write a part of an algorithm while the rest is
executed by the abstraction. This is common when creating libraries to ease in some complex
task or when reusability of some algorithm is compromised by only a part of it.
*/

package template

import "strings"

type MessageRetriever interface {
	Message() string
}

type Template interface {
	first() string
	third() string
	ExecuteAlgorithm(MessageRetriever) string
}

type TemplateImpl struct{}

func (t *TemplateImpl) first() string {
	return "hello"
}

func (t *TemplateImpl) third() string {
	return "template"
}

func (t *TemplateImpl) ExecuteAlgorithm(m MessageRetriever) string {
	return strings.Join([]string{t.first(), m.Message(), t.third()}, " ")
}
