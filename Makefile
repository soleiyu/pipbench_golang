default:
	go run thermalbench.go > plot
	gnuplot plot.txt
