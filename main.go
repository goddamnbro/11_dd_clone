package main

import (
	"bytes"
	"flag"
	"math"
	"os"
	"time"

	"github.com/schollz/progressbar"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// flags
	fromPath := *flag.String("from", "samples/from.txt", "a string")
	toPath := *flag.String("to", "samples/to.txt", "a string")

	// open file for reading
	f, err := os.Open(fromPath)
	check(err)
	defer f.Close()

	// get FileInfo object for the size property
	fi, err := f.Stat()
	check(err)
	fileSize := int(fi.Size())

	// more flags
	offset := *flag.Int("offset", 0, "an int")
	limit := *flag.Int("limit", int(fileSize), "an int")
	if limit > int(fileSize) {
		limit = int(fileSize)
	}

	// create buffer and copy content to buffer
	buf := make([]byte, fileSize)
	_, err = f.Read(buf)
	check(err)

	// create or open file for writing
	newFile, err := os.Create(toPath)
	check(err)
	defer newFile.Close()

	step := bytes.MinRead
	d := float64(limit) / float64(step)
	pages := int(math.Ceil(d))

	bar := progressbar.New(pages)
	for i := 0; i < pages; i++ {
		if offset+step > limit {
			step = limit - offset
		}

		newFile.Write(buf[offset : offset+step])
		offset += step

		bar.Add(1)
		// for demonstration only
		time.Sleep(100 * time.Millisecond)
	}
}
