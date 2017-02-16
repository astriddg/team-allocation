package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func persistDeptAndPers() error {
	deptFile, err := json.Marshal(departments)

	// Saving everything in files
	err = ioutil.WriteFile(fileNames["departments"], deptFile, os.ModeType)
	if err != nil {
		return err
	}

	persFile, err := json.Marshal(people)

	err = ioutil.WriteFile(fileNames["people"], persFile, os.ModeType)
	if err != nil {
		return err
	}
	return nil
}
