package deck

import "fmt"

type ReviewQueue struct {
	queue []*ShittyCard
}

func NewReviewQueue(max int) *ReviewQueue {
	return &ReviewQueue{
		queue: make([]*ShittyCard, 0, max),
	}
}

func (q *ReviewQueue) Push(card *ShittyCard) {
	q.queue = append(q.queue, card)
}

func (q *ReviewQueue) Pop() (*ShittyCard, error) {
	if len(q.queue) == 0 {
		return nil, fmt.Errorf("queue is empty")
	}

	card := q.queue[0]
	q.queue = q.queue[1:]
	return card, nil
}

func (q *ReviewQueue) Peek() *ShittyCard {
	return q.queue[0]
}

func (q *ReviewQueue) Size() int {
	return len(q.queue)
}

func (q *ReviewQueue) Clear() {
	q.queue = make([]*ShittyCard, 0)
}

func (q *ReviewQueue) Has(card *ShittyCard) bool {
	for _, c := range q.queue {
		if c.content == card.content && c.kind == card.kind {
			return true
		}
	}

	return false
}
