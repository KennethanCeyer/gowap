package squeak

import (
	"encoding/json"
	"github.com/golang/snappy"
	"os"
	"path"
)

type Nuts struct {
	Private string
	Public  string
}

func (n *Nuts) Save(archive string) error {
	archivePath := path.Join(ProjectPath, NutsPath, archive)
	jsonData, _ := json.Marshal(n)
	var data []byte
	snappy.Encode(jsonData, data)

	file, err := os.OpenFile(archivePath, os.O_WRONLY, 0644)

	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(data)
	return err
}
