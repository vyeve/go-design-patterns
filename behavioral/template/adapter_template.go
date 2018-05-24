package template

type adapter struct {
	myFunc func() string
}

func (a *adapter) Message() string {
	if a.myFunc != nil {
		return a.myFunc()
	}
	return ""
}

func MessageRetrieverAdapter(f func() string) MessageRetriever {
	return &adapter{myFunc: f}
}
