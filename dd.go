package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

type Args struct {
	From string
	To string
	Offset int
	Limit int
}

func parseArgs() *Args {
	args := new(Args)

	flag.StringVar(&args.From, "from", "", "file to read from")
	flag.StringVar(&args.To, "to", "", "file to read from")
	flag.IntVar(&args.Offset, "offset", 0, "offset in input file")
	flag.IntVar(&args.Limit, "limit", 0, "maximum number of bites to transfer")

	flag.Parse()

	return args
}

func validateArgs(args *Args) {
	if args.From == "" {
		log.Fatal("Please specify `from` argument")
	} else if args.To == "" {
		log.Fatal("Please specify `to` argument")
	} else if args.Limit < 0 {
		log.Fatal("Limit should be not negative")
	} else if args.Offset < 0 {
		log.Fatal("Offset should be not negative")
	}
}

func Copy(src io.ReadSeeker, dest io.Writer, offset int, limit int) (int, error) {
	var srcReader io.Reader = src

	if offset != 0 {
		_, err := src.Seek(int64(offset), 0)
		srcReader = src
		if err != nil {
			return 0, err
		}
	}

	if limit != 0 {
		srcReader = io.LimitReader(srcReader, int64(limit))
	}

	written, err := io.Copy(dest, srcReader)
	if err != nil {
		return 0, err
	}
	return int(written), nil
}

func CopyFiles(args *Args) (int, error) {
	source, err := os.Open(args.From)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	dest, err := os.OpenFile(args.To, os.O_TRUNC | os.O_WRONLY, 0644)
	if err != nil {
		if os.IsNotExist(err) {
			dest, err = os.Create(args.To)
			if err != nil {
				return 0, err
			}
		} else {
			return 0, err
		}
	}
	defer dest.Close()

	return Copy(source, dest, args.Offset, args.Limit)
}

// -from /files/source.txt -to /files/dest.txt -offset 6 -limit 3
func main() {
	args := parseArgs()
	validateArgs(args)
	bytesWritten, err := CopyFiles(args)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Copied bytes: %v\n", bytesWritten)
}
