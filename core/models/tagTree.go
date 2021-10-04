package models

type TagTree struct {
	Tag    Tag
	Childs []*TagTree
}
