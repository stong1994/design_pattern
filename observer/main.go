package main

type Observer interface {
	Update(msg string)
}

type Observer1 struct{}

func (n Observer1) Update(msg string) {
	// do something
}

type Observer2 struct{}

func (n Observer2) Update(msg string) {
	// do something
}

type Observer3 struct{}

func (n Observer3) Update(msg string) {
	// do something
}

type Topic struct {
	msg       string
	observers []Observer
}

func (t *Topic) Notify() {
	for _, observer := range t.observers {
		observer.Update(t.msg)
	}
}
