package queue

import "testing"

func TestAddQueue(t *testing.T) {
	q := New(3)

	for i := 0; i < len(q.items); i++ {
		if len(q.items) != i {
			t.Errorf("Incorrect queue element count: %v, want %v", len(q.items), i)
		}

		if !q.Append(i) {
			t.Errorf("Failed to append item %v to queue", i)
		}
	}

	if q.Append(4) {
		t.Errorf("Should not be able to append item to full queue")
	}
}

func TestNextQueue(t *testing.T) {
	q := New(3)

	for i := 0; i < len(q.items); i++ {
		q.Append(i)
	}

	for i := 0; i < len(q.items); i++ {
		item, ok := q.Next()

		if !ok {
			t.Errorf("Should be able to gran next item from queue")
		}

		if item != i {
			t.Errorf("Got item from queue in wrong order: %v, want %v", item, i)
		}
	}

	item, ok := q.Next()
	if ok {
		t.Errorf("Should not be able to grab item from empty queue, received %v", item)
	}
}