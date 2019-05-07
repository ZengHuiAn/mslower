package framework

type Server struct {
	name string

	Age int
}

func (s * Server) create() {
	s.name = "123"
}

func NewServer()  {
	s :=&Server{}
	s.name = "123444"
	s.Age = 100
	print(s.name);
}