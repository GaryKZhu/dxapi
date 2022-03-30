package main

import (
	//"flag"
	//"github.com/gin-gonic/gin"
	"fmt"
	"sort"
	//"strconv"
	"doctorx/pkg/db"
	"github.com/navossoc/bayesian"
	"strings"
)

type Disease struct {
	Name string
	S1   string
	S2   string
	S3   string
	S4   string
	S5   string
	S6   string
	S7   string
	S8   string
	S9   string
	S10   string
	S11  string
	S12   string
	S13   string
	S14  string
	S15  string
	S16  string
	S17  string
}

func main() {
	//body, err := ioutil.ReadAll(c.Request.Body)
	//var input query
	//err = json.Unmarshal(body,&query)
	convert := make(map[string]int);
	convert["fever"] = 0;
	convert["high_fever"] = 0; 
	convert["headache"] = 1;
	convert["tiredness"] = 2;
	convert["fatigue"] = 2;
	convert["weakness"] = 2;
	convert["nausea"] = 3;
	convert["dizziness"] = 3;
	convert["stuffy nose"] = 4; 
	convert["runny_nose"] = 4;
	convert["nasal congestion"] = 4;
	convert["sore throat"] = 5; 
	convert["wheezing"] = 6
	convert["coughing"] = 7; 
	convert["cough"] = 7;
	convert["sneezing"] = 8;
	convert["continuous_sneezing"] = 8
	convert["sneeze"] = 8;
	convert["shortness of breath"] = 9;
	convert["chest tightness"] = 10;
	convert["chest pain"] = 11; 
	convert["chest_pain"] = 11;
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
	convert["shivering"] = 22;
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

	conn := db.GetData()
	statement := "select * from symptoms1"
	rs, err := db.RunQuery(conn, statement)
	checkErr(err)
	tot := 0
	correct := 0
	for rs.Next() {
		var d1 Disease;

		err = rs.Scan(&d1.Name, &d1.S1, &d1.S2,&d1.S3,&d1.S4,&d1.S5,&d1.S6,&d1.S7,&d1.S8,&d1.S9,&d1.S10,&d1.S11,&d1.S12,&d1.S13,&d1.S14,&d1.S15, &d1.S16, &d1.S17)

		statement := "select * from symptoms2 where d_name = '" + d1.Name + "'"; 
		tmp, err := db.RunQuery(conn, statement)
		checkErr(err)
		yes := 0
		for tmp.Next() {
			yes = 1
		}
		if(yes == 0) {
			continue
		}

		tv := make([]float64, 24)
	
		counter := 1;

		if _, ok := convert[strings.TrimSpace(d1.S1)]; ok {
			if(tv[convert[strings.TrimSpace(d1.S1)]] == 0) {
				counter += 1
			}
			tv[convert[strings.TrimSpace(d1.S1)]] = 1	
		}
		if _, ok := convert[strings.TrimSpace(d1.S2)]; ok {
			if(tv[convert[strings.TrimSpace(d1.S2)]] == 0) {
				counter += 1
			}
			tv[convert[strings.TrimSpace(d1.S2)]] = 1		
		}
		if _, ok := convert[strings.TrimSpace(d1.S3)]; ok {
			if(tv[convert[strings.TrimSpace(d1.S3)]] == 0) {
				counter += 1
			}
			tv[convert[strings.TrimSpace(d1.S3)]] = 1		
		}
		if _, ok := convert[strings.TrimSpace(d1.S4)]; ok {
			if(tv[convert[strings.TrimSpace(d1.S4)]] == 0) {
				counter += 1
			}
			tv[convert[strings.TrimSpace(d1.S4)]] = 1		
		}
		if _, ok := convert[strings.TrimSpace(d1.S5)]; ok {
			if(tv[convert[strings.TrimSpace(d1.S5)]] == 0) {
				counter += 1
			}
			tv[convert[strings.TrimSpace(d1.S5)]] = 1		
		}
		if _, ok := convert[strings.TrimSpace(d1.S6)]; ok {
			if(tv[convert[strings.TrimSpace(d1.S6)]] == 0) {
				counter += 1
			}
			tv[convert[strings.TrimSpace(d1.S6)]] = 1		
		}
		if _, ok := convert[strings.TrimSpace(d1.S7)]; ok {
			if(tv[convert[strings.TrimSpace(d1.S7)]] == 0) {
				counter += 1
			}
			tv[convert[strings.TrimSpace(d1.S7)]] = 1		
		}
		if _, ok := convert[strings.TrimSpace(d1.S8)]; ok {
			if(tv[convert[strings.TrimSpace(d1.S8)]] == 0) {
				counter += 1
			}
			tv[convert[strings.TrimSpace(d1.S8)]] = 1		
		}
		if _, ok := convert[strings.TrimSpace(d1.S9)]; ok {
			if(tv[convert[strings.TrimSpace(d1.S9)]] == 0) {
				counter += 1
			}
			tv[convert[strings.TrimSpace(d1.S9)]] = 1		
		}
		if _, ok := convert[strings.TrimSpace(d1.S10)]; ok {
			if(tv[convert[strings.TrimSpace(d1.S10)]] == 0) {
				counter += 1
			}
			tv[convert[strings.TrimSpace(d1.S10)]] = 1		
		}
		if _, ok := convert[strings.TrimSpace(d1.S11)]; ok {
			if(tv[convert[strings.TrimSpace(d1.S11)]] == 0) {
				counter += 1
			}
			tv[convert[strings.TrimSpace(d1.S11)]] = 1		
		}
		if _, ok := convert[strings.TrimSpace(d1.S12)]; ok {
			if(tv[convert[strings.TrimSpace(d1.S12)]] == 0) {
				counter += 1
			}
			tv[convert[strings.TrimSpace(d1.S12)]] = 1		
		}
		if _, ok := convert[strings.TrimSpace(d1.S13)]; ok {
			if(tv[convert[strings.TrimSpace(d1.S13)]] == 0) {
				counter += 1
			}
			tv[convert[strings.TrimSpace(d1.S13)]] = 1		
		}
		if _, ok := convert[strings.TrimSpace(d1.S14)]; ok {
			if(tv[convert[strings.TrimSpace(d1.S14)]] == 0) {
				counter += 1
			}
			tv[convert[strings.TrimSpace(d1.S14)]] = 1		
		}
		if _, ok := convert[strings.TrimSpace(d1.S15)]; ok {
			if(tv[convert[strings.TrimSpace(d1.S15)]] == 0) {
				counter += 1
			}
			tv[convert[strings.TrimSpace(d1.S15)]] = 1		
		}		
		if _, ok := convert[strings.TrimSpace(d1.S16)]; ok {
			if(tv[convert[strings.TrimSpace(d1.S16)]] == 0) {
				counter += 1
			}
			tv[convert[strings.TrimSpace(d1.S16)]] = 1		
		}
		if _, ok := convert[strings.TrimSpace(d1.S17)]; ok {
			if(tv[convert[strings.TrimSpace(d1.S17)]] == 0) {
				counter += 1
			}
			tv[convert[strings.TrimSpace(d1.S17)]] = 1		
		}

		if(counter < 6) {
			continue;
		}
		fmt.Println(tv)

		type kv struct {
			Key   string
			Value float64
		}
		var ss []kv

		_, scores, names := MGetMostLikeDisease(tv, unconvert, convert)
		for i, score := range scores {
			ss = append(ss, kv{names[i], score})
		}
		sort.Slice(ss, func(i, j int) bool {
			return ss[i].Value > ss[j].Value
		})
		tot += 1
		if(ss[0].Key == d1.Name || ss[1].Key == d1.Name || ss[2].Key == d1.Name) {
			correct += 1
		}

		ans := make([]string, 0)
		
		fmt.Print(d1.Name)
		fmt.Print(" ")
		fmt.Print(ss[1].Key)
		fmt.Print(" ")
		fmt.Print(tot)
		fmt.Print(" ")
		fmt.Println(correct)
		for _, obj := range ss {
		//	fmt.Println(obj.Key)
			ans = append(ans, obj.Key)
		}
	}

	fmt.Println("*******RESULTS*******")
	fmt.Print("Amount Correct: ") 
	fmt.Println(correct)
	fmt.Print("Total Amount: ")
	fmt.Println(tot)
	fmt.Println("*******RESULTS*******")
}


//new idea
//compare slices one by one using knn
// find out most optimal this way


//return the most like vector
func MGetMostLikeDisease(query []float64, unconvert map[int]string, convert map[string]int) (m string, probs []float64, diseases []string) {
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
			fmt.Println(unconvert[i])
			cursymptoms = append(cursymptoms, unconvert[i])
		}
	}
	scores, likely, _ := classifier.LogScores(cursymptoms)

	return names[likely], scores, names;
} 



func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
