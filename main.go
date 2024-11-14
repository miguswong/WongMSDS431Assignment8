package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"

	"github.com/seehuhn/mt19937"
)

// Function to generate random normal data
func generateNormalData(rng *rand.Rand, mean, stdDev float64, size int) []float64 {
	data := make([]float64, size)
	for i := 0; i < size; i++ {
		data[i] = rng.NormFloat64()*stdDev + mean
	}
	return data
}

// Returns the mean of a slice
func mean(data []float64) float64 {
	sum := 0.0
	for _, v := range data {
		sum += v
	}
	return sum / float64(len(data))
}

// Returns the median of a slice
func median(data []float64) float64 {
	n := len(data)
	sortedData := make([]float64, n)
	copy(sortedData, data)
	sort.Float64s(sortedData)
	if n%2 == 0 {
		return (sortedData[n/2-1] + sortedData[n/2]) / 2.0
	}
	return sortedData[n/2]
}

// returns a filtered slices based off first column value
func filterRowsByn(studyResult [][]float64, value float64) [][]float64 {
	var filteredResults [][]float64
	for _, row := range studyResult {
		if row[0] == value {
			filteredResults = append(filteredResults, row)
		}
	}
	return filteredResults
}

// Returns Standard Deviation given a 2-d for a given "column"
func StdDevByn(nstudyResult [][]float64, colNum int) float64 {
	var stdDev float64
	var meanData []float64

	//Grab all sample means (2nd column) and store in slice
	for _, row := range nstudyResult {
		meanData = append(meanData, row[colNum])
	}

	//Calculate mean of sample means
	meanValue := mean(meanData)
	sum := 0.0
	for _, v := range meanData {
		sum += (v - meanValue) * (v - meanValue)
	}
	variance := sum / float64(len(meanData)-1)
	stdDev = math.Sqrt(variance)

	return stdDev
}

func main() {
	start := time.Now()

	B := 100
	fmt.Printf("\nRunning study with %d bootstrap samples\n", B)

	rng := rand.New(mt19937.New())
	seed := int64(9999)
	rng.Seed(seed)

	//Populaiton parameters
	popMean := 100.0
	popSD := 10.0
	studySampleSizes := []int{25, 100, 225, 400}
	studyResult := [][]float64{}

	fmt.Printf("\nStudy conditions:\n  Population mean: %.2f SD: %.2f\n", popMean, popSD)
	for iteration := 1; iteration <= 100; iteration++ {
		for _, n := range studySampleSizes { // Code to be executed in each iteration
			thisSample := generateNormalData(rng, popMean, popSD, n)
			thisSampleMean := mean(thisSample)

			bootstrapMeans := make([]float64, B)
			bootstrapMedians := make([]float64, B)

			for b := 0; b < B; b++ {
				thisBootstrapSample := make([]float64, n) //Initilize slice to hold random samples
				for i := 0; i < n; i++ {
					thisBootstrapSample[i] = thisSample[rand.Intn(n)] //Randomly select a sample from the population
				}
				thisBootstrapMean := mean(thisBootstrapSample)
				thisBootstrapMedian := median(thisBootstrapSample)
				bootstrapMeans[b] = thisBootstrapMean
				bootstrapMedians[b] = thisBootstrapMedian
			}
			bootMean := mean(bootstrapMeans)
			bootMedian := mean(bootstrapMedians)

			thisIterationResults := []float64{float64(n), thisSampleMean, bootMean, bootMedian}
			studyResult = append(studyResult, thisIterationResults)
		}
	}

	fmt.Printf("\nEstimated standard errors using %v bootstrap samples\n", B)
	for _, n := range studySampleSizes {
		nStudyResult := filterRowsByn(studyResult, float64(n))

		fmt.Println("\nSamples of size n = ", n)
		fmt.Printf("  SE Mean from Central Limit Theorem: %.2f", popSD/math.Sqrt(float64(n)))
		fmt.Printf("\n  SE Mean from Samples: %.2f", StdDevByn(nStudyResult, 1))
		fmt.Printf("\n  SE Mean from Bootstrap Samples: %.2f", StdDevByn(nStudyResult, 2))
		fmt.Printf("\n  SE Median from Bootstrap Samples: %.2f\n", StdDevByn(nStudyResult, 3))
	}

	fmt.Println("\n----- Run Complete -----")
	elapsed := time.Since(start)
	fmt.Printf("Elapsed time: %.2f seconds\n", elapsed.Seconds())
}
