// THERMAL Benchmark

package main

import (
	"fmt"
	"time"
	"sync"
	"os/exec"
	"./pictFunc"
)

func main(){
	t0 := time.Now()

	fflag := true
	cnt := 0.0

	var wg sync.WaitGroup
	wg.Add(1)

//## BEFORE COOLING
	for i := 0; i < 20; i ++ {
		freq, temp := singleMeas()
		fmt.Printf("%.1f %s %s\n", cnt, freq, temp)
		cnt += 0.5
		pt := time.Now()
		sec := (pt.Sub(t0)).Seconds()
		ms := int(1000.0 * sec) % 500
		next := 500 - int(ms)

		time.Sleep(time.Millisecond * time.Duration(next))
	}

//## CALCULATION
	go func () {
		outFile := pictFunc.MkMandelMapQ_pr(5120, 2880, 12, 0.016, -0.667, 0.009, 0, 0, 0.36)
		outFile.Save("res.png")
		fflag = false
	}()

	go func () {
		for ; fflag; {
			freq, temp := singleMeas()
			fmt.Printf("%.1f %s %s\n", cnt, freq, temp)
			cnt += 0.5
			pt := time.Now()
			sec := (pt.Sub(t0)).Seconds()
			ms := int(1000.0 * sec) % 500
			next := 500 - int(ms)

			time.Sleep(time.Millisecond * time.Duration(next))
		}
		wg.Done()
	}()

	wg.Wait()

//## AFTER COOLING
	for i := 0; i < 20; i ++ {
		freq, temp := singleMeas()
		fmt.Printf("%.1f %s %s\n", cnt, freq, temp)
		cnt += 0.5
		pt := time.Now()
		sec := (pt.Sub(t0)).Seconds()
		ms := int(1000.0 * sec) % 500
		next := 500 - int(ms)

		time.Sleep(time.Millisecond * time.Duration(next))
	}


//	fmt.Println("FINISH")
}

func singleMeas () (string, string) {
	freq, _ := exec.Command("perl", "freq.pl").Output()
	temp, _ := exec.Command("perl", "temp.pl").Output()

	return string(freq), string(temp)
}

