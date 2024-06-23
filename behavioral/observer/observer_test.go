package observer

import (
	"fmt"
	"testing"
)

func TestStockSubscribeAndPublish(t *testing.T) {
	observer := &mockObserver{}
	stock := &stock{name: "Apple"}
	stock.Subscribe(observer)
	stock.SetPrice(10)

	want := "Apple: 10.00"
	got := observer.updates[0]
	if want != got {
		t.Errorf("want observer update '%s', got '%s'", got, want)
	}
}

func TestStockUnubscribeAndPublish(t *testing.T) {
	someObserver := &mockObserver{}
	anotherObserver := &mockObserver{}
	stock := &stock{name: "Apple"}
	stock.Subscribe(someObserver)
	stock.Subscribe(anotherObserver)
	stock.Unsubscribe(anotherObserver)
	unsubsribed := anotherObserver
	stock.SetPrice(10)

	want := "Apple: 10.00"
	got := someObserver.updates[0]
	if want != got {
		t.Errorf("want observer update '%s', got '%s'", got, want)
	}

	if len(unsubsribed.updates) > 0 {
		t.Errorf("want observer without updates, got %s", unsubsribed.updates)
	}
}

type mockObserver struct {
	updates []string
}

func (o *mockObserver) Update(s stock) {
	update := fmt.Sprintf("%s: %.2f", s.name, s.price)
	o.updates = append(o.updates, update)
}
