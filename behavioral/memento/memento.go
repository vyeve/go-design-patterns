/*
The Memento design pattern usually has three players:
 - Memento: A type that stores the type we want to save.
 - Originator: A type that is in change of creating mementos and storing the current
active state.
 - Care Taker: A type that stores the list of mementos.
*/

package memento

import "errors"

type State struct {
	Description string
}

type memento struct {
	state State
}

type originator struct {
	state State
}

func (o *originator) NewMemento() memento {
	return memento{state: o.state}
}

func (o *originator) ExtractAndStoreState(m memento) {
	o.state = m.state
}

type careTaker struct {
	mementoList []memento
}

func (c *careTaker) Add(m memento) {
	c.mementoList = append(c.mementoList, m)
}

func (c *careTaker) Memento(i int) (memento, error) {
	if len(c.mementoList) < i || i < 0 {
		return memento{}, errors.New("index not found")
	}
	return c.mementoList[i], nil
}
