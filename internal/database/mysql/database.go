package mysql

import (
	"io"
)

type Database interface {
	io.Closer

	//databse methods definition
}
