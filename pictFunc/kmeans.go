package pictFunc

import (
	"fmt"
	"math/rand"
	"time"
)

func Kmeans256(a Pict, hol int) Pict {
	res := MkPict(a.Width, a.Height)

	cp := make([][]uint8, 256)
	cnt := 0
	
	//STEP0 : GENERATE RAND POINT
	rand.Seed(time.Now().UnixNano())	
	for i := 0; i < len(cp); i++ {
		cp[i] = make([]uint8, 4)
		cp[i][0] = 255
		cp[i][1] = uint8(rand.Intn(256))
		cp[i][2] = uint8(rand.Intn(256))
		cp[i][3] = uint8(rand.Intn(256))
	}

	for ;; {

		gcol := make([][]uint, len(cp))
		gnum := make([]uint, len(cp))

		for n := 0; n < len(cp); n++ {
			gcol[n] = make([]uint, 4)
			gcol[n][0] = 255
			gcol[n][1] = 0
			gcol[n][2] = 0
			gcol[n][3] = 0
			gnum[n] = 0
		}

		//STEP1 : que Distance and shozoku
		for x := 0; x < res.Width; x+=1 {
			for y := 0; y < res.Height; y+=1 {

				d := dist(a.Px[x][y], cp[0])
				dp := 0;
				dn := 0;

				for i := 1; i < len(cp); i++ {
					dp = dist(a.Px[x][y], cp[i])
					if dp < d {
						d = dp
						dn = i
					}
				}

				res.Px[x][y][0] = uint8(dn)

				gcol[dn][1] += uint(a.Px[x][y][1])
				gcol[dn][2] += uint(a.Px[x][y][2])
				gcol[dn][3] += uint(a.Px[x][y][3])
				gnum[dn] ++
			}
		}

		//STEP2 : que G re cp

		chnum := 0

		for i := 0; i < len(cp); i++ {
			if gnum[i] != 0 {
				gcol[i][1] /= gnum[i]
				gcol[i][2] /= gnum[i]
				gcol[i][3] /= gnum[i]
			}

			if cp[i][1] != uint8(gcol[i][1]) ||
					cp[i][2] != uint8(gcol[i][2]) ||
					cp[i][3] != uint8(gcol[i][3]) {
				chnum ++
			}
		}
		
		pcp := [][]uint8{}

		for n := 0; n < len(cp); n++ {
			if gnum[n] != 0 {
				ppcp := []uint8{255, 
					uint8(gcol[n][1]), uint8(gcol[n][2]), uint8(gcol[n][3])}
				pcp = append(pcp, ppcp)
			}	else {
				ppcp := []uint8{255, 
					uint8(rand.Intn(256)),
					uint8(rand.Intn(256)),
					uint8(rand.Intn(256))}
				pcp = append(pcp, ppcp)
			}
		}
		cp = pcp

		fmt.Println(chnum, len(cp))

		cnt ++
		if hol < cnt {
			break
		}
	}

	//STEP1 : que Distance and shozoku
	for x := 0; x < res.Width; x++ {
		for y := 0; y < res.Height; y++ {
			d := dist(a.Px[x][y], cp[0])
			dp := 0;
			dn := 0;

			for i := 1; i < len(cp); i++ {
				dp = dist(a.Px[x][y], cp[i])
				if dp < d {
					d = dp
					dn = i
				}
			}
			res.Px[x][y] = cp[dn]
			res.Px[x][y][0] = 255
		}
	}

	return res
}

