package core

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

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

//Ping --
func (c *Core) Ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	fmt.Fprintf(w, "{\"status\" : \"ok\"}")
}

//UserHandler --
func (c *Core) UserHandler(w http.ResponseWriter, r *http.Request) {
	var req user
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(300)
		fmt.Fprintf(w, "{\"status\" : \"error\"}")
		panic(err)
		return
	}
	db, err := sql.Open("mysql", c.DbMaster)
	if err != nil {
		w.WriteHeader(300)
		fmt.Fprintf(w, "{\"status\" : \"error\"}")
		panic(err)
		return
	}
	defer db.Close()
	statement, err := db.Prepare("insert into user(age,gender,school,testkey) values(?,?,?,?)")
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "{\"status\" : \"error\"}")
		panic(err)
		return
	}
	var id int
	_ = statement.QueryRow(req.Vek, req.Rod, req.Skola, req.Kluc)
	statement2, err := db.Prepare("SELECT LAST_INSERT_ID")
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "{\"status\" : \"error\"}")
		panic(err)
		return
	}
	err = statement2.QueryRow().Scan(&id)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "{\"status\" : \"error\"}")
		panic(err)
		return
	}
	w.WriteHeader(200)
	fmt.Fprintf(w, "{\"user\" : \""+strconv.Itoa(id)+"\"}")
	panic(err)
	return
}

type getTest struct {
	Test   string `json:"test"`
	UserID int    `json:"userid"`
}

//GetTest --
func (c *Core) GetTest(w http.ResponseWriter, r *http.Request) {
	var req getTest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(300)
		fmt.Fprintf(w, "{\"status\" : \"error\"}")
		panic(err)
		return
	}
	for _, test := range c.Tests {
		if test.Key == req.Test {
			resp, err := json.MarshalIndent(test, "  ", "    ")
			if err != nil {
				w.WriteHeader(500)
				fmt.Fprintf(w, "{\"status\" : \"marshal error\"}")
				panic(err)
				return
			}
			db, err := sql.Open("mysql", c.DbMaster)
			if err != nil {
				w.WriteHeader(300)
				fmt.Fprintf(w, "{\"status\" : \"mysql error\"}")
				panic(err)
				return
			}
			defer db.Close()
			statement, err := db.Prepare("update user set testid = ? where iduser = ?")
			if err != nil {
				w.WriteHeader(500)
				fmt.Fprintf(w, "{\"status\" : \"sql error\"}")
				panic(err)
				return
			}
			_, err = statement.Exec(test.TestID, req.UserID)
			if err != nil {
				w.WriteHeader(500)
				fmt.Fprintf(w, "{\"status\" : \"sql 2 error\"}")
				panic(err)
				return
			}
			w.WriteHeader(200)
			fmt.Fprintf(w, string(resp))
		}
	}
	w.WriteHeader(500)
	fmt.Fprintf(w, "{\"status\" : \"error\"}")
	panic(err)
	return

}

//LoadTests -
func (c *Core) LoadTests() error {
	var swap structures.Test
	swap, err := readTest("tests1.json")
	if err != nil {
		return err
	}
	c.Tests = append(c.Tests, swap)
	swap, err = readTest("tests2.json")
	if err != nil {
		return err
	}
	c.Tests = append(c.Tests, swap)
	swap, err = readTest("tests3.json")
	if err != nil {
		return err
	}
	c.Tests = append(c.Tests, swap)
	return nil
}

//readTest -
func readTest(name string) (structures.Test, error) {
	out := structures.Test{}
	temFile, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Println("Reading tests failed")
		return out, err
	}
	err = json.Unmarshal(temFile, &out)
	if err != nil {
		fmt.Println("Unmarshal config failed")
		return out, err
	}
	return out, nil
}
