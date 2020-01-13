package connector

import (
	"sync"
	"gandalf-go/message"
)

//Iterator : queue allowing access via a string key
type Iterator struct {
	queue *Queue
	//seen    map[string]bool
	current string
	m       sync.Mutex
}

// NewIterator : constructor
func NewIterator(queue *Queue) *Iterator {
	i := new(Iterator)
	i.Init(queue)
	queue.iters[i] = true
	return i
}

// Init : initialisation
func (i *Iterator) Init(queue *Queue) {
	i.queue = queue
	//i.seen = make(map[string]bool)
}

// Close : fermeture de l'iterateur
func (i *Iterator) Close() {
	delete(i.queue.iters, i)
}

// Get : get next unseen element
func (i *Iterator) Get() *message.Message {
	i.m.Lock()
	defer i.m.Unlock()

	var message *message.Message
	// Si la queue est vide, on ne renvoie rien
	if i.queue.IsEmpty() {
		return nil
	}

	// Si l'iterateur n'a pas été initialisé,
	if i.current == "" {
		message = i.queue.First() // premier message de la queue
	} else {
		message = i.queue.Next(i.current) // message suivant
	}

	// si on a trouvé un nouveau message à renvoyer
	if message != nil {
		i.current = (*message).GetUUID() // on pointe dessus
	}
	return message

	/*
		// on cherche une valeur qui n'a pas déjà été lue
		for i.seen[i.current] {
			message = i.queue.Next(i.current)
			if message == nil {
			return nil
			}
		}
	*/

	// la valeur a été vue
	//i.seen[i.current] = true

	//return message
}
