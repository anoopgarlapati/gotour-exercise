package main

import (
	"fmt"
	"image"
	"image/color"
	"io"
	"math"
	"os"
	"strings"
	"time"

	"golang.org/x/tour/pic"
	"golang.org/x/tour/reader"
)

func main() {
	// Stringers
	pangolier := Person{"Donte Panlin", 6}
	darkwillow := Person{"Mireska Sunbreeze", 6}
	fmt.Println(pangolier, darkwillow)

	// Exercise: Stringers IPv4 Address
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}

	// Errors
	if err := run(); err != nil {
		fmt.Println(err)
	}

	// Exercise: Errors Square Root
	fmt.Println(nonNegativeSqrt(23))
	fmt.Println(nonNegativeSqrt(-23))

	// Readers
	r := strings.NewReader("Hello, Dio !!")
	read(r)

	// Exercise: Readers Infinite 'A's
	reader.Validate(MyReader{})

	// Exercise: Readers ROT13 Reader
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r2 := rot13Reader{s}
	io.Copy(os.Stdout, &r2)
	fmt.Println()

	// Images
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())

	// Exercise: Images Pic (result saved as images-pic.png for reference)
	img := Image{}
	pic.ShowImage(img)
}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

type IPAddr [4]byte

func (ipaddr IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ipaddr[0], ipaddr[1], ipaddr[2], ipaddr[3])
}

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{time.Now(), "it failed"}
}

func nonNegativeSqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := float64(1)
	delta := float64(2)
	for delta > 0.0000000001 {
		tmp := z
		z -= (z*z - x) / (2 * z)
		delta = math.Abs(z - tmp)
	}
	return z, nil
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Cannot square root a negative number: %v", float64(e))
}

func read(r io.Reader) {
	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v, err = %v, b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}

type MyReader struct{}

func (r MyReader) Read(buffer []byte) (int, error) {
	alloc := len(buffer)
	for i := 0; i < alloc; i++ {
		buffer[i] = 'A'
	}
	return alloc, nil
}

type rot13Reader struct {
	r io.Reader
}

func (rot13 *rot13Reader) Read(buffer []byte) (alloc int, err error) {
	alloc, err = rot13.r.Read(buffer)
	for i := 0; i < alloc; i++ {
		if buffer[i] >= 'A' && buffer[i] <= 'Z' {
			buffer[i] = 'A' + ((buffer[i]-'A')+13)%26
		} else if buffer[i] >= 'a' && buffer[i] <= 'z' {
			buffer[i] = 'a' + ((buffer[i]-'a')+13)%26
		}
	}
	return alloc, err
}

type Image struct{}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 100, 100)
}

func (img Image) At(x, y int) color.Color {
	return color.RGBA{uint8(x ^ y), uint8(x ^ y), 255, 255}
}
