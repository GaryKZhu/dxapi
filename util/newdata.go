package main

import (
	//"flag"
	"fmt"
	//"sort"
	"strconv"
	"strings"
	"doctorx/pkg/db"
//	"github.com/gaspiman/cosine_similarity"
	//"reflect"
	"github.com/sjwhitworth/golearn/base"
//	"github.com/sjwhitworth/golearn/evaluation"
//	"github.com/sjwhitworth/golearn/knn"
)




func main() {
	//body, err := ioutil.ReadAll(c.Request.Body)
	//var input query
	//err = json.Unmarshal(body,&query)

	convert := make([]map[string]int, 5);
	hc := make(map[string]int);
	luc := make(map[string]int);
	cc := make(map[string]int);
	ac := make(map[string]int);
	lic := make(map[string]int);

	hc["fever"] = 0;
	hc["headache"] = 1;
	hc["tiredness"] = 2;
	hc["fatigue"] = 2;
	hc["weakness"] = 2;
	hc["nausea"] = 3;
	hc["dizziness"] = 3;
	hc["nausea"] = 3;
	hc["dizziness"] = 3;
	luc["stuffy nose"] = 0;
	luc["nasal congestion"] = 0;
	luc["sore throat"] = 1; 
	luc["coughing"] = 2; 
	luc["cough"] = 2;
	luc["sneezing"] = 3;
	luc["sneeze"] = 3;
	luc["shortness of breath"] = 4;
	cc["chest tightness"] = 0;
	cc["chest pain"] = 1; 
	cc["sharp chest pain"] = 1;
	cc["burning chest pain"] = 1;
	cc["back pain"] = 2;
	ac["vomiting"] = 0;
	ac["sharp abdominal pain"] = 1;
	ac["burning abdominal pain"] = 1;
	ac["stomach swelling"] = 2;
	ac["bloating"] = 2;
	ac["diarrhea"] = 3;
	ac["constipation"] = 4;
	ac["bloody stools"] = 4
	ac["blood in stool"] = 4;
	lic["arm pain"] = 0;
	lic["arm itch"] = 1;
	lic["leg pain"] = 2;
	lic["leg itch"] = 3;
	lic["chills"] = 4;
	lic["Skin dryness, peeling, scaliness, or roughness"] = 5;
	convert[0] = hc;
	convert[1] = luc;
	convert[2] = cc;
	convert[3] = ac;
	convert[4] = lic; 

	for i := 0; i < 5; i += 1 {
		FitData(i, convert[i]);
	}
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
func FitData(id int, convert map[string]int) () {
	//fetch vector result set from db
	conn := db.GetData()
	attrs := make([]base.Attribute, 4)
	switch id {
		case 0: 
			attrs[0] = base.NewFloatAttribute("Fever")
			attrs[1] = base.NewFloatAttribute("Headache")
			attrs[2] = base.NewFloatAttribute("Tiredness")
			attrs[3] = base.NewFloatAttribute("Nausea/Dizziness")
			attrs = append(attrs, base.NewCategoricalAttribute())
			attrs[4].SetName("Names")
		case 1: 
			attrs[0] = base.NewFloatAttribute("Stuffy Nose")
			attrs[1] = base.NewFloatAttribute("Sore Throat")
			attrs[2] = base.NewFloatAttribute("Wheezing")
			attrs[3] = base.NewFloatAttribute("Coughing")
			attrs = append(attrs, base.NewFloatAttribute("Sneezing"))
			attrs = append(attrs, base.NewFloatAttribute("Shallow Breathing/Shortness of Breadth"))
			attrs = append(attrs, base.NewCategoricalAttribute())
			attrs[6].SetName("Names")
		case 2: 
			attrs[0] = base.NewFloatAttribute("Chest Tightness")
			attrs[1] = base.NewFloatAttribute("Chest Pain")
			attrs[2] = base.NewFloatAttribute("Back Pain")
			attrs[3] = base.NewCategoricalAttribute()
			attrs[3].SetName("Names")
		case 3: 
			attrs[0] = base.NewFloatAttribute("Vomiting")
			attrs[1] = base.NewFloatAttribute("Abdominal Pain")
			attrs[2] = base.NewFloatAttribute("Swelling")
			attrs[3] = base.NewFloatAttribute("Constipation/Diarrhea")
			attrs = append(attrs, base.NewFloatAttribute("Blood Stools"))
			attrs = append(attrs, base.NewCategoricalAttribute())
			attrs[5].SetName("Names")
		case 4: 
			attrs[0] = base.NewFloatAttribute("Arm Pain")
			attrs[1] = base.NewFloatAttribute("Leg Pain")
			attrs[2] = base.NewFloatAttribute("Arm Itch")
			attrs[3] = base.NewFloatAttribute("Leg Itch")
			attrs = append(attrs, base.NewFloatAttribute("Chills"))
			attrs = append(attrs, base.NewFloatAttribute("Dry Skin"))
			attrs = append(attrs, base.NewCategoricalAttribute())
			attrs[6].SetName("Names")
	}
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
} 



func String2Float64(s []string) (f []float64) {
	for i := 0; i < len(s); i += 1 {
		f64, err := strconv.ParseFloat(s[i], 64)
		checkErr(err)
		f = append(f, float64(f64))
	}
	return f
}
