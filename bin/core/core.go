package core

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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

type user struct {
	Kluc  string `json:"kluc"`
	Vek   string `json:"vek"`
	Rod   string `json:"rod"`
	Skola string `json:"skola"`
}

//UserHandler --
func (c *Core) UserHandler(w http.ResponseWriter, r *http.Request) {
	var req user
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(300)
		fmt.Fprintf(w, "{\"status\" : \"error\"}")
	}
	db, err := sql.Open("mysql", c.DbMaster)
	if err != nil {
		w.WriteHeader(300)
		fmt.Fprintf(w, "{\"status\" : \"error\"}")
	}
	defer db.Close()
	statement, err := db.Prepare("insert into user(age,gender,school,key) OUTPUT Inserted.ID values(?,?,?,?)")
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "{\"status\" : \"error\"}")
	}
	var id int
	err = statement.QueryRow(req.Vek, req.Rod, req.Skola, req.Kluc).Scan(&id)
	w.WriteHeader(200)
	fmt.Fprintf(w, "{\"user\" : \""+string(id)+"\"}")
}

func (c *Core) GetTest(w http.ResponseWriter, r *http.Request) {
	var req string
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
