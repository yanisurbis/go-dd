package main

import (
	"flag"
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
		log.Fatal("Please specify `from` argument.")
	} else if args.To == "" {
		log.Fatal("Please specify `to` argument")
	} else if args.Limit < 0 {
		log.Fatal("Limit should be not negative")
	} else if args.Offset < 0 {
		log.Fatal("Offset should be not negative")
	}
}

func Copy(args *Args) error {
	source, err := os.Open(args.From)
	if err != nil {
		return err
	}
	defer source.Close()

	dst, err := os.OpenFile(args.To, os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer source.Close()

	sourceReader, err := func() (io.Reader, error) {
		if args.Offset != 0 {
			_, err := source.Seek(int64(args.Offset), 0)
			if err != nil {
				return nil, err
			}
		}

		if args.Limit != 0 {
			return io.LimitReader(source, int64(args.Limit)), nil
		}

		return source, nil
	}()
	if err != nil {
		return err
	}

	_, err = io.Copy(dst, sourceReader)
	return err
}

func main() {
	args := parseArgs()
	validateArgs(args)
	err := Copy(args)
	if err != nil {
		log.Fatal(err)
	}
}
