package interfaces

type Response interface {
	HttpStatus() int
	ApplicationStatus() string
	HasData() bool
	GetData() interface{}
}
