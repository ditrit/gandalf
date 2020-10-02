package models

type DomainClosure struct {
	AncestorID   uint
	Ancestor     Domain
	DescendantID uint
	Descendant   Domain
	Depth        uint
}
