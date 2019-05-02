package squeak

import (
	"encoding/base64"
	"encoding/json"
	"github.com/golang/snappy"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"path/filepath"
)

type Nuts struct {
	Private string
	Public  string
}

func (n *Nuts) Save(nutsPath string) error {
	keyStore := make([]string, 2)
	for key, keyPath := range []string{n.Private, n.Public} {
		keyData, err := ioutil.ReadFile(keyPath)
		if err != nil {
			return err
		}
		keyStore[key] = string(keyData)
	}

	jsonData, _ := json.Marshal(keyStore)
	var buf []byte
	buf = snappy.Encode(buf, jsonData)
	data := make([]byte, base64.StdEncoding.EncodedLen(len(buf)))
	base64.StdEncoding.Encode(data, buf)

	log.WithFields(log.Fields{
		"nuts": filepath.Dir(nutsPath),
	}).Info("new nuts has been added")

	return ioutil.WriteFile(nutsPath, data, 0644)
}
