package models

type TagClosure struct {
	AncestorID   uint
	Ancestor     Tag
	DescendantID uint
	Descendant   Tag
	Depth        uint
}
