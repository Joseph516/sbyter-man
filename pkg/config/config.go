package config

import (
	"github.com/spf13/viper"
	"path"
)

type ViperConfig struct {
	vp *viper.Viper
}

func NewViperConfig(fpath string) (*ViperConfig, error) {
	ext := path.Ext(fpath)[1:]

	vp := viper.New()
	vp.AddConfigPath(path.Dir(fpath))
	vp.SetConfigName(path.Base(fpath))
	vp.SetConfigType(ext)
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return &ViperConfig{vp}, nil
}

func (s *ViperConfig) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}

	return nil
}
