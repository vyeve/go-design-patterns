package publisher_subscriber

type Subscriber interface {
	Notify(interface{}) error
	Close()
}
