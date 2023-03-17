package config

import (
	"bytes"
	"encoding/json"
	"os"
	"sync"
)

type File[TConfig any] struct {
	filename string
	Config   TConfig
	mu       sync.Mutex
}

func NewFile[TConfig any](filename string, defaults TConfig) *File[TConfig] {
	cf := &File[TConfig]{
		filename: filename,
		Config:   defaults,
	}

	_ = cf.Read()

	return cf
}

func (c *File[TConfig]) Read() error {
	cfb, err := os.ReadFile(c.filename)
	if err != nil {
		return err
	}
	return json.NewDecoder(bytes.NewReader(cfb)).Decode(&c.Config)
}

func (c *File[TConfig]) Write() error {
	c.mu.Lock()
	defer c.mu.Unlock()
	cfb := &bytes.Buffer{}
	if err := json.NewEncoder(cfb).Encode(&c.Config); err != nil {
		return err
	}
	return os.WriteFile(c.filename, cfb.Bytes(), os.ModePerm)
}

func (c *File[TConfig]) Flush() error {
	return c.Write()
}
