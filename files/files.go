package files

import (
	"errors"
	"os"
	"os/exec"

	"github.com/4azoff/VKRPlayGround/config"

	"github.com/google/uuid"
)

func genLink() (string, error) {
	link, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}

	return link.String(), nil
}

func removeFile(filename string) {
	cmd := exec.Command("rm", "-rf", filename)
	cmd.Run()
}

func createFile(filename string) (string, error) {
	goFilename := filename + ".go"
	cmd := exec.Command("touch", goFilename)
	err := cmd.Run()
	if err != nil {
		return "", err
	}

	return goFilename, nil
}

func PrepareFile(data []byte, cnf *config.PlayConfig) (string, error) {
	link, err := genLink()
	if err != nil {
		return "", err
	}

	filename, err := createFile(link)
	if err != nil {
		return "", err
	}

	err = os.WriteFile(filename, data, 0700)
	if err != nil {
		return "", err
	}

	res, err := exec.Command("go", "build", "-o", cnf.BinariesFolder, filename).CombinedOutput()
	if err != nil {
		err = errors.New(string(res))
		removeFile(filename)
		return "", err
	}

	cmd := exec.Command("mv", filename, cnf.FilesFolder)
	err = cmd.Run()
	if err != nil {
		return "", err
	}

	return link, nil
}
