package faraday

type FieldType int

const (
	Segment FieldType = iota
	SegGroup
)

type StructField struct {
	Idx int
	Typ FieldType
}
