package encoding

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Yandex-Practicum/final-project-encoding-go/models"
	"gopkg.in/yaml.v3"
)

// JSONData тип для перекодирования из JSON в YAML
type JSONData struct {
	DockerCompose *models.DockerCompose
	FileInput     string
	FileOutput    string
}

// YAMLData тип для перекодирования из YAML в JSON
type YAMLData struct {
	DockerCompose *models.DockerCompose
	FileInput     string
	FileOutput    string
}

// MyEncoder интерфейс для структур YAMLData и JSONData
type MyEncoder interface {
	Encoding() error
}

// Encoding перекодирует файл из JSON в YAML
func (j *JSONData) Encoding() error {
	// ниже реализуйте метод
	// ...

	jsonFileContent, err := ioutil.ReadFile(j.FileInput)

	if err != nil {
		return fmt.Errorf("read json file fail: %w", err)
	}

	if j.DockerCompose == nil {
		j.DockerCompose = &models.DockerCompose{}
	}

	err = json.Unmarshal(jsonFileContent, j.DockerCompose)

	if err != nil {
		return fmt.Errorf("unmarshal json data fail: %w", err)
	}

	yamlFile, err := os.OpenFile(j.FileOutput, os.O_RDWR|os.O_CREATE, 0777)

	if err != nil {
		return fmt.Errorf("file open during yaml encoding fail: %w", err)
	}

	out, err := yaml.Marshal(j.DockerCompose)
	if err != nil {
		return fmt.Errorf("yaml encoding fail: %w", err)
	}

	_, err = yamlFile.Write(out)
	if err != nil {
		return fmt.Errorf("writing yaml data fail: %s", err)
	}

	return nil
}

// Encoding перекодирует файл из YAML в JSON
func (y *YAMLData) Encoding() error {
	// Ниже реализуйте метод
	// ...
	yamlFileContent, err := ioutil.ReadFile(y.FileInput)

	if err != nil {
		return fmt.Errorf("read yaml file fail: %w", err)
	}

	if y.DockerCompose == nil {
		y.DockerCompose = &models.DockerCompose{}
	}

	err = yaml.Unmarshal(yamlFileContent, y.DockerCompose)

	if err != nil {
		return fmt.Errorf("unmarshal yaml content fail: %w", err)
	}

	jsonFile, err := os.OpenFile(y.FileOutput, os.O_RDWR|os.O_CREATE, 0777)

	if err != nil {
		return fmt.Errorf("file open during json encoding fail: %w", err)
	}

	out, err := json.Marshal(y.DockerCompose)
	if err != nil {
		return fmt.Errorf("json encoding fail: %w", err)
	}

	_, err = jsonFile.Write(out)
	if err != nil {
		return fmt.Errorf("writing json data fail: %w", err)
	}

	return nil
}
