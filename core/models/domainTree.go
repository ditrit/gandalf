package models

type DomainTree struct {
	Domain Domain
	Childs []*DomainTree
}
