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
type Biochemistry struct {
	Userid int `json:userid`
	Collection_date string `json:collection_date`
	Timestamp int `json:timestamp`
    Creatinine float64 `json:creatinine`
    Cystatin_c float64 `json:cystatin_c`
    Egfr_crea float64 `json:egfr_crea`
    Egfr_cys float64 `json:egfr_cys`
    Sgpt_alt float64 `json:sgpt_alt`
	Glucose float64 `json:glucose`
    Sodium float64 `json:sodium`
    Potassium float64 `json:potassium`
    Albumin float64 `json:albumin`
    Bilirubin float64 `json:bilrubin`
    Alkaline float64 `json:alkaline`
    Gamma float64 `json:gamma`
	Alanine float64 `json:alanine`
    Lactate float64 `json:lactate`
    Lipase float64 `json:lipase`
    Cr_protein float64 `json:cr_protein`
}


func InsertBiochemistry(c *gin.Context) {
	//connect to the db
	conn := db.GetSQLite3Connection()

	var report Biochemistry
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
		u := ""
		r := ""

		switch i {
			case 3: 
				t = "Creatinine"
				u = "umol/L"
				r = "44 - 88"
			case 4: 
				t = "Cystatin C"
				u = "mg/L"
				r = "0.61 - 0.95"
			case 5: 
				t = "eGFR Creatinine"
				u = "mL/min/1.73 m*2"
				r = ">60"	
			case 6: 
				t = "eGFR Cystatin C"
				u = "mL/min/1.73 m*2"
				r = ">60"	
			case 7: 
				t = "SGPT (ALT)"
				u = "U/L"
				r = "<33"	
			case 8: 
				t = "Glucose Fasting"
				u = "mmol/L"
				r = "3.6 - 6.0"
			case 9: 
				t = "Sodium"
				u = "mmol/L"
				r = "135 - 145"
			case 10: 
				t = "Potassium"
				u = "mmol/L"
				r = "3.5 - 5.2"
			case 11: 
				t = "Albumin"
				u = "g/L"
				r = "35 - 52"
			case 12: 
				t = "Bilirubin Total"
				u = "umol/L"
				r = "<20"
			case 13: 
				t = "Alkaline Phosphate"
				u = "U/L"
				r = "35 - 120"
			case 14: 
				t = "Gamma Glutamyl Transferase"
				u = "U/L"
				r = "12 - 37"
			case 15: 
				t = "Alanine Aminotransferase"
				u = "U/L"
				r = "<36"
			case 16: 
				t = "Lactate Dehydrogenase"
				u = "U/L"
				r = "110 - 230"
			case 17: 
				t = "Lipase"
				u = "U/L"
				r = "<60"
			case 18: 
				t = "C Reactive Protein"
				u = "mg/L"
				r = "<5.0"
		}
		statement := "insert into details(reportid, userid, type, name, value, range, units, timestamp) values (" + strconv.Itoa(reportid) + "," + strconv.Itoa(report.Userid) + ",'Biochemistry','" + t + "'," + strconv.FormatFloat(tmp, 'E', -1, 64) + ",'" + r + "','" + u + "','" + strconv.Itoa(report.Timestamp) + "')"	
		fmt.Println(statement)
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
		"message": "Succeed to insert biochemistry.",
	})
}