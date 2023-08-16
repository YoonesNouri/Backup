package main

import (
	"fmt"
	"math"
)

type NaiveBayes struct {
	classes      []string
	classCounts  map[string]int
	featureProbs map[string]map[string]map[string]float64
}

func NewNaiveBayes() *NaiveBayes {
	return &NaiveBayes{
		classes:      make([]string, 0),
		classCounts:  make(map[string]int),
		featureProbs: make(map[string]map[string]map[string]float64),
	}
}

func (nb *NaiveBayes) Train(features []map[string]string, labels []string) {
	// Count the occurrences of each class
	for _, class := range labels {
		nb.classCounts[class]++
	}

	// Count the occurrences of each feature given a class
	for i, featuresMap := range features {
		class := labels[i]

		if _, ok := nb.featureProbs[class]; !ok {
			nb.featureProbs[class] = make(map[string]map[string]float64)
		}

		for feature, value := range featuresMap {
			if _, ok := nb.featureProbs[class][feature]; !ok {
				nb.featureProbs[class][feature] = make(map[string]float64)
			}

			nb.featureProbs[class][feature][value]++
		}
	}

	// Calculate the probabilities of each feature given a class
	for class := range nb.featureProbs {
		for feature := range nb.featureProbs[class] {
			totalCount := float64(nb.classCounts[class])

			for value := range nb.featureProbs[class][feature] {
				count := nb.featureProbs[class][feature][value]
				prob := count / totalCount
				nb.featureProbs[class][feature][value] = prob
			}
		}
	}
}

func (nb *NaiveBayes) Predict(features map[string]string) string {
	maxProb := math.Inf(-1)
	var predictedClass string

	for class := range nb.classCounts {
		prob := float64(nb.classCounts[class])

		for feature, value := range features {
			if _, ok := nb.featureProbs[class][feature]; ok {
				if featureProb, ok := nb.featureProbs[class][feature][value]; ok {
					prob *= featureProb
				} else {
					// Handle missing feature value case
					prob = 0
					break
				}
			} else {
				// Handle missing feature case
				prob = 0
				break
			}
		}

		if prob > maxProb {
			maxProb = prob
			predictedClass = class
		}
	}

	return predictedClass
}

func main() {
	// Example usage
	nb := NewNaiveBayes()

	// Training data
	features := []map[string]string{
		{"Outlook": "Sunny", "Temperature": "Hot", "Humidity": "High", "Windy": "False"},
		{"Outlook": "Sunny", "Temperature": "Hot", "Humidity": "High", "Windy": "True"},
		{"Outlook": "Overcast", "Temperature": "Hot", "Humidity": "High", "Windy": "False"},
		{"Outlook": "Rain", "Temperature": "Mild", "Humidity": "High", "Windy": "False"},
		{"Outlook": "Rain", "Temperature": "Cool", "Humidity": "Normal", "Windy": "False"},
		{"Outlook": "Rain", "Temperature": "Cool", "Humidity": "Normal", "Windy": "True"},
		{"Outlook": "Overcast", "Temperature": "Cool", "Humidity": "Normal", "Windy": "True"},
		{"Outlook": "Sunny", "Temperature": "Mild", "Humidity": "High", "Windy": "False"},
		{"Outlook": "Sunny", "Temperature": "Cool", "Humidity": "Normal", "Windy": "False"},
		{"Outlook": "Rain", "Temperature": "Mild", "Humidity": "Normal", "Windy": "False"},
		{"Outlook": "Sunny", "Temperature": "Mild", "Humidity": "Normal", "Windy": "True"},
		{"Outlook": "Overcast", "Temperature": "Mild", "Humidity": "High", "Windy": "True"},
		{"Outlook": "Overcast", "Temperature": "Hot", "Humidity": "Normal", "Windy": "False"},
		{"Outlook": "Rain", "Temperature": "Mild", "Humidity": "High", "Windy": "True"},
	}

	labels := []string{
		"No", "No", "Yes", "Yes", "Yes", "No", "Yes", "No", "Yes", "Yes", "Yes", "Yes", "Yes", "No",
	}

	nb.Train(features, labels)

	// Test data
	testData := map[string]string{
		"Outlook":     "Sunny",
		"Temperature": "Cool",
		"Humidity":    "High",
		"Windy":       "False",
	}

	prediction := nb.Predict(testData)
	fmt.Println("Prediction:", prediction)
}
