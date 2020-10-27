package external

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/pkg/errors"
)

var (
	execFilePath = ""
)

func CreateFile() {
	if file, err := os.Create("../exec"); err != nil {
		panic("error creating script file" + err.Error())
	} else {
		fmt.Fprint(file, execFile)

		wd, _ := os.Getwd()
		execFilePath = filepath.Dir(wd)
		execFilePath = filepath.Join(execFilePath, "/exec")

		file.Close()
	}
}

func ExecFile() error {
	var exFile = execFilePath

	if exFile == "" {
		return errors.New("file is empty")
	}

	if file, err := os.Open(exFile); err != nil {
		return errors.Wrap(err, "open file")
	} else {
		defer file.Close()

		h := md5.New()

		if _, err = io.Copy(h, file); err != nil {
			return errors.Wrap(err, "read file")
		}

		if "7efc9dcd90660dfc5636c516f1788717" != fmt.Sprintf("%x", h.Sum(nil)) {
			return errors.New("comparison md5 error")
		}
	}

	execCmd := exec.Command("source", exFile)

	if execCmd == nil {
		return errors.New("execution error")
	}

	return nil
}
