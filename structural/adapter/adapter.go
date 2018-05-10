/*
The Adapter pattern will help yow fit the needs of two parts of the code that are incompatible
at first: two interfaces that are incompatible, but which must work together.
*/

package adapter

import "fmt"

// LegacyPrinter is an old interface
type LegacyPrinter interface {
	Print(string) string
}

// MyLegacyPrinter is a struct that implement LegacyPrinter interface
type MyLegacyPrinter struct{}

func (l *MyLegacyPrinter) Print(s string) (newMsg string) {
	newMsg = fmt.Sprintf("Legacy Printer: %s", s)
	fmt.Println(newMsg)
	return newMsg
}

// ModernPrinter is a new interface
type ModernPrinter interface {
	PrintStored() string
}

// PrinterAdapter is a struct that implements ModernPrinter interface and is an
// adapter for LegacyPrinter interface
type PrinterAdapter struct {
	OldPrinter LegacyPrinter
	Msg        string
}

func (p *PrinterAdapter) PrintStored() (newMsg string) {
	if p.OldPrinter != nil {
		newMsg = fmt.Sprintf("Adapter: %s", p.Msg)
		return p.OldPrinter.Print(newMsg)

	}
	return p.Msg
}
