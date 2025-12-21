package v0

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

type Package struct {
	Destination string        `json:"description"`
	Extract     bool          `json:"extract"`
	Filename    string        `json:"filename"`
	Install     DLangDateTime `json:"install"`
	Installer   bool          `json:"installer"`
	Rename      string        `json:"rename"`
	Repository  string        `json:"repository"`
}

type DLangDateTime struct {
	time.Time
}

var DLangDateTimeLayouts = []string{
	"2006-01-02T15:04:05",
}

func (t *DLangDateTime) UnmarshalJSON(b []byte) error {
	// allow null
	if bytes.Equal(b, []byte("null")) {
		t.Time = time.Time{}
		return nil
	}

	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return fmt.Errorf("datetime must be string: %w", err)
	}

	s = strings.TrimSpace(s)
	if s == "" {
		t.Time = time.Time{}
		return nil
	}

	var lastErr error
	for _, layout := range DLangDateTimeLayouts {
		tt, err := time.Parse(layout, s)
		if err == nil {
			t.Time = tt
			return nil
		}

		lastErr = err
	}

	return fmt.Errorf("invalid datetime value %q: %v", s, lastErr)
}

func LoadJson(path string) (map[string]Package, any) {
	_, err := os.Stat(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", fmt.Sprintf("File %s does not exist", path))
		return nil, err
	}

	jsonData, err := os.ReadFile(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", fmt.Sprintf("Error has occurred in read: %s", err.Error()))
		return nil, err
	}

	// fmt.Println(string(jsonData))

	var packages map[string]Package
	err = json.Unmarshal([]byte(jsonData), &packages)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", fmt.Sprintf("Error has occurred in json unmarshal: %s", err.Error()))
		return nil, err
	}

	return packages, nil
}
