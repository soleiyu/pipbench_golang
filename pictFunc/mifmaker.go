	package pictFunc

import (
	"fmt"
	"os"
	"strconv"
)

func MifOutR(inp Pict, fn string) {
	pixelnum := inp.Width * inp.Height

	fmt.Println(pixelnum)
	fmt.Printf("%06X\n", pixelnum/ 4)

	writer(32, pixelnum, fn, inp)
}

func writer (width, depth int, fn string, data Pict) {
	file, _ := os.Create(fn)
	defer file.Close()

	file.Write(([]byte)("--  A____A\n"))
	file.Write(([]byte)("-- |・ㅅ・|\n"))
	file.Write(([]byte)("-- |っ　ｃ|\n"))
	file.Write(([]byte)("-- |　　　|\n"))
	file.Write(([]byte)("--  U￣￣U\n\n"))

	file.Write(([]byte)("WIDTH="))
	file.Write(([]byte)(strconv.Itoa(width)))
	file.Write(([]byte)(";\n"))
	file.Write(([]byte)("DEPTH="))
	file.Write(([]byte)(strconv.Itoa(depth)))
	file.Write(([]byte)(";\n\n"))

	file.Write(([]byte)("ADDRESS_RADIX=HEX;\n"))
	file.Write(([]byte)("DATA_RADIX=HEX;\n\n"))

	file.Write(([]byte)("CONTENT BEGIN\n"))

for i := 0; i < depth; i++ {
		h :=	fmt.Sprintf("  %06X  :  ", i / 4)
		file.Write(([]byte)(h))


		v1 := fmt.Sprintf("%02X",(data.Px[i % data.Width][i / data.Width][1]))
		i++
		v2 := fmt.Sprintf("%02X",(data.Px[i % data.Width][i / data.Width][1]))
		i++
		v3 := fmt.Sprintf("%02X",(data.Px[i % data.Width][i / data.Width][1]))
		i++
		v4 := fmt.Sprintf("%02X",(data.Px[i % data.Width][i / data.Width][1]))
		file.Write(([]byte)(v1))
		file.Write(([]byte)(v2))
		file.Write(([]byte)(v3))
		file.Write(([]byte)(v4))
		file.Write(([]byte)(";\n"))	
	}

	file.Write(([]byte)("END;\n"))
}

