package settings

import (
	"golang.org/x/crypto/openpgp/errors"
	"gopkg.in/yaml.v3"
	"io"
	"os"
)

type Setting struct {
	PhotoDirectories []string `yaml:"photo_directories"`
	OutputDirectory  string   `yaml:"output_directory"`
}

func New() (*Setting, error) {
	settingFilePhat := os.Getenv("SETTING_FILE_PATH")
	settingFile, err := os.OpenFile(settingFilePhat, os.O_RDONLY, 0755)
	if err != nil {
		return nil, err
	}
	defer settingFile.Close()
	readYAML, err := io.ReadAll(settingFile)
	if err != nil {
		return nil, err
	}

	var setting *Setting
	if err := yaml.Unmarshal(readYAML, setting); err != nil {
		return nil, err
	}

	return setting, nil
}

func (s *Setting) Get(dirType dirType) any {
	switch dirType {
	case PhotoDir:
		return s.PhotoDirectories
	case OutputDir:
		return s.OutputDirectory
	default:
		panic(errors.InvalidArgumentError("不正なSettingsのディレクトリタイプを取得することはできません。"))
	}
}
