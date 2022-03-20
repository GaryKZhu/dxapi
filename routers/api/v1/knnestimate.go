package v1

import (
	"github.com/gin-gonic/gin"
	"doctorx/pkg/db"
	//"encoding/json"
	"strconv"
//	"io/ioutil"
	"strings"
//	"github.com/gaspiman/cosine_similarity"
	"github.com/sjwhitworth/golearn/base"
//	"github.com/sjwhitworth/golearn/evaluation"
	"github.com/sjwhitworth/golearn/knn"
	"fmt"
)

func KnnEstimate(c *gin.Context) {
	convert := make(map[string]int);

	convert["fever"] = 0;
	convert["headache"] = 1;
	convert["tiredness"] = 2;
	convert["fatigue"] = 2;
	convert["weakness"] = 2;
	convert["nausea"] = 3;
	convert["dizziness"] = 3;
	convert["stuffy nose"] = 4; 
	convert["nasal congestion"] = 4;
	convert["sore throat"] = 5; 
	convert["wheezing"] = 6
	convert["coughing"] = 7; 
	convert["cough"] = 7;
	convert["sneezing"] = 8;
	convert["sneeze"] = 8;
	convert["shortness of breath"] = 9;
	convert["chest tightness"] = 10;
	convert["chest pain"] = 11; 
	convert["sharp chest pain"] = 11;
	convert["burning chest pain"] = 11;
	convert["back pain"] = 12;
	convert["vomiting"] = 13;
	convert["sharp abdominal pain"] = 14;
	convert["burning abdominal pain"] = 14;
	convert["stomach swelling"] = 15;
	convert["bloating"] = 15;
	convert["diarrhea"] = 16;
	convert["constipation"] = 16;
	convert["bloody stools"] = 17
	convert["blood in stool"] = 17;
	convert["arm pain"] = 18;
	convert["arm itch"] = 19;
	convert["leg pain"] = 20;
	convert["leg itch"] = 21;
	convert["chills"] = 22;
	convert["Skin dryness, peeling, scaliness, or roughness"] = 23;

	userid := c.Request.URL.Query().Get("userid")
	start := c.Request.URL.Query().Get("start")
	end := c.Request.URL.Query().Get("end")

	conn := db.GetSQLite3Connection()
	statement := "select head,lung,chest,abdomen,limbs from symptoms where userid=" + userid + " and timestamp >= " + start + " and timestamp <= " + end
	fmt.Println(statement)
	rs, err := db.RunQuery(conn, statement)
	counter := 0
	tv := make([]float64, 24)
	tmph := ""
	tmplu := ""
	tmpc := ""
	tmpa := ""
	tmpli := ""	
	for rs.Next() {
		err = rs.Scan(&tmph, &tmplu, &tmpc, &tmpa, &tmpli)
		checkErr(err)
		ptr := 0
		arr := String2Float64(strings.Split(tmph, ","))
		for _, a := range arr {
			tv[ptr] += a
			ptr++
		}

		arr = String2Float64(strings.Split(tmplu, ","))
		for _, a := range arr {
			tv[ptr] += a
			ptr++
		}

		arr = String2Float64(strings.Split(tmpc, ","))
		for _, a := range arr {
			tv[ptr] += a
			ptr++
		}

		arr = String2Float64(strings.Split(tmpa, ","))
		for _, a := range arr {
			tv[ptr] += a
			ptr++
		}

		arr = String2Float64(strings.Split(tmpli, ","))
		for _, a := range arr {
			tv[ptr] += a
			ptr++
		}

		counter++;
	}
	
	c.JSON(200, KGetMostLikeDisease(tv, convert))
}


//new idea
//compare slices one by one using knn
// find out most optimal this way


