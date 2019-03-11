package core

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/michalnov/psychology/bin/core/structures"

	//need
	_ "github.com/go-sql-driver/mysql" //needed
)

//Core -
type Core struct {
	DbMaster string
	DbSlave  string
	Tests    []structures.Test
}

type user struct {
	Kluc  string `json:"kluc,ommitempty"`
	Vek   string `json:"vek,ommitempty"`
	Rod   string `json:"rod,ommitempty"`
	Skola string `json:"skola,ommitempty"`
}

//Ping --
func (c *Core) Ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	fmt.Fprintf(w, "{\"status\" : \"ok\"}")
}

//UserHandler --
func (c *Core) UserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(200)
		fmt.Fprintf(w, "{\"status\" : \"error\"}")
	}
	var req user
	err := json.NewDecoder(r.Body).Decode(&req)
	fmt.Println("\n\n", req.Kluc)
	if err != nil {
		w.WriteHeader(300)
		fmt.Fprintf(w, "{\"status\" : \"error\"}")
		return
	}
	db, err := sql.Open("mysql", c.DbMaster)
	if err != nil {
		w.WriteHeader(300)
		fmt.Fprintf(w, "{\"status\" : \"error\"}")
		return
	}
	defer db.Close()
	statement, err := db.Prepare("insert into user(age,gender,school,testkey) values(?,?,?,?)")
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "{\"status\" : \"error\"}")
		return
	}
	res, err := statement.Exec(req.Vek, req.Rod, req.Skola, req.Kluc)
	id6, err := res.LastInsertId()
	id := strconv.FormatInt(id6, 10)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "{\"status\" : \"error\"}")
		return
	}
	w.WriteHeader(200)
	fmt.Fprintf(w, "{\"user\" : \""+id+"\"}")
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
		w.WriteHeader(400)
		fmt.Fprintf(w, "{\"status\" : \"json error\"}")
		return
	}
	for _, test := range c.Tests {
		if test.Key == req.Test {
			resp, err := json.MarshalIndent(test, "  ", "    ")
			if err != nil {
				w.WriteHeader(500)
				fmt.Fprintf(w, "{\"status\" : \"marshal error\"}")
				return
			}
			db, err := sql.Open("mysql", c.DbMaster)
			if err != nil {
				w.WriteHeader(500)
				fmt.Fprintf(w, "{\"status\" : \"mysql error\"}")
				return
			}
			defer db.Close()
			statement, err := db.Prepare("update user set testid = ? where iduser = ?")
			if err != nil {
				w.WriteHeader(500)
				fmt.Fprintf(w, "{\"status\" : \"sql error\"}")
				return
			}
			_, err = statement.Exec(test.TestID, req.UserID)
			if err != nil {
				w.WriteHeader(500)
				fmt.Fprintf(w, "{\"status\" : \"sql 2 error\"}")
				return
			}
			w.WriteHeader(200)
			fmt.Fprintf(w, string(resp))
			return
		}
	}
	w.WriteHeader(500)
	fmt.Fprintf(w, "{\"status\" : \"test not found error\"}")
	return

}

type naswer struct {
	Userid     int `json:"userid"`
	Testid     int `json:"testid"`
	Questionid int `json:"questionid"`
	Ansid      int `json:"answerid"`
}

//Answerque --
func (c *Core) Answerque(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(200)
		fmt.Fprintf(w, "{\"status\" : \"error\"}")
	}
	var req naswer
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(300)
		fmt.Fprintf(w, "{\"status\" : \"error\"}")
		return
	}
	db, err := sql.Open("mysql", c.DbMaster)
	if err != nil {
		w.WriteHeader(300)
		fmt.Fprintf(w, "{\"status\" : \"error\"}")
		return
	}
	defer db.Close()
	statement, err := db.Prepare("insert into answered(userid,testid,questionid,answerid) values(?,?,?,?)")
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "{\"status\" : \"error\"}")
		return
	}
	_, err = statement.Exec(req.Userid, req.Testid, req.Questionid, req.Ansid)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "{\"status\" : \"error\"}")
		return
	}
	w.WriteHeader(200)
	fmt.Fprintf(w, "{\"status\" : \"ok\"}")
	return
}

type gas struct {
	Userid int `json:"userid"`
	Gas1   int `json:"gas1"`
	Gas2   int `json:"gas2"`
	Gas3   int `json:"gas3"`
	Gas4   int `json:"gas4"`
	Gas5   int `json:"gas5"`
	Gas6   int `json:"gas6"`
	Gas7   int `json:"gas7"`
	Gas8   int `json:"gas8"`
	Gas9   int `json:"gas9"`
	Gas10  int `json:"gas10"`
}

//GetGas --
func (c *Core) GetGas(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(200)
		fmt.Fprintf(w, "{\"status\" : \"error\"}")
	}
	var req gas
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(300)
		fmt.Fprintf(w, "{\"status\" : \"error\"}")
		return
	}
	db, err := sql.Open("mysql", c.DbMaster)
	if err != nil {
		w.WriteHeader(300)
		fmt.Fprintf(w, "{\"status\" : \"error\"}")
		return
	}
	defer db.Close()
	statement, err := db.Prepare("insert into gas(userid,gas1,gas2,gas3,gas4,gas5,gas6,gas7,gas8,gas9,gas10) values(?,?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "{\"status\" : \"error\"}")
		return
	}
	_, err = statement.Exec(req.Userid, req.Gas1, req.Gas2, req.Gas3, req.Gas4, req.Gas5, req.Gas6, req.Gas7, req.Gas8, req.Gas9, req.Gas10)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "{\"status\" : \"error\"}")
		return
	}
	w.WriteHeader(200)
	fmt.Fprintf(w, "{\"status\" : \"ok\"}")
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

/*
func Activationmail(email string, tokenn string, mailer str.Mail) {
	m := gomail.NewMessage()
	m.SetHeader("From", "martinhercka1@gmail.com")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Hello")
	m.SetBody("text/html", mailer.Host+mailer.Port+"/auth/activate?token="+tokenn)

	if mailer.Username == "" || mailer.Password == "" {
		fmt.Println(mailer.Host + mailer.Port + "/auth/activate?token=" + tokenn)
		return
	}
	d := gomail.NewDialer("smtp.gmail.com", 587, mailer.Username, mailer.Password)

	if err := d.DialAndSend(m); err != nil {
		panic(err)
}
*/
