package template

import "strings"

type AnonymousTemplate struct{}

func (*AnonymousTemplate) first() string {
	return "hello"
}

func (*AnonymousTemplate) third() string {
	return "template"
}

func (a *AnonymousTemplate) ExecuteAlgorithm(f func() string) string {
	return strings.Join([]string{a.first(), f(), a.third()}, " ")
}
