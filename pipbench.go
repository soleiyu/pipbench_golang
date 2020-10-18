// Parallel Image Processing Benchmark

package main

import (
	"fmt"
	"time"
	"./pictFunc"
)

func main(){
	start := time.Now()
	outFile := pictFunc.MkMandelMapQ_pr(3840, 2160, 12, 0.016, -0.667, 0.009, 0, 0, 0.36)
	end := time.Now()
	fmt.Printf("%0.2fsec\n", (end.Sub(start)).Seconds())

	outFile.Save("res.png")

}
