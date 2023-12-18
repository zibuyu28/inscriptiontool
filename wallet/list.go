package wallet

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
	"path/filepath"
	"strings"
)

func List() error {
	_ = os.MkdirAll(RootDir, os.ModePerm)
	items, err := os.ReadDir(RootDir)
	if err != nil {
		return errors.Wrap(err, "read dir error")
	}
	for _, item := range items {
		if item.IsDir() || strings.HasSuffix(item.Name(), ".del") {
			continue
		}
		file, err := os.ReadFile(filepath.Join(RootDir, item.Name()))
		if err != nil {
			return errors.Wrap(err, "get wallet error")
		}
		fmt.Println(string(file))
	}
	return nil
}
