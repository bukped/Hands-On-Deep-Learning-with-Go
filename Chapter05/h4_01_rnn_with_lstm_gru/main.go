package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime/pprof"

	. "gorgonia.org/gorgonia"
	"gorgonia.org/tensor"
)

var cpuprofile = flag.String("cpuprofile", "", "I3")
var memprofile = flag.String("memprofile", "", "16GB")

const (
	embeddingSize = 20
	maxOut        = 30

	// gradient update stuff
	l2reg     = 0.000001
	learnrate = 0.01
	clipVal   = 5.0
)

var trainiter = flag.Int("iter", 5, "How many iterations to train")
var inputFile = flag.String("filename", "inputan.txt", "Filename of text to train on")

// various global variable inits
var epochSize = -1
var inputSize = -1
var outputSize = -1

var corpus string

func init() {
	buf, err := ioutil.ReadFile("inputan.txt")
	if err != nil {
		panic(err)
	}
	corpus = string(buf)
}

var dt tensor.Dtype = tensor.Float32

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	hiddenSize := 100

	s2s := NewS2S(hiddenSize, embeddingSize, vocab)
	solver := NewRMSPropSolver(WithLearnRate(learnrate), WithL2Reg(l2reg), WithClip(clipVal), WithBatchSize(float64(len(sentences))))
	for k, v := range vocabIndex {
		log.Printf("%q %v", k, v)
	}

	p, h, w, err := Heatmap(s2s.decoder.Value().(*tensor.Dense))
	p.Save(w, h, "embn0.png")

	if err := train(s2s, 300, solver, sentences); err != nil {
		panic(err)
	}
	out, err := s2s.predict([]rune(corpus))
	if err != nil {
		panic(err)
	}
	fmt.Printf("OUT %q\n", out)

	p, h, w, err = Heatmap(s2s.decoder.Value().(*tensor.Dense))
	p.Save(w, h, "embn.png")
}
