package v1


import (
	"github.com/gin-gonic/gin"
	"doctorx/pkg/db"
	"strconv"
	"io/ioutil"
	"encoding/json"
	"fmt"
)
type Users struct {
    Name  string `json:"name"`
	Userid int `json:"userid"`
}

type NewUser struct {
    Userid  int `json:"userid"`
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
	Timestamp int `json:"timestamp"`
	Email string `json:"email"`
	Height int `json:"height"`
	Weight int `json:"weight"`
	Birthday string `json:birthday`
	Gender string `json:gender`
}
func GetUser(c *gin.Context) {
	
	conn := db.GetSQLite3Connection()
	statement := "select firstname, userid from users"
	rs, err := db.RunQuery(conn, statement)

	
	vals := make([]Users, 0)
	var tmp Users

	for rs.Next() {
		err = rs.Scan(&tmp.Name, &tmp.Userid)
		checkErr(err)
		vals = append(vals, tmp)
	}
	
	rs.Close()
	conn.Close()
	c.JSON(200, vals)
	return
}

func InsertUser(c * gin.Context) {
	conn := db.GetSQLite3Connection()
	
	body, err := ioutil.ReadAll(c.Request.Body)
	statement := "select userid from users"
	rs, err := db.RunQuery(conn, statement)

	maxval := 0
	tmp := 0
	arr := make([]int, 0)
	newuserid := 1
	for rs.Next() {
		err = rs.Scan(&tmp)
		arr = append(arr, tmp)
		maxval = Max(maxval, tmp)
	}

	for i := 1; i <= maxval+1; i++ {
		if(i > len(arr)) {
			newuserid = i;
			break;
		}
		if(arr[i-1] != i) {
			newuserid = i
			break
		}
	}
	
	var user NewUser
	
	err = json.Unmarshal(body,&user)
	user.Userid = newuserid
	fmt.Println(user)
	user.Timestamp /= 1000
	statement = "insert into users(userid, firstname, lastname, timestamp, email, height, weight, birthday, gender) values(" + strconv.Itoa(user.Userid) + ",'" + user.Firstname + "','" + user.Lastname + "','" + strconv.Itoa(user.Timestamp) + "','" + user.Email + "',"  + strconv.Itoa(user.Height) + "," + strconv.Itoa(user.Weight) + ",'" + user.Birthday + "','" + user.Gender + "')"
	stmt,err := conn.Prepare(statement)
	CheckError(err)
	_,err = stmt.Exec()

	statement = "insert into sdetails(userid, timestamp, type, value, category) values (" + strconv.Itoa(user.Userid) + "," + strconv.Itoa(user.Timestamp) + ",'" + "Weight" + "'," + strconv.Itoa(user.Weight) + "," + "'Other'" + ")"
	stmt,err = conn.Prepare(statement)
	CheckError(err)
	_,err = stmt.Exec()
	CheckError(err)
	stmt.Close()
	conn.Close()

	CheckError(err)
}

func Max(x, y int) int {
    if x < y {
        return y
    }
    return x
}