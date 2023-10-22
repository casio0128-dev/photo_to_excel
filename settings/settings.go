package settings

import (
	"golang.org/x/crypto/openpgp/errors"
	"gopkg.in/yaml.v3"
	"io"
	"os"
	"photo2excel/commons"
)

type Setting struct {
	Version  string   `yaml:"version"`
	Settings Settings `yaml:"settings"`
}

type Settings struct {
	PhotoDirectories []string `yaml:"photo_directories"`
	OutputDirectory  string   `yaml:"output_directory"`
}

func New() (*Setting, error) {
	settingFilePath := os.Getenv("SETTING_FILE_PATH")
	settingFile, err := os.OpenFile(settingFilePath, os.O_RDONLY, 0755)
	if err != nil {
		return nil, err
	}
	defer settingFile.Close()
	readYAML, err := io.ReadAll(settingFile)
	if err != nil {
		return nil, err
	}

	setting := Setting{}
	if err := yaml.Unmarshal(readYAML, &setting); err != nil {
		return nil, err
	}

	return &setting, nil
}

func (s *Setting) Get(dirType commons.DirType) any {
	switch dirType {
	case commons.PhotoDir:
		return s.Settings.PhotoDirectories
	case commons.OutputDir:
		return s.Settings.OutputDirectory
	default:
		panic(errors.InvalidArgumentError("不正なSettingsのディレクトリタイプを取得することはできません。"))
	}
}
