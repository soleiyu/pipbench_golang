package pictFunc

import (
	"fmt"
	"os"
	"image"
	"image/png"
	"image/color"
)

type Pict struct{
	Width int
	Height int
	Px [][][]uint8
}

func (this *Pict) ShowSize(){
    fmt.Printf("%d x %d\n", this.Width, this.Height)
}

func (this *Pict) Load(fn string){
	file, _ := os.Open(fn)
	defer	file.Close()
	file2, _ := os.Open(fn)
	defer	file2.Close()

	cfg, _, _ := image.DecodeConfig(file)
	img, _, err := image.Decode(file2)
	if err != nil{
		fmt.Println("IMG ERROR")
	}

	this.Width = cfg.Width
	this.Height = cfg.Height

	this.Px = make([][][]uint8, this.Width)
	for x := 0; x < this.Width; x++ {
	  this.Px[x] = make([][]uint8, this.Height)
		for y := 0; y < this.Height; y++{
			this.Px[x][y] = make([]uint8, 4)
		}
	}

	for x := 0; x < this.Width; x++ {
		for y := 0; y < this.Height; y++ {
		  r, g, b, a := img.At(x, y).RGBA()
			this.Px[x][y][0] = uint8(a)
			this.Px[x][y][1] = uint8(r)
			this.Px[x][y][2] = uint8(g)
			this.Px[x][y][3] = uint8(b)
		}
	}

	this.ShowSize()
}

func (this *Pict) Save(fn string){
//	outRect := image.Rectangle{image.Pt(0, 0), dstImg.Bounds().Size()}
	out := image.NewRGBA(image.Rect(0, 0, this.Width, this.Height))
	var c color.RGBA	
	
	for x := 0; x < this.Width; x++ {
		for y := 0; y < this.Height; y++ {
			c = mkCol(this.Px[x][y][0], this.Px[x][y][1], this.Px[x][y][2], this.Px[x][y][3])
			out.Set(x, y, c)
		}
	}

	outfile, _ := os.Create(fn)
  defer outfile.Close()
    // 書き出し
  png.Encode(outfile, out)
}

func MkPict(w int, h int) Pict{
	var p Pict
	p.Width = w
	p.Height = h

	p.Px = make([][][]uint8, w)
	for x := 0; x < w; x++ {
	  p.Px[x] = make([][]uint8, h)
		for y := 0; y < h; y++ {
			p.Px[x][y] = make([]uint8, 4)
		}
	}	

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			p.Px[x][y][0] = 0
			p.Px[x][y][1] = 0
			p.Px[x][y][2] = 0
			p.Px[x][y][3] = 0
		}
	}
	return p
}

func mkCol(a uint8, r uint8, g uint8, b uint8) color.RGBA{
	var c color.RGBA

	c.A = a
	c.R = r
	c.G = g
	c.B = b

	return c
}
