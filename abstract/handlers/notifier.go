package handlers

type Notifier interface {
	Notify(message string) error
}

