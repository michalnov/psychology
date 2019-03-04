package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"

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
func (c *Core) LoadTests() error {
	var swap structures.Test
	swap, err := readTest("test1")
	if err != nil {
		return err
	}
	c.Tests = append(c.Tests, swap)
	swap, err = readTest("test2")
	if err != nil {
		return err
	}
	c.Tests = append(c.Tests, swap)
	swap, err = readTest("test3")
	if err != nil {
		return err
	}
	c.Tests = append(c.Tests, swap)
	return nil
}

//readTest -
func readTest(name string) (structures.Test, error) {
	out := structures.Test{}
	absPath, _ := filepath.Abs(name + ".json")
	fmt.Println("Database Opening configuration File")
	temFile, err := ioutil.ReadFile(absPath)
	if err != nil {
		fmt.Println("Reading Db config failed")
		return out, err
	}
	err = json.Unmarshal(temFile, &out)
	if err != nil {
		fmt.Println("Unmarshal config failed")
		return out, err
	}
	return out, nil
}
