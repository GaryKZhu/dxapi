package v1


import (
	"github.com/gin-gonic/gin"
	"doctorx/pkg/db"
	"encoding/json"
	"log"
	"io/ioutil"
	"strconv"
)

// define the struct, the field must be Capitalized, otherwise it
type Details struct {
    Reportid  int `json:"reportid"`
    Type string `json:"type"`
	Name string `json:"name"`
	Value string `json:"value"`
	Range string `json:"range"`
	Units string `json:"units"`
}

func GetDetails(c *gin.Context) {
	//check header
	reportid := c.Request.URL.Query().Get("reportid")

	//verify if sql statement is empty
	if reportid == "" {
		c.JSON(500, gin.H{
			"message": "missing report id.",
		})
		return
	}
	
	//connect to the db
	conn := db.GetSQLite3Connection()
	
	//run the query
	statement := "select reportid, type, name, value, range, units from details where reportid="+reportid
	rs, err := db.RunQuery(conn, statement)
	CheckError(err)
	// scan row results
	var details []Details 
	for rs.Next() {
		detail := new(Details) 
		err := rs.Scan(&detail.Reportid, &detail.Type, &detail.Name, &detail.Value, &detail.Range, &detail.Units)
		CheckError(err)
		//print the struct
		//log.Printf("%+v\n", detail)
		details = append(details, *detail)
	}
	
	//close connection
	rs.Close()
	conn.Close()

	//reply the response
	c.JSON(200, details)

	//additional
	//print slices 
	//log.Println(details)
	//marshal to convert to []byte //unmarshal to covert []byte to json
	//r, _ := json.Marshal(details)
	//log.Printf(string(r))
}

func InsertDetails(c *gin.Context) {
	//connect to the db
	conn := db.GetSQLite3Connection()
	
	var details []Details
	body, err := ioutil.ReadAll(c.Request.Body)
	CheckError(err)
	err = json.Unmarshal(body,&details)
	CheckError(err)
	
	log.Print("###########################Post Report Details#############")
	for _, detail := range details {
		statement := "insert into details(reportid, type, name, value, range, units) values (" + strconv.Itoa(detail.Reportid) + ",'" + detail.Type + "','" + detail.Name + "','" + detail.Value+ "','" + detail.Range + "','"+detail.Units+ "')"
		log.Print(statement)
		stmt,err := conn.Prepare(statement)
		CheckError(err)
		_,err = stmt.Exec()
		CheckError(err)
		stmt.Close()
 	}
	 conn.Close()
	//reply the response
	c.JSON(200, gin.H{
		"message": "Succeed to insert report details.",
	})
}