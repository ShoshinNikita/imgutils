package imgutils

import (
	"image"
	"image/jpeg"
	"log"
	"os"
	"strconv"
	"sync"
	"testing"
)

const (
	testdata = "testdata/"
	folder   = "tmp/"
)

func TestMain(m *testing.M) {
	if err := os.Mkdir(testdata, 0666); err != nil && !os.IsExist(err) {
		log.Fatalln(err)
	}
	if err := os.Mkdir(folder, 0666); err != nil && !os.IsExist(err) {
		log.Fatalln(err)
	}

	code := m.Run()

	if err := os.RemoveAll(folder); err != nil {
		log.Fatalln(err)
	}

	os.Exit(code)
}

func BenchmarkCrop(b *testing.B) {
	f, err := os.Open(testdata + "1.jpeg")
	if err != nil {
		log.Fatalln(err)
	}

	img, err := jpeg.Decode(f)
	if err != nil {
		log.Fatalln(err)
	}

	f.Close()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		min := image.Point{379, 32}
		max := image.Point{379 + 198, 32 + 147}
		newImg := Crop(img, min, max)

		filename := strconv.Itoa(i) + ".jpeg"
		newFile, _ := os.Create(folder + filename)
		jpeg.Encode(newFile, newImg, nil)
		newFile.Close()
	}
}

func BenchmarkCropInParallel(b *testing.B) {
	f, err := os.Open(testdata + "1.jpeg")
	if err != nil {
		log.Fatalln(err)
	}

	img, err := jpeg.Decode(f)
	if err != nil {
		log.Fatalln(err)
	}

	f.Close()

	min := image.Point{0, 0}
	max := image.Point{200, 200}

	wg := new(sync.WaitGroup)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		wg.Add(1)

		go func(id int) {
			newImg := Crop(img, min, max)

			filename := strconv.Itoa(id) + ".jpeg"
			newFile, _ := os.Create(folder + filename)
			jpeg.Encode(newFile, newImg, nil)
			newFile.Close()

			wg.Done()
		}(i)
	}

	wg.Wait()
}

func BenchmarkConcatenate(b *testing.B) {
	// There are must be files 1.jpeg and 2.jpeg

	// First image
	f, _ := os.Open(testdata + "1.jpeg")
	img1, err := jpeg.Decode(f)
	if err != nil {
		b.Errorf("can't encode file '1.jpg': %s", err)
	}
	f.Close()
	// Second image
	f, _ = os.Open(testdata + "2.jpeg")
	img2, err := jpeg.Decode(f)
	if err != nil {
		b.Errorf("can't encode file '1.jpg': %s", err)
	}
	f.Close()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		newImg := Concatenate(img1, img2, ConcatHorizontalMode)

		filename := strconv.Itoa(i) + ".jpeg"
		newFile, _ := os.Create(folder + filename)
		jpeg.Encode(newFile, newImg, nil)
		newFile.Close()
	}
}

func BenchmarkConcatenateInParallel(b *testing.B) {
	// There are must be files 1.jpeg and 2.jpeg

	// First image
	f, _ := os.Open(testdata + "1.jpeg")
	img1, err := jpeg.Decode(f)
	if err != nil {
		b.Errorf("can't encode file '1.jpg': %s", err)
	}
	f.Close()
	// Second image
	f, _ = os.Open(testdata + "2.jpeg")
	img2, err := jpeg.Decode(f)
	if err != nil {
		b.Errorf("can't encode file '1.jpg': %s", err)
	}
	f.Close()

	wg := new(sync.WaitGroup)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		wg.Add(1)

		go func(id int) {
			newImg := Concatenate(img1, img2, ConcatHorizontalMode)

			filename := strconv.Itoa(id) + ".jpeg"
			newFile, _ := os.Create(folder + filename)
			jpeg.Encode(newFile, newImg, nil)
			newFile.Close()

			wg.Done()
		}(i)
	}

	wg.Wait()
}
