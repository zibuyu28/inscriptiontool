package wallet

import (
	"fmt"
	"os"
	"path/filepath"
)

func Delete(name string) {
	if name == "" {
		fmt.Println("name is empty")
		return
	}
	exist, err := checkWalletExist(name)
	if err != nil {
		fmt.Printf("check wallet exist failed: %v\n", err)
		return
	}
	if !exist {
		fmt.Printf("wallet %s not exist\n", name)
		return
	}

	err = os.Rename(filepath.Join(RootDir, fmt.Sprintf("%s.json", name)), filepath.Join(RootDir, fmt.Sprintf("%s.json.del", name)))
	if err != nil {
		fmt.Printf("delete wallet failed: %v\n", err)
		return
	}
}
