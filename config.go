package main

import "os"

type confInterface struct {
	inFileHndl *os.File
}

func dataInSTDIN() (bool, error) {
	fh, err := os.Stdin.Stat()
	if err != nil {
		return false, err
	}
	return bool(fh.Size() > 0), nil
}

func prodConf() (confInterface, error) {
	got, err := dataInSTDIN()
	if err != nil {
		return confInterface{}, err
	}
	if !got {
		return confInterface{}, nil
	}
	return confInterface{inFileHndl: os.Stdin}, nil
}
