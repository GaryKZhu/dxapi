package v1


import (
	"github.com/gin-gonic/gin"
	"doctorx/pkg/db"
	"encoding/json"
	"log"
	"io/ioutil"
	"strconv"
	"strings"
)

// define the struct, the field must be Capitalized, otherwise it
type Symptoms struct {
    Reportid  int `json:"reportid"`
	Userid int `json:"userid"`
    Collection_date string `json:"collection_date"`
	Timestamp int `json:"timestamp"`
	Head string `json:"head"`
	Lung string `json:"lung"`
	Chest string `json:"chest"`
	Abdomen string `json:"abdomen"`
	Limbs string `json:"limbs"`
	Other string `json:"other"`
	Comment string `json:"comment"`
}

func GetSymptoms(c *gin.Context) {

	//check header
	reportid := c.Request.URL.Query().Get("reportid")

	//verify if sql statement is empty
	if reportid == "" {
		c.JSON(500, gin.H{
			"message": "missing report id.",
		})
		return
	}

}

func InsertSymptom(c *gin.Context) {
	//connect to the db
	conn := db.GetSQLite3Connection()
	
	var symptom Symptoms
	body, err := ioutil.ReadAll(c.Request.Body)
	CheckError(err)
	err = json.Unmarshal(body,&symptom)
	CheckError(err)
	
	log.Print("###########################Post Symptom Details#############")
	statement := "insert into symptoms(userid, collection_date, head, lung, chest, abdomen, limbs, other, timestamp, comments) values (" + strconv.Itoa(symptom.Userid) + ",'" + symptom.Collection_date + "','" + symptom.Head + "','" + symptom.Lung+ "','" + symptom.Chest + "','"+symptom.Abdomen + "','"+symptom.Limbs+ "','" + symptom.Other + "','" + strconv.Itoa(symptom.Timestamp) + "','" + symptom.Comment + "')"
	log.Print(statement)
	stmt,err := conn.Prepare(statement)
	CheckError(err)
	_,err = stmt.Exec()
	CheckError(err)
	stmt.Close()
	symptom.Timestamp /= 1000; 

	s := strings.Split(symptom.Head, ",")

	for index, element := range s {
		var t string; 
		switch index {
			case 0: t = "Fever"
			case 1: t = "Headache"
			case 2: t = "Tiredness"
			case 3: t = "Dizziness"
		}
		statement := "insert into sdetails(userid, timestamp, type, value, category) values (" + strconv.Itoa(symptom.Userid) + "," + strconv.Itoa(symptom.Timestamp) + ",'" + t + "'," + element + "," + "'Head'" + ")"
		log.Print(statement)
		stmt,err := conn.Prepare(statement)
		CheckError(err)
		_,err = stmt.Exec()
		CheckError(err)
		stmt.Close()
	}

	
	s = strings.Split(symptom.Lung, ",")

	for index, element := range s {
		var t string; 
		switch index {
			case 0: t = "Stuffy Nose"
			case 1: t = "Sore Throat"
			case 2: t = "Wheezing"
			case 3: t = "Coughing"
		}
		statement := "insert into sdetails(userid, timestamp, type, value, category) values (" + strconv.Itoa(symptom.Userid) + "," + strconv.Itoa(symptom.Timestamp) + ",'" + t + "'," + element + "," + "'Lung'" + ")"
		stmt,err := conn.Prepare(statement)
		CheckError(err)
		_,err = stmt.Exec()
		CheckError(err)
		stmt.Close()
	}



	s = strings.Split(symptom.Chest, ",")

	for index, element := range s {
		var t string; 
		switch index {
			case 0: t = "Chest Tightness"
			case 1: t = "Chest Pain"
			case 2: t = "Back Pain"
		}
		statement := "insert into sdetails(userid, timestamp, type, value, category) values (" + strconv.Itoa(symptom.Userid) + "," + strconv.Itoa(symptom.Timestamp) + ",'" + t + "'," + element + "," + "'Chest'" + ")"		
		stmt,err := conn.Prepare(statement)
		CheckError(err)
		_,err = stmt.Exec()
		CheckError(err)
		stmt.Close()
	}

	s = strings.Split(symptom.Abdomen, ",")

	for index, element := range s {
		var t string; 
		switch index {
			case 0: t = "Vomiting"
			case 1: t = "Stomach Ache"
			case 2: t = "Swelling"
			case 3: t = "Constipation/Diarrhea"
			case 4: t = "Bloody Stools" 
		}
		statement := "insert into sdetails(userid, timestamp, type, value, category) values (" + strconv.Itoa(symptom.Userid) + "," + strconv.Itoa(symptom.Timestamp) + ",'" + t + "'," + element + "," + "'Abdomen'" + ")"	
		stmt,err := conn.Prepare(statement)
		CheckError(err)
		_,err = stmt.Exec()
		CheckError(err)
		stmt.Close()
	}


	s = strings.Split(symptom.Limbs, ",")

	for index, element := range s {
		var t string; 
		switch index {
			case 0: t = "Arm Pain"
			case 1: t = "Leg Pain"
			case 2: t = "Arm Itch"
			case 3: t = "Leg Itch"
			case 4: t = "Chills"
			case 5: t = "Dry Skin"
		}
		statement := "insert into sdetails(userid, timestamp, type, value, category) values (" + strconv.Itoa(symptom.Userid) + "," + strconv.Itoa(symptom.Timestamp) + ",'" + t + "'," + element + "," + "'Limbs'" + ")"		
		stmt,err := conn.Prepare(statement)
		CheckError(err)
		_,err = stmt.Exec()
		CheckError(err)
		stmt.Close()
	}

	s = strings.Split(symptom.Other, ",")
	for index, element := range s {
		var t string; 
		switch index {
			case 0: t = "Systolic Blood Pressure"
			case 1: t = "Diabolic Blood Pressure"
			case 2: t = "Blood Sugar"
			case 3: t = "Heart Rate"
			case 4: t = "Weight"
		}
		statement := "insert into sdetails(userid, timestamp, type, value, category) values (" + strconv.Itoa(symptom.Userid) + "," + strconv.Itoa(symptom.Timestamp) + ",'" + t + "'," + element + "," + "'Other'" + ")"		
		stmt,err := conn.Prepare(statement)
		CheckError(err)
		_,err = stmt.Exec()
		CheckError(err)
		stmt.Close()
	}

	conn.Close()
	//reply the response
	c.JSON(200, gin.H{
		"message": "Succeed to insert symptom.",
	})
}