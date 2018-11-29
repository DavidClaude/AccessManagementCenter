package codex

import (
	"github.com/go-ini/ini"
	"fmt"
	"errors"
	"strconv"
)

type TomlConfig struct {
	file   *ini.File
	tables map[string]map[string]string
}

func (tc *TomlConfig) Init(path string) (err error) {
	tc.file, err = ini.Load(path)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	tc.tables = make(map[string]map[string]string)
	return nil
}

func (tc *TomlConfig) Fill(sec string) (err error) {
	s, err := tc.file.GetSection(sec)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	table := s.KeysHash()
	tc.tables[sec] = table
	return nil
}

func (tc *TomlConfig) GetString(sec, key string) (val string, err error) {
	table, ok := tc.tables[sec]
	if !ok {
		return "", errors.New("no section")
	}
	v, ok := table[key]
	if !ok {
		return "", errors.New("no key")
	}
	return v, nil
}

func (tc *TomlConfig) GetInt(sec, key string) (val int, err error) {
	table, ok := tc.tables[sec]
	if !ok {
		return 0, errors.New("no section")
	}
	v, ok := table[key]
	if !ok {
		return 0, errors.New("no key")
	}
	val, err = strconv.Atoi(v)
	if err != nil {
		return 0, err
	}
	return val, nil
}

func (tc *TomlConfig) GetInt64(sec, key string) (val int64, err error) {
	table, ok := tc.tables[sec]
	if !ok {
		return 0, errors.New("no section")
	}
	v, ok := table[key]
	if !ok {
		return 0, errors.New("no key")
	}
	val, err = strconv.ParseInt(v, 10, 64)
	if err != nil {
		return 0, err
	}
	return val, nil
}

func (tc *TomlConfig) GetFloat64(sec, key string) (val float64, err error) {
	table, ok := tc.tables[sec]
	if !ok {
		return 0, errors.New("no section")
	}
	v, ok := table[key]
	if !ok {
		return 0, errors.New("no key")
	}
	val, err = strconv.ParseFloat(v, 64)
	if err != nil {
		return 0, err
	}
	return val, nil
}

func (tc *TomlConfig) GetBool(sec, key string) (val bool, err error) {
	table, ok := tc.tables[sec]
	if !ok {
		return false, errors.New("no section")
	}
	v, ok := table[key]
	if !ok {
		return false, errors.New("no key")
	}
	if v == "true" || v == "1" {
		return true, nil
	}
	return false, nil
}
