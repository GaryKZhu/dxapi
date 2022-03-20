package v1

import (
	"fmt"
	"github.com/sjwhitworth/golearn/base"
//	"github.com/sjwhitworth/golearn/evaluation"
	"github.com/sjwhitworth/golearn/knn"
)

func main() {
	// Load in a dataset, with headers. Header attributes will be stored.
	// Think of instances as a Data Frame structure in R or Pandas.
	// You can also create instances from scratch.
	rawData, err := base.ParseCSVToInstances("routers/api/v1/test.csv", true)	
	if err != nil {
		panic(err)
	}
	// Print a pleasant summary of your data.
	//fmt.Println(rawData)

	//Initialises a new KNN classifier
	//euclidean and manhattan are consistent with linear and kdtree
	// cosine is not very consistent and also gives trash estimations
	cls := knn.NewKnnClassifier("euclidean", "kdtree", 3)

	//Do a training-test split
   // trainData, testData := base.InstancesTrainTestSplit(rawData, 0.50)
	cls.Fit(rawData)

	
	attrs := make([]base.Attribute, 5)
	attrs[0] = base.NewFloatAttribute("Fever")
	attrs[1] = base.NewFloatAttribute("Headache")
	attrs[2] = base.NewFloatAttribute("Tiredness")
	attrs[3] = base.NewFloatAttribute("Nausea/Dizziness")
	attrs[4] = base.NewCategoricalAttribute()
	attrs[4].SetName("Names")
	// Now let's create the final instances set
	newInst := base.NewDenseInstances()

	// Add the attributes
	newSpecs := make([]base.AttributeSpec, len(attrs))
	for i, a := range attrs {
		newSpecs[i] = newInst.AddAttribute(a)
	}

	// Allocate space
	newInst.Extend(1)

	// Write the data
	newInst.Set(newSpecs[0], 0, newSpecs[0].GetAttribute().GetSysValFromString("3.0"))
	newInst.Set(newSpecs[1], 0, newSpecs[1].GetAttribute().GetSysValFromString("3.0"))
	newInst.Set(newSpecs[2], 0, newSpecs[2].GetAttribute().GetSysValFromString("2.0"))
	newInst.Set(newSpecs[3], 0, newSpecs[3].GetAttribute().GetSysValFromString("0.0"))
	newInst.Set(newSpecs[4], 0, newSpecs[4].GetAttribute().GetSysValFromString("Flu"))
	newInst.Set(newSpecs[4], 0, newSpecs[4].GetAttribute().GetSysValFromString("Cold"))
	newInst.Set(newSpecs[4], 0, newSpecs[4].GetAttribute().GetSysValFromString("Pneumonia"))
	newInst.Set(newSpecs[4], 0, newSpecs[4].GetAttribute().GetSysValFromString("Stomach Flu"))
	newInst.Set(newSpecs[4], 0, newSpecs[4].GetAttribute().GetSysValFromString("Eczema"))
	newInst.Set(newSpecs[4], 0, newSpecs[4].GetAttribute().GetSysValFromString("Unknown"))
	newInst.AddClassAttribute(attrs[len(attrs)-1])

	fmt.Println(newInst)
	fmt.Println(rawData)
	predictions, err := cls.Predict(newInst)
	if err != nil {
		panic(err)
	}
	//fmt.Println(predictions)
	fmt.Println(base.GetClass(predictions, 0))
	//println(base.GetClass(predictions, 1))
	//fmt.Println(base.GetClass(predictions, 1))
	//fmt.Println(base.GetClass(predictions, 2))
//	println(base.GetClass(predictions, 3))
//	println(base.GetClass(predictions, 4))
//	if err != nil {
//		panic(err)
//	}
//
	// Prints precision/recall metrics
	//confusionMat, err := evaluation.GetConfusionMatrix(testData, predictions)
	//if err != nil {
	//	panic(fmt.Sprintf("Unable to get confusion matrix: %s", err.Error()))
	//}
	//fmt.Println(evaluation.GetSummary(confusionMat))
}