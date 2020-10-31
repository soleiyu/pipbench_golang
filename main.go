// THERMAL Benchmark

package main

import (
	"fmt"
	"time"
	"./pictFunc"
)

func main(){
	t0 := time.Now()

	outFile := pictFunc.MkMandelMapQ_pr(3840, 2160, 4, 0.016, -0.667, 0.009, 0, 0, 0.36)
	t1 := time.Now()

	outFile.Save("res.png")
	fmt.Printf("%f\n", (t1.Sub(t0)).Seconds())

//	fmt.Println("FINISH")
}

