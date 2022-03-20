package v1

import (
	"doctorx/pkg/db"
	"encoding/json"
	"log"
	"strconv"
	"io/ioutil"
	"github.com/gin-gonic/gin"
)

// define the struct, the field must be Capitalized, otherwise it
type Report struct {
    Reportid  int `json:"reportid"`
	Userid int `json:"userid"`
    Collection_date string `json:"collection_date"`
	Organization string `json:"organization"`
	Section string `json:"section"`
	Status string `json:"status"`
}

// define the struct, the field must be Capitalized, otherwise it
type NewReport struct {
    Userid  int
    Collection_date string
	Organization string
	Section string
	Status string

}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func GetReports(c *gin.Context) {
	//check header
	//userid := c.Request.URL.Query().Get("userid")

	//connect to the db
	conn := db.GetSQLite3Connection()
	
	//run the query
	statement := "select reportid, userid, collection_date, organization, type, status from reports"//where userid="+userid
	rs, err := db.RunQuery(conn, statement)
	CheckError(err)
	// scan row results
	var reports []Report 
	for rs.Next() {
		tmp := Report{0, 0, "", "", "", ""}
		err := rs.Scan(&tmp.Reportid, &tmp.Userid,&tmp.Collection_date, &tmp.Organization, &tmp.Section, &tmp.Status)
		CheckError(err)
		//print the struct
		log.Printf("%+v\n", tmp.Collection_date)
		//fmt.Println(report.Collection_date)
		reports = append(reports, tmp)
	}
	
	//close connection
	rs.Close()
	conn.Close()

	//reply the response
	c.JSON(200, reports)

	//additional
	//print slices 
	//log.Println(reports)
	//marshal to convert to []byte //unmarshal to covert []byte to json
	//r, _ := json.Marshal(reports)
	//log.Printf(string(r))
}

func InsertReport(c *gin.Context) {
	//connect to the db
	conn := db.GetSQLite3Connection()
	
	body, err := ioutil.ReadAll(c.Request.Body)
	log.Println("----body")
	log.Println(string(body))
	
	var report NewReport
	//err = json.NewDecoder(c.Request.Body).Decode(&report)
	//CheckError(err)
	err = json.Unmarshal(body,&report)
	CheckError(err)
	
	log.Print("###########################Post Report#############")
	//run the query
	log.Println(report.Collection_date)
	statement := "insert into reports(userid, collection_date, organization, type, status) values(" + strconv.Itoa(report.Userid) + ",'" + report.Collection_date + "','" + report.Organization + "','" + report.Section+ "','" + report.Status + "')"
	log.Print(statement)
	stmt,err := conn.Prepare(statement)
	CheckError(err)
	_,err = stmt.Exec()
	CheckError(err)
	stmt.Close()
	conn.Close()

	//reply the response
	c.JSON(200, gin.H{
		"message": "Succeed to insert a new record.",
	})
}