package connector

import (
	"container/list"
	"fmt"
	"gandalf-go/message"
	"strconv"
	"sync"
	"time"
)

//Queue : queue allowing access via a string key
type Queue struct {
	qlist list.List
	dict  map[string]*list.Element
	iters map[*Iterator]bool
	m     sync.Mutex
}

// NewQueue : constructor
func NewQueue() *Queue {
	q := new(Queue)
	q.Init()
	return q
}

// Init :
func (q *Queue) Init() {
	q.qlist.Init()
	q.dict = make(map[string]*list.Element)
	q.iters = make(map[*Iterator]bool)
}

// Push : insert a new value in the queue except if the UUID is already present and remove after timeout expiration
func (q *Queue) Push(m message.Message) {
	fmt.Println("Push a message!")
	fmt.Println(m)

	key := m.GetUUID()
	timeout, _ := strconv.Atoi(m.GetTimeout())
	fmt.Println("TIME OUT")
	fmt.Println(timeout)

	q.m.Lock()
	defer q.m.Unlock()

	ele := q.dict[key]
	if ele != nil {
		return
	}

	ele = q.qlist.PushFront(m)
	q.dict[key] = ele
	go func() {
		time.Sleep(time.Duration(timeout) * time.Millisecond)
		fmt.Println("REMOVED")
		q.remove(key)
	}()
}

// First :
func (q *Queue) First() *message.Message {
	q.m.Lock()
	defer q.m.Unlock()

	ele := q.qlist.Back()
	if ele != nil {
		value := ele.Value.(message.Message)
		return &value
	}
	return nil
}

// Next :
func (q *Queue) Next(key string) *message.Message {
	q.m.Lock()
	defer q.m.Unlock()
	eleFromKey := q.dict[key]
	if eleFromKey != nil {
		nextEle := eleFromKey.Prev()
		if nextEle != nil {
			nextMessage := nextEle.Value.(message.Message)
			return &nextMessage
		}
	}
	return nil
}

// remove :
func (q *Queue) remove(key string) {
	q.m.Lock()
	defer q.m.Unlock()

	// Repositionner les iterateurs positionnés sur le message à supprimer
	// sur le message :
	// 1. suivant s'il existe
	// 2. sinon sur le précédent s'il existe
	// 3. sinon c'est que la queue est vide
	ele := q.dict[key]
	nextEle := ele.Prev() // cas 1.
	if nextEle == nil {
		nextEle = ele.Next() // cas 2.
	}
	nextUUID := ""
	if nextEle != nil {
		nextUUID = nextEle.Value.(message.Message).GetUUID()
	}

	// suppression et repositionnement pour chaque iterateur
	for i := range q.iters {
		i.m.Lock()
		if i.current == key { // si literateur pointe sur le message à supprimer
			i.current = nextUUID // repositionnement
		}
		// supprimer le message de la liste des messages déjà consultés
		//delete(i.seen, key)
		i.m.Unlock()
	}

	// supprimer le message dans la queue (dans la liste et dans la map)
	delete(q.dict, key)
	q.qlist.Remove(ele)
}

// IsEmpty : the event queue is empty
func (q *Queue) IsEmpty() bool {
	return q.qlist.Len() == 0
}

// Print :
func (q *Queue) Print() {
	ele := q.qlist.Back()
	fmt.Printf("   Queue{\n")
	for ele != nil {
		fmt.Printf("      %s,\n", ele.Value.(message.Message).GetUUID())
		ele = ele.Prev()
	}
	fmt.Printf("nb eles : %d\n", q.qlist.Len())
}
