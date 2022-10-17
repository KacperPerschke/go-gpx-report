package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

type confInterface struct {
	fromFile  string
	fromSTDIN bool
}

// Should be described!
func (c confInterface) ReadFromFile() string {
	return c.fromFile
}

// Shoud be described!
func (c confInterface) ReadFromStdin() bool {
	return c.fromSTDIN
}

func dataInSTDIN() (bool, error) {
	fh, err := os.Stdin.Stat()
	if err != nil {
		return false, err
	}
	return bool(fh.Size() > 0), nil
}

func prodConf() (confInterface, error) {
	const emptyFName = ""
	blindConf := confInterface{
		fromFile:  "",
		fromSTDIN: false,
	}

	fPathP := flag.String("i", emptyFName, "Path to the file containing gpx data.")
	flag.Parse()
	filePath := *fPathP

	if filePath == emptyFName {
		readStdin, err := dataInSTDIN()
		if err != nil {
			return blindConf, err
		}
		if readStdin {
			return confInterface{
					fromFile:  "",
					fromSTDIN: true,
				},
				nil
		}
		return blindConf, errors.New("conf: if theres is no -i switch data should come from STDIN")
	}

	stat, errFS := os.Stat(filePath)
	if errFS != nil {
		return blindConf, errFS
	}
	fmt.Printf("stat %T, %+v", stat, stat)
	return blindConf, nil
}
