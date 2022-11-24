package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"sort"
	"strings"
)

func countMean(bunch []int) float64 {
	var sum int
	for _, elem := range bunch {
		sum += elem
	}
	var mean = float64(sum) / float64(len(bunch))
	return mean
}

func countMedian(bunch []int) float64 {
	var size = len(bunch)
	var median float64
	if size%2 == 0 {
		median = float64(bunch[size/2-1]+bunch[size/2]) / 2.0
	} else {
		median = float64(bunch[size/2])
	}
	return median
}

func countMode(bunch []int) int {
	bunchMap := make(map[int]int)
	var max int
	var mode int
	for i := range bunch {
		bunchMap[bunch[i]] += 1
		if bunchMap[bunch[i]] > max {
			max = bunchMap[bunch[i]]
			mode = bunch[i]
		}
	}
	return mode
}

func countSD(bunch []int) float64 {
	mean := countMean(bunch)
	var sum float64
	var sd float64
	for i := range bunch {
		sum += math.Pow(float64(bunch[i])-mean, 2)
	}
	sd = math.Sqrt(sum / float64(len(bunch)))
	return sd
}

func getMetrics() map[string]bool {
	var metrics = make(map[string]bool)
	if len(os.Args) == 1 {
		metrics["mean"] = true
		metrics["median"] = true
		metrics["mode"] = true
		metrics["sd"] = true
	} else if len(os.Args) > 1 && len(os.Args) < 5 {
		args := os.Args[1:]
		for i := range args {
			str := strings.ToLower(args[i])
			if str == "mean" || str == "median" || str == "mode" || str == "sd" {
				if metrics[str] == true {
					fmt.Fprintf(os.Stderr, "error: '%s' is repeated argument\n", str)
					os.Exit(1)
				}
				metrics[str] = true
			} else {
				fmt.Fprintf(os.Stderr, "error: '%s' is incorrect argument\n", str)
				os.Exit(1)
			}
		}
	} else {
		fmt.Fprintln(os.Stderr, "error: too much arguments")
		os.Exit(1)
	}
	return metrics
}

func getBunch() []int {
	var bunch []int
	fmt.Println("Enter numbers between -100000 and 100000 separated by newlines:")
	for {
		var number int
		_, err := fmt.Scanf("%d\n", &number)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Fprintf(os.Stderr, "error: %v\n", err)
				os.Exit(1)
			}
		}
		if number < -100000 || number > 100000 {
			fmt.Fprintf(os.Stderr, "error: the number %d is out of range\n", number)
			os.Exit(1)
		}
		bunch = append(bunch, number)
	}
	if len(bunch) == 0 {
		fmt.Fprintln(os.Stderr, "error: empty set")
		os.Exit(1)
	}
	sort.Ints(bunch)
	return bunch
}

func main() {
	metrics := getMetrics()
	bunch := getBunch()

	for i := range metrics {
		if metrics[i] == true {
			switch i {
			case "mean":
				fmt.Printf("Mean: %.1f\n", countMean(bunch))
			case "median":
				fmt.Printf("Median: %.1f\n", countMedian(bunch))
			case "mode":
				fmt.Printf("Mode: %d\n", countMode(bunch))
			case "sd":
				fmt.Printf("SD: %.1f\n", countSD(bunch))
			}
		}
	}
}
