package qzschool

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"testing"
)

type qz struct {
	LoginCode string `json:"loginCode"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	ProjectId string `json:"projectId"`
	Version   string `json:"version"`
}

func getQz() *QzSchool {
	u := &qz{}
	file, err := os.Open("config.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	bs, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(bs, u)
	if err != nil {
		panic(err)
	}
	if u.LoginCode == "" || u.ProjectId == "" || u.Version == "" {
		return NewQzSchool(u.Username, u.Password)
	} else {
		return NewDetailQzSchool(u.Username, u.Password, u.LoginCode, u.ProjectId, u.Version)
	}
}

func TestQzSchool_Login(t *testing.T) {
	q := getQz()
	err := q.Login(DefaultVersion)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(q)
}

func TestQzSchool_GetAccountMoney(t *testing.T) {
	q := getQz()
	m, err := q.GetAccountMoney()
	if err != nil {
		panic(err)
	}
	fmt.Println(m)
}

func TestQzSchool_GetPageMonthBill(t *testing.T) {
	q := getQz()
	bill, err := q.GetPageMonthBill(BillAllType, "2022-04", 1)
	if err != nil {
		panic(err)
	}
	for i, it := range bill {
		fmt.Println(i+1, it)
	}
}
func TestQzSchool_GetAllMonthBill(t *testing.T) {
	q := getQz()
	bill, err := q.GetAllMonthBill(BillAllType, "2022-04")
	if err != nil {
		panic(err)
	}
	for i, it := range bill {
		fmt.Println(i+1, it)
	}
}
