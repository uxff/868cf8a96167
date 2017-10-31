// Autogenerated by Thrift Compiler (0.10.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package tcliservice

import (
	"bytes"
	"fmt"

	"git.apache.org/thrift.git/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

var PRIMITIVE_TYPES map[TTypeId]struct{}
var COMPLEX_TYPES map[TTypeId]struct{}
var COLLECTION_TYPES map[TTypeId]struct{}
var TYPE_NAMES map[TTypeId]string

const CHARACTER_MAXIMUM_LENGTH = "characterMaximumLength"
const PRECISION = "precision"
const SCALE = "scale"

func init() {
	PRIMITIVE_TYPES = map[TTypeId]struct{}{
		0:  struct{}{},
		1:  struct{}{},
		2:  struct{}{},
		3:  struct{}{},
		4:  struct{}{},
		5:  struct{}{},
		6:  struct{}{},
		7:  struct{}{},
		8:  struct{}{},
		9:  struct{}{},
		15: struct{}{},
		16: struct{}{},
		17: struct{}{},
		18: struct{}{},
		19: struct{}{},
		20: struct{}{},
		21: struct{}{},
	}

	COMPLEX_TYPES = map[TTypeId]struct{}{
		10: struct{}{},
		11: struct{}{},
		12: struct{}{},
		13: struct{}{},
		14: struct{}{},
	}

	COLLECTION_TYPES = map[TTypeId]struct{}{
		10: struct{}{},
		11: struct{}{},
	}

	TYPE_NAMES = map[TTypeId]string{
		0:  "BOOLEAN",
		1:  "TINYINT",
		2:  "SMALLINT",
		3:  "INT",
		4:  "BIGINT",
		5:  "FLOAT",
		6:  "DOUBLE",
		7:  "STRING",
		8:  "TIMESTAMP",
		9:  "BINARY",
		10: "ARRAY",
		11: "MAP",
		12: "STRUCT",
		13: "UNIONTYPE",
		15: "DECIMAL",
		16: "NULL",
		17: "DATE",
		18: "VARCHAR",
		19: "CHAR",
		20: "INTERVAL_YEAR_MONTH",
		21: "INTERVAL_DAY_TIME",
	}

}
