package entity

const (
	SkipLabel LabelOperation = 1 << iota
	CreateLabel
	DeleteLabel
	UpdateLabel
)

// Label represents a GitHub label.
type Label struct {
	ID    int64
	Name  string
	Color string
	Desc  string
}

type LabelOperation uint

type LabelTransform struct {
	Operation LabelOperation
	From      Label
	To        Label
}
