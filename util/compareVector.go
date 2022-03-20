package main

import (
	//"flag"
	"fmt"
	"strconv"
	"strings"
	"doctorx/pkg/db"
	"github.com/gaspiman/cosine_similarity"
	"reflect"
)

// define the struct, the field must be Capitalized, otherwise it
type Input struct {
    Start  int `json:"start"`
    End string `json:"end"`
	UserID int `json:"userid"`
}

type Vector struct {
	head     string
	lungs      string
	chest  string
	abdomen       string
	limbs string
	name string
	similarity float64
}

/*func CompareVector() {

	userid := c.Request.URL.Query().Get("userid")

	//verify if sql statement is empty
	if userid == "" {
		c.JSON(500, gin.H{
			"message": "missing user id.",
		})
		return
	}
	
	//connect to the db
	conn := db.GetSQLite3Connection()
	
	//run the query
	var input Input
	body, err := ioutil.ReadAll(c.Request.Body)
	CheckError(err)
	err = json.Unmarshal(body,&input)

	//get all data from db in timerange
	statement := "select head,lungs,chest,limbs,abdomen from details where userid="+input.userid + " and timestamp >= " + input.start + " and timestamp <= " + input.end
	rs, err := db.RunQuery(conn, statement)
	CheckError(err)
	
	//next we should take the average of the vector over the time range
	// ignore this step for now
	/*for rs.Next() {
		detail := new(Details) 
		err := rs.Scan(&detail.Reportid, &detail.Type, &detail.Name, &detail.Value, &detail.Range, &detail.Units)
		CheckError(err)
		//print the struct
		//log.Printf("%+v\n", detail)
		details = append(details, *detail)
	} */ 

/*
	//close connection
	rs.Close()
	conn.Close()


	// then pass it over into the function

	
	//reply the response
	//c.JSON(200, details)


	//test vector
	testvector := flag.String("3,3,2,0","3,2,0,4,2,0","0,0,0","0,0,0,0,0,0","1,1,0,0,0,2")
	flag.Parse()

	//get most like vector
	m := GetMostlikeVector(String2Float64(strings.Split(*testvector, ",")))

	//print out the result
	fmt.Println("*********** RCA Result ************")
	if m.similarity > 0 {
		fmt.Println("Disease: " + m.name)
		fmt.Printf("Similarity: %0.2f %% \n", m.similarity*100)
	} else {
		fmt.Printf("Does not find any closed feature vector. Please consider to build a new one.")
	}
	fmt.Println("***********************************")
} */


func main() {
	testhead := "3,3,2,0"
	testlungs := "3,2,0,4,2,0"
	testchest := "0,0,0"
	testabdomen := "0,0,0,0,0"
	testlimbs := "1,1,0,0,2,0"
	
	
	testvector := make([][]float64, 5)
	for i := 0; i < len(testvector); i++ {
		testvector[i] = make([]float64, 6)
	}

	testvector[0] = String2Float64(strings.Split(testhead,","))
	testvector[1] = String2Float64(strings.Split(testlungs,","))
	testvector[2] = String2Float64(strings.Split(testchest,","))
	testvector[3] = String2Float64(strings.Split(testabdomen,","))
	testvector[4] = String2Float64(strings.Split(testlimbs,","))
	tot := len(testvector[0]) + len(testvector[1]) + len(testvector[2]) + len(testvector[3]) + len(testvector[4])

	m := GetMostlikeVector(testvector, tot)
	//print out the result
	fmt.Println("*********** Result ************")
	if m.similarity > 0 {
		fmt.Println("Disease: " + m.name)
		fmt.Printf("Similarity: %0.2f %% \n", m.similarity*100)
	} else {
		fmt.Printf("No disease matched.")
	}
	fmt.Println("***********************************")
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}


//new idea
//compare slices one by one using knn
// find out most optimal this way


//return the most like vector
func GetMostlikeVector(tv [][]float64, tot int) (m *Vector) {
	//fetch vector result set from db
	conn := db.GetSQLite3Connection()
	
	//run the query
	statement := "select * from diseases"
	rs, err := db.RunQuery(conn, statement)
	checkErr(err)

	similarity := 0.0
	tmp := 0.0
	//iterate the result set
	var v Vector
	m = new(Vector)

	all := make([]float64, 5)

	for rs.Next() {
		tmp = 0.0

		//v = new(Vector)
		//get one record
		err = rs.Scan(&v.head, &v.lungs, &v.chest, &v.abdomen, &v.limbs, &v.name)
		checkErr(err)

		s := 0.0 
		println(v.name)
		//println(v.vector + v.cause + v.link)
		for i := 0; i < len(tv); i += 1 {
			tcase := make([]float64, len(tv[i]))
			copy(tcase, tv[i])
			compare := ""
			switch i {
				case 0: compare = v.head
				case 1: compare = v.lungs
				case 2: compare = v.chest
				case 3: compare = v.abdomen
				case 4: compare = v.limbs
			}
			
			f := make([]float64, len(tcase))
			f = String2Float64(strings.Split(compare, ","))

			for j := 0; j < len(tcase); j += 1 {
				if(f[j] == 0.0) {
					f[j] += 1.0
				}
				if(tcase[j] == 0.0) {
					tcase[j] += 1.0
				}
				all = append(all, tcase[j])
			}

			if(reflect.DeepEqual(tcase, f)) {
				s = 1.0
			} else {
				s, err = cosine_similarity.Cosine(tcase, f)
				checkErr(err)
			}
			tmp += s * float64(len(tcase)) / float64(tot)
		}
		
		fmt.Printf("Similarity: %0.2f %% \n", tmp*100)
		if tmp >= similarity {
			m.similarity = tmp
			m.head = v.head
			m.lungs = v.lungs
			m.chest = v.chest
			m.abdomen = v.abdomen
			m.limbs = v.limbs
			m.name = v.name
			similarity = tmp
		} 
	}

	//close connection
	rs.Close()
	conn.Close()
	//return the result
	return m
} 



func String2Float64(s []string) (f []float64) {
	for i := 0; i < len(s); i += 1 {
		f64, err := strconv.ParseFloat(s[i], 64)
		checkErr(err)
		f = append(f, float64(f64))
	}
	return f
}
