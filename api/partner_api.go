package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/getfarook/crud-http-server/models"
	"github.com/getfarook/crud-http-server/utils"
	"github.com/gorilla/mux"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func GetPartner(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetPartner function is invoked")

	vars := mux.Vars(r)
	id := vars["id"]

	db := utils.DbHandle()
	defer db.Close()

	var partner models.Partner

	err := db.QueryRow("SELECT id, name, age, dob, balance, access FROM partners WHERE id = $1", id).Scan(&partner.Id, &partner.Name, &partner.Age, &partner.Dob, &partner.Balance, &partner.Access)
	fmt.Println("Before switch")
	fmt.Println(partner)
	switch {
	case err == sql.ErrNoRows:
		//log.Fatalf("no user with id %d", id)
		json.NewEncoder(w).Encode(fmt.Sprint("Partner with id:", id, " does not exist"))

	case err != nil:
		//log.Fatal(err)
		panic(err)
	default:
		//log.Printf("name is %s\n", partner.Name)
		json.NewEncoder(w).Encode(partner)
	}

}

func GetAllPartners(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetAllPartners function is invoked")

	db := utils.DbHandle()
	defer db.Close()

	rows, err := db.Query("SELECT id, name, age, dob, balance, access FROM partners ")
	if err != nil {
		// handle this error better than this
		panic(err)
	}
	defer rows.Close()

	retrievedPartners := []*models.Partner{}

	for rows.Next() {
		var partner models.Partner

		err1 := rows.Scan(&partner.Id, &partner.Name, &partner.Age, &partner.Dob, &partner.Balance, &partner.Access)

		if err1 != nil {
			panic(err1)
		}

		retrievedPartners = append(retrievedPartners, &partner)

	}

	json.NewEncoder(w).Encode(retrievedPartners)

}

func AddParner(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Add Partner function is invoked")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var partner models.Partner
	json.Unmarshal(reqBody, &partner)

	db := utils.DbHandle()
	defer db.Close()

	sqlStatement := "INSERT INTO partners (name, age, dob, balance, access) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	var id int32 = 0

	err := db.QueryRow(sqlStatement, partner.Name, partner.Age, partner.Dob, partner.Balance, partner.Access).Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Println("New Partner ID is: ", id)

	//fmt.Println("Parnter received is :")
	//fmt.Println(partner)
	//json.NewEncoder(w).Encode(partner)
	message := fmt.Sprint("New partner is created with id: ", id)
	json.NewEncoder(w).Encode(message)

	//fmt.Fprintf(w, "%+v", string(reqBody))
}

func DeleteParner(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DeletePartner function is invoked")
	vars := mux.Vars(r)
	id := vars["id"]

	db := utils.DbHandle()
	defer db.Close()

	sqlStatement := "DELETE FROM partners WHERE id = $1;"

	res, err1 := db.Exec(sqlStatement, id)
	if err1 != nil {
		panic(err1)
	}

	count, err2 := res.RowsAffected()
	if err2 != nil {
		panic(err2)
	}

	var message string
	if count > 0 {
		message = fmt.Sprint("Partner with id:", id, " is successfully deleted")
	} else {
		message = fmt.Sprint("Partner with id:", id, " does not exist")
	}

	json.NewEncoder(w).Encode(message)

}

func UpdateParner(w http.ResponseWriter, r *http.Request) {
	fmt.Println("UpdatePartner function is invoked")

	reqBody, _ := ioutil.ReadAll(r.Body)
	var partner models.Partner
	json.Unmarshal(reqBody, &partner)

	vars := mux.Vars(r)
	id := vars["id"]

	db := utils.DbHandle()
	defer db.Close()

	sqlStatement := "UPDATE partners SET name = $1, age = $2, dob = $3, balance = $4, access = $5 WHERE id = $6;"

	res, err1 := db.Exec(sqlStatement, partner.Name, partner.Age, partner.Dob, partner.Balance, partner.Access, id)
	if err1 != nil {
		panic(err1)
	}

	count, err2 := res.RowsAffected()
	if err2 != nil {
		panic(err2)
	}

	var message string
	if count > 0 {
		message = fmt.Sprint("Partner with id:", id, " is successfully updated")
	} else {
		message = fmt.Sprint("Partner with id:", id, " does not exist")
	}

	json.NewEncoder(w).Encode(message)

}
