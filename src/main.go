package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func main() {
	fmt.Println("Input bunch of inyeger numbers, strictly between -100000 and 100000.")

	scanner := bufio.NewScanner(os.Stdin)
	var arr []int
	var sum, i, mean, median, sd float64
	var Mean, Median, Mode, SD bool

	flag.BoolVar(&Mean, "mean", true, "display Mean")
	flag.BoolVar(&Median, "median", true, "display Median")
	flag.BoolVar(&Mode, "mode", true, "display Mode")
	flag.BoolVar(&SD, "sd", true, "display SD")
	flag.Parse()

	for scanner.Scan() {
		strIn := scanner.Text()
		numIn, err := strconv.Atoi(strIn)
		if err == nil {
			if numIn > 100000 || numIn < -100000 {
				fmt.Printf("Out of bounds value\n")
				continue
			} else {
				arr = append(arr, numIn)
				sum += float64(numIn)
				i++
			}
		} else {
			fmt.Printf("Non-int input\n")
			continue
		}
	}
	sort.Ints(arr)
	mean = sum / i

	if Mean && Median && Mode && SD {
		median = medianCalc(arr, int(i))
		mode := modeCalc(arr)
		sd = sdCalc(arr, mean, i)
		fmt.Printf("Mean: %.2f\nMedian: %.2f\nMode: %d\nSD: %.2f\n", mean, median, mode, sd)
	} else {
		if Mean {
			fmt.Printf("Mean: %.2f\n", mean)
		}
		if Median {
			median = medianCalc(arr, int(i))
			fmt.Printf("Median: %.2f\n", median)
		}
		if Mode {
			mode := modeCalc(arr)
			fmt.Printf("Mode: %d\n", mode)
		}
		if SD {
			sd = sdCalc(arr, mean, i)
			fmt.Printf("SD: %.2f\n", sd)
		}
	}
}

func medianCalc(arr []int, i int) (x float64) {
	if i%2 != 0 {
		x = float64(arr[i/2])
	} else {
		x = float64(arr[i/2]+arr[i/2-1]) / 2.0
	}
	return
}

func modeCalc(arr []int) int {
	m := map[int]int{}
	var mfreq, maxf int
	if len(arr) == 1 {
		return arr[0]
	}
	for i, a := range arr {
		m[a]++
		if i == 0 {
			maxf = a
			mfreq = 1
		} else {
			if m[a] > mfreq {
				maxf = a
				mfreq = m[a]
			} else if m[a] == mfreq && a < maxf {
				maxf = a
			}
		}
	}
	return maxf
}

func sdCalc(arr []int, m, i float64) (x float64) {
	for _, j := range arr {
		x += math.Pow(float64(j)-m, 2)
	}
	x = math.Sqrt(x / (i - 1))
	return
}
