package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/mumax/3/data"
	"github.com/mumax/3/util"
)

func resize(f *data.Slice, arg string) {
	s := parseSize(arg)
	resized := data.Resize(f, s)
	*f = *resized
}

func parseSize(arg string) (size [3]int) {
	words := strings.Split(arg, "x")
	if len(words) != 3 {
		log.Fatal("resize: need N0xN1xN2 argument")
	}
	for i, w := range words {
		v, err := strconv.Atoi(w)
		util.FatalErr(err)
		size[i] = v
	}
	return
}

func Resize(in *data.Slice, N [3]int) *data.Slice{
        if in.Size() == N{
                return in // nothing to do
        }
        In := in.Tensors()
        out := data.NewSlice(in.NComp(), N)
        Out := out.Tensors()

        srcsize := data.SizeOf(In[0])
        dstsize := data.SizeOf(Out[0])

        Dx := dstsize[X]
        Dy := dstsize[Y]
        Dz := dstsize[Z]
        Sx := srcsize[X]
        Sy := srcsize[Y]
        Sz := srcsize[Z]
        scalex := Sx / Dx
        scaley := Sy / Dy
        scalez := Sz / Dz
        util.Assert(scalex > 0 && scaley > 0)

        for c := range Out {

                for iz := 0; iz < Dz; iz++ {
                        for iy := 0; iy < Dy; iy++ {
                                for ix := 0; ix < Dx; ix++ {
                                        sum, n := 0.0, 0.0

                                        for I := 0; I < scalez; I++ {
                                                i2 := iz*scalez + I
                                                for J := 0; J < scaley; J++ {
                                                        j2 := iy*scaley + J
                                                        for K := 0; K < scalex; K++ {
                                                                k2 := ix*scalex + K

                                                                if i2 < Sz && j2 < Sy && k2 < Sx {
                                                                        sum += float64(In[c][i2][j2][k2])
                                                                        n++
                                                                }
                                                        }
                                                }
                                        }
                                        Out[c][iz][iy][ix] = float32(sum / n)
                                }
                        }
                }
        }

        return out
}
