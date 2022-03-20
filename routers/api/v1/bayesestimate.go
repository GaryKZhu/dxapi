package v1

import (
	//"flag"
	"github.com/gin-gonic/gin"
	"fmt"
	"strings"
	"strconv"
	"doctorx/pkg/db"
	"github.com/navossoc/bayesian"
)



func BayesEstimate(c *gin.Context) {
	//body, err := ioutil.ReadAll(c.Request.Body)
	//var input query
	//err = json.Unmarshal(body,&query)
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

	unconvert := make(map[int]string);
	unconvert[0] = "fever";
	unconvert[1] = "headache";
	unconvert[2] = "tiredness";
	unconvert[3] = "nausea";
	unconvert[4] = "stuffy nose";
	unconvert[5] = "sore throat";
	unconvert[6] = "wheezing";
	unconvert[7] = "coughing";
	unconvert[8] = "sneezing";
	unconvert[9] = "shortness of breath";
	unconvert[10] = "chest tightness";
	unconvert[11] = "chest pain";
	unconvert[12] = "back pain";
	unconvert[13] = "vomiting";
	unconvert[14] = "abdominal pain";
	unconvert[15] = "bloating";
	unconvert[16] = "diarrhea";
	unconvert[17] = "bloody stools";
	unconvert[18] = "arm pain";
	unconvert[19] = "arm itch";
	unconvert[20] = "leg pain";
	unconvert[21] = "leg itch";
	unconvert[22] = "chills";
	unconvert[23] = "dry skin";

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
	
	c.JSON(200, BGetMostLikeDisease(tv, unconvert, convert))
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
func BGetMostLikeDisease(query []float64, unconvert map[int]string, convert map[string]int) (m string) {
	//fetch vector result set from db
	conn := db.GetData()

	statement := "select distinct d_name from symptoms2";
	rs, err := db.RunQuery(conn, statement)
	checkErr(err)
	names := make([]string, 0)
	classes := make([]bayesian.Class, 0)

	for rs.Next() {
		name := ""
		err = rs.Scan(&name)
		names = append(names, name)
		classes = append(classes, bayesian.Class(name))
		checkErr(err)
	}

	classifier := bayesian.NewClassifier(classes...)

	for _,name := range names {
		statement := "select s_name,s_value from symptoms2 where d_name=\""+name+"\"";
		rs, err := db.RunQuery(conn, statement)
		checkErr(err)
		tc := make([]int, len(query))
		tv := make([]float64, len(query))
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

		symptoms := make([]string, 0)
		probs := make([]float64, 0)
		for i := 0; i < len(tc); i += 1 { 
			if tc[i] != 0 {
				tv[i] /= float64(tc[i])
				symptoms = append(symptoms, unconvert[i])
				probs = append(probs, tv[i])
			}
		}
		classifier.LearnProbs(symptoms, probs, bayesian.Class(name))
	}

	
	cursymptoms := make([]string, 0)
	for i := 0; i<len(query); i += 1 {
		if query[i] > 0 {
			cursymptoms = append(cursymptoms, unconvert[i])
		}
	}
	_, likely, _ := classifier.LogScores(cursymptoms)

	return names[likely];
} 



func String2Float64(s []string) (f []float64) {
	for i := 0; i < len(s); i += 1 {
		f64, err := strconv.ParseFloat(s[i], 64)
		checkErr(err)
		f = append(f, float64(f64))
	}
	return f
}

