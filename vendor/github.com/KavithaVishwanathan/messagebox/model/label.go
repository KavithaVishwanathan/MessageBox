package model

type LabelType int

const (
	Personal LabelType = iota
	Purchases
	Promotions
)

var LabelNames = [...]string{"Personal","Purchases","Promotions"}

type Label struct {
	Id	LabelType	`json:"id"`
	Name	string	`json:"name"`
}

func CreateLabels() []Label{
	labels := []Label {
		Label{Personal,LabelNames[Personal]},
		Label{Purchases, LabelNames[Purchases]},
		Label{Promotions,LabelNames[Promotions]},
	}
	return labels
}



