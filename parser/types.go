package parser

type Service struct {
	Annotations []*Annotation
}

type Annotation struct {
	Name  string
	Value string
}
