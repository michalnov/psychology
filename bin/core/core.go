package core

import (
	"github.com/michalnov/psychology/bin/core/structures"

	_ "github.com/go-sql-driver/mysql" //needed
)

//Core -
type Core struct {
	DbMaster string
	DbSlave  string
	Tests    []structures.Test
}

//LoadTests -
func LoadTests() error {

	return nil
}
