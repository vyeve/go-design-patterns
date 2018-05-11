/*
The Bridge pattern tries to  decouple abstraction (an object) from its implementation (the thing that
the object does). The object of the Bridge pattern is to bring flexibility to a struct that change often.
*/

package bridge

import (
	"errors"
	"fmt"
	"io"
)

type PrinterAPI interface {
	PrintMessage(string) error
}

type PrinterImpl1 struct{}

func (p *PrinterImpl1) PrintMessage(msg string) error {
	_, err := fmt.Println(msg)
	return err
}

type PrinterImpl2 struct {
	Writer io.Writer
}

func (d *PrinterImpl2) PrintMessage(msg string) error {
	if d.Writer == nil {
		return errors.New("you need to pass an io.Writer to PrinterImpl2")
	}
	_, err := fmt.Fprint(d.Writer, msg)
	return err
}

type PrinterAbstraction interface {
	Print() error
}

type NormalPrinter struct {
	Msg     string
	Printer PrinterAPI
}

func (c *NormalPrinter) Print() error {
	return c.Printer.PrintMessage(c.Msg)
}

type PacktPrinter struct {
	Msg     string
	Printer PrinterAPI
}

func (c *PacktPrinter) Print() error {
	return c.Printer.PrintMessage(fmt.Sprintf("Message from Packt: %s", c.Msg))
}
