package models

type ObjectClosure struct {
	AncestorID   uint
	Ancestor     Domain
	DescendantID uint
	Descendant   Domain
	Depth        uint
}
