package interfaces

type HttpResponse interface {
	HttpStatus() int
	ApplicationStatus() string
	HasData() bool
	GetData() interface{}
}
