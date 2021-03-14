package interfaces

type (
	Effect interface {
		IsPermit() bool
		IsEmpty() bool
	}
)