func KColi(a Pict, hol, rnum, gnum, bnum int) Pict {
	res := MkPict(a.Width, a.Height)

	cpr := make([]uint8, rnum)
	cpg := make([]uint8, gnum)
	cpb := make([]uint8, bnum)

	//STEP0 : GENERATE COLOR POINT
	for i := 0; i < rnum; i++ {
		cpr[i] = uint8(255 * (i + 1) / (rnum + 1))
	}
	for i := 0; i < gnum; i++ {
		cpg[i] = uint8(255 * (i + 1) / (gnum + 1))
	}
	for i := 0; i < bnum; i++ {
		cpb[i] = uint8(255 * (i + 1) / (bnum + 1))
	}

	for ;; {
		
		sumr := make([]uint, rnum)
		sumg := make([]uint, gnum)
		sumb := make([]uint, bnum)
		
		cntr := make([]uint, rnum)
		cntg := make([]uint, gnum)
		cntb := make([]uint, bnum)
			
		for i := 0; i < rnum; i++ {
			sumr[i] = 0
			cntr[i] = 0
		}
		for i := 0; i < gnum; i++ {
			sumg[i] = 0
			cntg[i] = 0
		}
		for i := 0; i < bnum; i++ {
			sumb[i] = 0
			cntb[i] = 0
		}

		//STEP1 : que Distance and shozoku
		for x := 0; x < res.Width; x+=1 {
			for y := 0; y < res.Height; y+=1 {

				dr := uint8(255)
				dg := uint8(255)
				db := uint8(255)

				dcr := 0
				dcg := 0
				dcb := 0

				for i := 0; i < rnum; i++ {
					dpr := abs(a.Px[x][y][1], cpr[i])

					if dpr < dr {
							dr = dpr
							dcr = i
					}
				}

				for i := 0; i < gnum; i++ {
					dpg := abs(a.Px[x][y][2], cpg[i])

					if dpg < dg {
							dg = dpg
							dcg = i
					}
				}

				for i := 0; i < bnum; i++ {
					dpb := abs(a.Px[x][y][3], cpb[i])

					if dpb < db {
							db = dpb
							dcb = i
					}
				}

				sumr[dcr] += uint(a.Px[x][y][1])
				sumg[dcg] += uint(a.Px[x][y][2])
				sumb[dcb] += uint(a.Px[x][y][3])

				cntr[dcr] ++
				cntg[dcg] ++
				cntb[dcb] ++
			}
		}

		//STEP2 : que G re cp

		chnum := 0

		for i := 0; i < rnum; i++ {
			if cntr[i] != 0 {
				sumr[i] /= cntr[i]
			}
			if cpr[i] != uint8(sumr[i]) {
				chnum ++
				cpr[i] = uint8(sumr[i])
			}
		}

		for i := 0; i < gnum; i++ {
			if cntg[i] != 0 {
				sumg[i] /= cntg[i]
			}
			if cpg[i] != uint8(sumg[i]) {
				chnum ++
				cpg[i] = uint8(sumg[i])
			}
		}

		for i := 0; i < bnum; i++ {
			if cntb[i] != 0 {
				sumb[i] /= cntb[i]
			}
			if cpb[i] != uint8(sumb[i]) {
				chnum ++
				cpb[i] = uint8(sumb[i])
			}
		}

		fmt.Println(chnum)

		if chnum < hol {
			break
		}
	}

	shcr := make([]uint, rnum)
	shcg := make([]uint, gnum)
	shcb := make([]uint, bnum)

	//STEP1 : que Distance and shozoku
	for x := 0; x < res.Width; x++ {
		for y := 0; y < res.Height; y++ {
			dr := uint8(255)
			dg := uint8(255)
			db := uint8(255)
			cr := 0
			cg := 0
			cb := 0

			for i := 0; i < rnum; i++ {
				dpr := abs(a.Px[x][y][1], cpr[i])
				
				if dpr < dr {
					dr = dpr
					cr = i
				}
			}

			for i := 0; i < gnum; i++ {
				dpg := abs(a.Px[x][y][2], cpg[i])

				if dpg < dg {
					dg = dpg
					cg = i
				}
			}

			for i := 0; i < bnum; i++ {
				dpb := abs(a.Px[x][y][3], cpb[i])
				
				if dpb < db {
					db = dpb
					cb = i
				}
			}

			shcr[cr] ++
			shcg[cg] ++
			shcb[cb] ++

			res.Px[x][y][1] = cpr[cr]
			res.Px[x][y][2] = cpg[cg]
			res.Px[x][y][3] = cpb[cb]
			res.Px[x][y][0] = 255
		}
	}

	fmt.Println("vr ; ", cpr)
	fmt.Println("vg ; ", cpg)
	fmt.Println("vb ; ", cpb)

	fmt.Printf("cr : ")
	for i := 0; i < rnum; i++ {
		fmt.Printf("%.1f, ", float32(100) * float32(shcr[i]) / float32(res.Width * res.Height))
	}
	fmt.Printf("\ncg : ")
	for i := 0; i < gnum; i++ {
		fmt.Printf("%.1f, ", float32(100) * float32(shcg[i]) / float32(res.Width * res.Height))
	}
	fmt.Printf("\ncb : ")
	for i := 0; i < bnum; i++ {
		fmt.Printf("%.1f, ", float32(100) * float32(shcb[i]) / float32(res.Width * res.Height))
	}
	fmt.Println()

	return res
}


