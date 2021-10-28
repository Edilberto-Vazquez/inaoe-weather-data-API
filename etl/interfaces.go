package etl

import "os"

type File interface {
	ProcessFile(*os.File, string)
}
