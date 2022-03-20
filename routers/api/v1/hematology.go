package v1


import (
	"github.com/gin-gonic/gin"
	"doctorx/pkg/db"
	"encoding/json"
	"io/ioutil"
	"fmt"
	"reflect"
	"strconv"
)

// define the struct, the field must be Capitalized, otherwise it
type Hematology struct {
	Userid int `json:userid`
	Collection_date string `json:collection_date`
	Timestamp int `json:timestamp`
	Sedimentation_rate float64 `json:sedimentation_rate` 
    Hemoglobin float64 `json:hemoglobin` 
    Hematocrit float64 `json:hematocrit` 
    Wbcc float64 `json:wbcc` 
    Rbcc float64 `json:rbcc` 
    Mcv float64 `json:mcv` 
    Mch float64 `json:mch` 
    Mchc float64 `json:mchc` 
    Rdw float64 `json:rdw` 
    Platelets float64 `json:platelets` 
    Neutrophils float64 `json:neutrophils` 
    Lymphocytes float64 `json:lymphocytes` 
    Monocytes float64 `json:monocytes` 
    Eosinophils float64 `j4son:eosinophils` 
    Basophils float64 `json:basophils` 
    Granulocytes float64 `json:granulcytes` 
    Nucleatedrbc float64 `json:nucleatedrbc` 
}

func InsertHematology(c *gin.Context) {
	//connect to the db
	conn := db.GetSQLite3Connection()
	
	var report Hematology
	body, err := ioutil.ReadAll(c.Request.Body)
	CheckError(err)
	err = json.Unmarshal(body,&report)
	CheckError(err)

	reportid := 0
	statement := "select reportid from reports"
	rs, err := db.RunQuery(conn, statement)

	for rs.Next() {
		tmp := 0
		err = rs.Scan(&tmp)
		reportid = Max(reportid, tmp)
	}

	reportid += 1
	fmt.Println(reportid)
	rs.Close()

	parsed := reflect.ValueOf(report);
	fmt.Println(parsed)

    for i := 3; i < parsed.NumField(); i++ {
		fmt.Println(parsed.Field(i).Interface())
        tmp := parsed.Field(i).Interface().(float64)
		if(tmp == -1) {
			continue
		}
		t := ""
		r := ""
		u := ""
		switch i {
			case 3: 
				t = "sedimentation_rate" 
				u = "mm/hr"
				r = "0-30"
			
			case 4: 
				t = "hemoglobin"
				u = "g/L"
				r = "115 - 155"
			
			case 5: 
				t = "hematocrit"
				u = "L/L"
				r = "0.33 - 0.46"
			case 6: 
				t = "wbcc"
				u = "x10E9/L"
				r = "4.0 - 11.0"
			case 7: 
				t = "rbcc"
				u = "x10E12/L"
				r = "3.60 - 5.20"
			case 8: 
				t = "mcv"
				u = "fl"
				r = "80 - 100"
			case 9: 
				t = "mch"
				u = "pg"
				r = "27 - 33"
			case 10: 
				t = "mchc"
				u = "g/L"
				r = "320 - 360"
			case 11: 
				t = "rdw"
				u = "%CV"
				r = "11.5 - 14.5"
			case 12: 
				t = "platelets"
				u = "x10E9/L"
				r = "150 - 400"
			case 13: 
				t = "neutrophils"
				u = "x10E9/L"
				r = "2.0 - 7.5"
			case 14: 
				t = "lymphocytes"
				u = "x10E9/L"
				r = "1.0 - 4.0"
			case 15: 
				t = "monocytes"
				u = "x10E9/L"
				r = "0.0 - 1.2"
			case 16: 
				t = "eosinophils"
				u = "x10E9/L"
				r = "0.0 - 0.7"
			case 17: 
				t = "basophils"
				u = "x10E9/L"
				r = "0.0 - 0.4"
			
			case 18: 
				t = "granulocytes"
				u = "x10E9/L"
				r = "0.0 - 0.1"
		
			case 19: 
				t = "nucleatedrbc"
				u = "/100 WBC"
				r = "0"
		}
		statement := "insert into details(reportid, userid, type, name, value, range, units, timestamp) values (" + strconv.Itoa(reportid) + "," + strconv.Itoa(report.Userid) + ",'Hematology','" + t + "'," + strconv.FormatFloat(tmp, 'E', -1, 64) + ",'" + r + "','" + u + "','" + strconv.Itoa(report.Timestamp) + "')"
		stmt,err := conn.Prepare(statement)
		CheckError(err)
		_,err = stmt.Exec()
		CheckError(err)
		stmt.Close()

    }


	conn.Close()
	fmt.Println(report)
	//reply the response
	c.JSON(200, gin.H{
		"message": "Succeed to insert hematology.",
	})
}