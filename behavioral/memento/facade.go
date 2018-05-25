package memento

type MementoFacade struct {
	originator Originator
	careTaker  CareTaker
}

func (f *MementoFacade) SaveSettings(c Command) {
	f.originator.Command = c
	f.careTaker.Add(f.originator.NewMemento())
}

func (f *MementoFacade) RestoreSettings(i int) Command {
	f.originator.ExtractAndStoreCommand(f.careTaker.Memento(i))
	return f.originator.Command
}
