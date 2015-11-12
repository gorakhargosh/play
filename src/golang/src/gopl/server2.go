package main

import (
	"fmt"
	"image"
	"image/color"
	"image/color/palette"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

var mu sync.Mutex
var count int

// var palette = []color.Color{
// 	color.White,
// 	color.Black,
// 	color.RGBA{0xcc, 0xff, 0xff, 0xff},
// }

const (
	cycles  = 5
	res     = 0.001
	size    = 100
	nframes = 64
	delay   = 8
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// lissajous writes a Lissajous animated GIF to the output.
func lissajous(out io.Writer) {
	pal := []color.Color{color.White}
	pal = append(pal, palette.WebSafe...)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, pal)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(rand.Intn(len(pal))))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}

func main() {
	http.HandleFunc("/count", counter)
	http.HandleFunc("/meta", handleMeta)
	http.HandleFunc("/lissajous", lissajousHandler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func lissajousHandler(w http.ResponseWriter, r *http.Request) {
	lissajous(w)
}

// handler mutates the counter.
func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	log.Print("hmmmm")
	fmt.Fprintf(w, "URL path = %q\n", r.URL.Path)
}

// counter displays the current count.
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

func handleMeta(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}