func KCol(a Pict, hol, cnum int) Pict {
	res := MkPict(a.Width, a.Height)

	cpr := make([]uint8, cnum)
	cpg := make([]uint8, cnum)
	cpb := make([]uint8, cnum)

	//STEP0 : GENERATE COLOR POINT
	for i := 0; i < cnum; i++ {
		cpr[i] = uint8(255 * (i + 1) / (cnum + 1))
		cpg[i] = uint8(255 * (i + 1) / (cnum + 1))
		cpb[i] = uint8(255 * (i + 1) / (cnum + 1))
	}

	for ;; {
		
		sumr := make([]uint, cnum)
		sumg := make([]uint, cnum)
		sumb := make([]uint, cnum)
		
		cntr := make([]uint, cnum)
		cntg := make([]uint, cnum)
		cntb := make([]uint, cnum)
	
		for i := 0; i < cnum; i++ {
			sumr[i] = 0
			sumg[i] = 0
			sumb[i] = 0
			cntr[i] = 0
			cntg[i] = 0
			cntb[i] = 0
		}

		//STEP1 : que Distance and shozoku
		for x := 0; x < res.Width; x+=1 {
			for y := 0; y < res.Height; y+=1 {

				dr := uint8(255)
				dg := uint8(255)
				db := uint8(255)

				dcr := 0
				dcg := 0
				dcb := 0

				for i := 0; i < cnum; i++ {
					dpr := abs(a.Px[x][y][1], cpr[i])
					dpg := abs(a.Px[x][y][2], cpg[i])
					dpb := abs(a.Px[x][y][3], cpb[i])

					if dpr < dr {
							dr = dpr
							dcr = i
					}
					if dpg < dg {
							dg = dpg
							dcg = i
					}
					if dpb < db {
							db = dpb
							dcb = i
					}
				}

				sumr[dcr] += uint(a.Px[x][y][1])
				sumg[dcg] += uint(a.Px[x][y][2])
				sumb[dcb] += uint(a.Px[x][y][3])

				cntr[dcr] ++
				cntg[dcg] ++
				cntb[dcb] ++
			}
		}

		//STEP2 : que G re cp

		chnum := 0

		for i := 0; i < cnum; i++ {
			if cntr[i] != 0 {
				sumr[i] /= cntr[i]
			}
			if cntg[i] != 0 {
				sumg[i] /= cntg[i]
			}
			if cntb[i] != 0 {
				sumb[i] /= cntb[i]
			}

			if cpr[i] != uint8(sumr[i]) {
				chnum ++
				cpr[i] = uint8(sumr[i])
			}
			if cpg[i] != uint8(sumg[i]) {
				chnum ++
				cpg[i] = uint8(sumg[i])
			}
			if cpb[i] != uint8(sumb[i]) {
				chnum ++
				cpb[i] = uint8(sumb[i])
			}
		}

		fmt.Println(chnum)

		if chnum < hol {
			break
		}
	}

	shcr := make([]uint, cnum)
	shcg := make([]uint, cnum)
	shcb := make([]uint, cnum)

	//STEP1 : que Distance and shozoku
	for x := 0; x < res.Width; x++ {
		for y := 0; y < res.Height; y++ {
			dr := uint8(255)
			dg := uint8(255)
			db := uint8(255)
			cr := 0
			cg := 0
			cb := 0

			for i := 0; i < cnum; i++ {
				dpr := abs(a.Px[x][y][1], cpr[i])
				dpg := abs(a.Px[x][y][2], cpg[i])
				dpb := abs(a.Px[x][y][3], cpb[i])
				
				if dpr < dr {
					dr = dpr
					cr = i
				}
				if dpg < dg {
					dg = dpg
					cg = i
				}
				if dpb < db {
					db = dpb
					cb = i
				}
			}

			shcr[cr] ++
			shcg[cg] ++
			shcb[cb] ++

			res.Px[x][y][1] = cpr[cr]
			res.Px[x][y][2] = cpg[cg]
			res.Px[x][y][3] = cpb[cb]
			res.Px[x][y][0] = 255
		}
	}

	fmt.Println("vr ; ", cpr)
	fmt.Println("vg ; ", cpg)
	fmt.Println("vb ; ", cpb)

	fmt.Printf("cr : ")
	for i := 0; i < cnum; i++ {
		fmt.Printf("%.1f, ", float32(100) * float32(shcr[i]) / float32(res.Width * res.Height))
	}
	fmt.Printf("\ncg : ")
	for i := 0; i < cnum; i++ {
		fmt.Printf("%.1f, ", float32(100) * float32(shcg[i]) / float32(res.Width * res.Height))
	}
	fmt.Printf("\ncb : ")
	for i := 0; i < cnum; i++ {
		fmt.Printf("%.1f, ", float32(100) * float32(shcb[i]) / float32(res.Width * res.Height))
	}
	fmt.Println()

	return res
}

func dist(a, b []uint8) int {
	dx := uint(a[1] - b[1]) * uint(a[1] - b[1])
	dy := uint(a[2] - b[2]) * uint(a[2] - b[2])
	dz := uint(a[3] - b[3]) * uint(a[3] - b[3])

	return int(dx + dy + dz)
}

func abs(a, b uint8) uint8 {
	if a < b {
		return b - a
	} else {
		return a - b
	}
}
