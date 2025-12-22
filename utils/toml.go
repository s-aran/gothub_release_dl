package utils

import (
	"errors"
	"fmt"
	"os"
	"sort"

	"github.com/pelletier/go-toml/v2"
)

var ErrVersionMismatch = errors.New("version mismatch")

func WriteToToml[T any](root map[string]any, main_data map[string]T, filename string) error {
	keys := make([]string, 0, len(main_data))
	for k := range main_data {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		root[k] = main_data[k]
	}

	b, err := toml.Marshal(root)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, b, 0o644)
}

type FromAnyMap[T any] func(map[string]any) (T, error)

func ReadFromToml[T any](filename string, parse FromAnyMap[T]) (uint, map[string]T, error) {
	b, err := os.ReadFile(filename)
	if err != nil {
		return 0, nil, err
	}

	var root map[string]any
	if err := toml.Unmarshal(b, &root); err != nil {
		return 0, nil, err
	}

	// get version
	maybe_version, ok := root["version"]
	if !ok {
		return 0, nil, fmt.Errorf("version not found")
	}
	maybe_int64_version, ok := maybe_version.(int64)
	if !ok {
		return 0, nil, fmt.Errorf("version is not int64")
	}
	version := uint(maybe_int64_version)

	items := make(map[string]T)
	for k, v := range root {
		if k == "version" {
			continue
		}

		sub, ok := v.(map[string]any)
		if !ok {
			return 0, nil, fmt.Errorf("[%s] must be table, got %T", k, v)
		}

		t, err := parse(sub)
		if err != nil {
			return 0, nil, fmt.Errorf("[%s] %w", k, err)
		}

		items[k] = t
	}

	return version, items, nil
}
