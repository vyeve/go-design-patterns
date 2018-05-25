package memento

type Command interface {
	GetValue() interface{}
}

type Volume byte

func (v Volume) GetValue() interface{} {
	return v
}

type Mute bool

func (m Mute) GetValue() interface{} {
	return m
}

type Memento struct {
	memento Command
}

type Originator struct {
	Command Command
}

func (o *Originator) NewMemento() Memento {
	return Memento{memento: o.Command}
}

func (o *Originator) ExtractAndStoreCommand(m Memento) {
	o.Command = m.memento
}

type CareTaker struct {
	mementoList []Memento
}

func (c *CareTaker) Add(m Memento) {
	c.mementoList = append(c.mementoList, m)
}

func (c *CareTaker) Pop() Memento {
	if len(c.mementoList) > 0 {
		tempMemento := c.mementoList[len(c.mementoList)-1]
		c.mementoList = c.mementoList[:len(c.mementoList)-1]
		return tempMemento
	}
	return Memento{}
}

func (c *CareTaker) Memento(i int) Memento {
	if len(c.mementoList) == 0 {
		return Memento{}
	}
	if len(c.mementoList) < i+1 {
		return c.mementoList[len(c.mementoList)-1]
	}
	if i < 0 {
		return c.mementoList[0]
	}
	return c.mementoList[i]
}
