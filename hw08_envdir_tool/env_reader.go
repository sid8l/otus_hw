package main

import (
	"bufio"
	"bytes"
	"os"
	"path"
)

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

func ReadFile(file string) (EnvValue, error) {
	f, err := os.Open(file)
	if err != nil {
		return EnvValue{Value: "", NeedRemove: false}, err
	}
	fileStat, err := f.Stat()
	if err != nil {
		return EnvValue{Value: "", NeedRemove: false}, err
	}
	if fileStat.Size() == 0 {
		return EnvValue{Value: "", NeedRemove: true}, nil
	}
	br := bufio.NewReader(f)
	line, _ := br.ReadBytes('\n')
	line = bytes.ReplaceAll(line, []byte{0}, []byte{'\n'})
	line = bytes.TrimRight(line, " \n\t")
	if err := f.Close(); err != nil {
		return EnvValue{Value: "", NeedRemove: false}, err
	}
	return EnvValue{Value: string(line), NeedRemove: false}, nil
}

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	env := make(Environment)
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		tmpEnvValue, err := ReadFile(path.Join(dir, entry.Name()))
		if err != nil {
			return nil, err
		}
		env[entry.Name()] = tmpEnvValue
	}

	return env, nil
}