//return the most like vector
func KGetMostLikeDisease(tv []float64, convert map[string]int) (m string) {
	//fetch vector result set from db
	conn := db.GetData()
	attrs := make([]base.Attribute, 25)
	attrs[0] = base.NewFloatAttribute("Fever")
	attrs[1] = base.NewFloatAttribute("Headache")
	attrs[2] = base.NewFloatAttribute("Tiredness")
	attrs[3] = base.NewFloatAttribute("Nausea/Dizziness")
	attrs[4] = base.NewFloatAttribute("Stuffy Nose")
	attrs[5] = base.NewFloatAttribute("Sore Throat")
	attrs[6] = base.NewFloatAttribute("Wheezing")
	attrs[7] = base.NewFloatAttribute("Coughing")
	attrs[8] = base.NewFloatAttribute("Sneezing")
	attrs[9] = base.NewFloatAttribute("Shallow Breathing/Shortness of Breadth")
	attrs[10] = base.NewFloatAttribute("Chest Tightness")
	attrs[11] = base.NewFloatAttribute("Chest Pain")
	attrs[12] = base.NewFloatAttribute("Back Pain")
	attrs[13] = base.NewFloatAttribute("Vomiting")
	attrs[14] = base.NewFloatAttribute("Abdominal Pain")
	attrs[15] = base.NewFloatAttribute("Swelling")
	attrs[16] = base.NewFloatAttribute("Constipation/Diarrhea")
	attrs[17] = base.NewFloatAttribute("Bloody Stools")
	attrs[18] = base.NewFloatAttribute("Arm Pain")
	attrs[19] = base.NewFloatAttribute("Leg Pain")
	attrs[20] = base.NewFloatAttribute("Arm Itch")
	attrs[21] = base.NewFloatAttribute("Leg Itch")
	attrs[22] = base.NewFloatAttribute("Chills")
	attrs[23] = base.NewFloatAttribute("Dry Skin")
	attrs[24] = base.NewCategoricalAttribute()
	attrs[24].SetName("Names")
	//run the query
	knownInst := base.NewDenseInstances()
	Specs := make([]base.AttributeSpec, len(attrs))
	for i, a := range attrs {
		Specs[i] = knownInst.AddAttribute(a)
	}
	knownInst.AddClassAttribute(attrs[len(attrs)-1])

	statement := "select distinct d_name from symptoms2";
	rs, err := db.RunQuery(conn, statement)
	checkErr(err)
	names := make([]string, 0)

	for rs.Next() {
		name := ""
		err = rs.Scan(&name)
		names = append(names, name)
		checkErr(err)
	}
	counter := 0
	cls := knn.NewKnnClassifier("euclidean", "kdtree", 6)
	for _,name := range names {
		knownInst.Extend(1)
		statement := "select s_name,s_value from symptoms2 where d_name=\""+name+"\"";
		rs, err := db.RunQuery(conn, statement)
		checkErr(err)
		tc := make([]int, len(attrs)-1)
		tv := make([]float64, len(attrs)-1)
		for rs.Next() {
			symptom := ""
			value := 0.00
			err := rs.Scan(&symptom, &value)
			checkErr(err);
			value /= 20;
			symptom = strings.ToLower(symptom)
			if _, ok := convert[symptom]; ok {
				tv[convert[symptom]] += value;
				tc[convert[symptom]]++;
			}
		}
		for i := 0; i < len(tc); i += 1 { 
			if tc[i] != 0 {
				knownInst.Set(Specs[i], counter, Specs[i].GetAttribute().GetSysValFromString(strconv.FormatFloat((tv[i]/float64(tc[i])), 'f', 3, 64))) 
			} else {
				knownInst.Set(Specs[i], counter, Specs[i].GetAttribute().GetSysValFromString(strconv.FormatFloat(0, 'f', 3, 64)))
			}
		}
		knownInst.Set(Specs[len(attrs)-1], counter, Specs[len(attrs)-1].GetAttribute().GetSysValFromString(name))
		counter++
	}

	cls.Fit(knownInst)

	checkInst := base.NewDenseInstances()

	// Add the attributes
	for i, a := range attrs {
		Specs[i] = checkInst.AddAttribute(a)
	}

	// Allocate spconverte
	checkInst.Extend(1)

	for i := 0; i<len(tv); i += 1 {
		checkInst.Set(Specs[i], 0, Specs[i].GetAttribute().GetSysValFromString(strconv.FormatFloat(tv[i], 'f', 2, 64)))
	}

	for i := 0; i<len(names); i += 1 {
		checkInst.Set(Specs[len(attrs)-1], 0, Specs[len(attrs)-1].GetAttribute().GetSysValFromString(names[i]))
	}
	checkInst.Set(Specs[len(attrs)-1], 0, Specs[len(attrs)-1].GetAttribute().GetSysValFromString("Unknown"))
	// Write the data
	checkInst.AddClassAttribute(attrs[len(attrs)-1])
	
	//fmt.Println(checkInst)
	//fmt.Println(knownInst)
//	checkInst.Set(checkSpecs[0], 0, checkSpecs[0].GetAttribute().GetSysValFromString("3.0"))


	predictions, err := cls.Predict(checkInst)
	if err != nil {
		panic(err)
	}
	//close connection
	rs.Close()
	conn.Close()
	//return the result

	return base.GetClass(predictions, 0)
} 

