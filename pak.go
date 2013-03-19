package pak

import (
	"fmt"
	"io/ioutil"
	"launchpad.net/goyaml"
	"os"
	"path/filepath"
    . "github.com/theplant/pak/share"
    "github.com/theplant/pak/gitpkg"
)

type PakPkg struct {
    gitpkg.GitPkg
    GetOption
}

func GetPakInfo() (pakInfo PakInfo, err error) {
	var pakInfoBytes []byte
	pakInfoBytes, err = pakRead(Pakfile)
	if err != nil {
		return
	}

	err = goyaml.Unmarshal(pakInfoBytes, &pakInfo)

	return
}

func GetPaklockInfo() (paklockInfo PaklockInfo, err error) {
	var content []byte
	content, err = pakRead(Paklock)
	if err != nil {
		return
	}

	err = goyaml.Unmarshal(content, &paklockInfo)

	return
}

// var ErrorNotFound = errors.New("Can't find.")
func pakRead(file string) (fileContent []byte, err error) {
	absPakfilePath := ""
	originalFile := file
	for true {
		absPakfilePath, err = filepath.Abs(file)
		if err != nil {
			return
		}
		if absPakfilePath == Gopath+"/Pakfile" {
			return nil, fmt.Errorf("Can't find %s" + originalFile)
		}

		_, err = os.Stat(file)
		if os.IsNotExist(err) {
			file = "../" + file
			continue
		}

		break
	}

	return ioutil.ReadFile(file)
}

func writePaklockInfo(paklockInfo PaklockInfo) error {
	content, err := goyaml.Marshal(&paklockInfo)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(Paklock, content, os.FileMode(0644))
}
