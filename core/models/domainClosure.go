package models

//TODO ADD FOREIGN KEY
type DomainClosure struct {
	AncestorID   uint
	Ancestor     Domain
	DescendantID uint
	Descendant   Domain
}
