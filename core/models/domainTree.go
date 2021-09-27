package models

type DomainTree struct {
	Name   string
	Childs []*DomainTree
}
