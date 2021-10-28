package etl

type FilesTypes interface {
	ProcessFile(string)
	GetRecords() []interface{}
}
