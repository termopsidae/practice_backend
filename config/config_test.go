package config

import (
	"fmt"
	"testing"
)

func TestSetConfig(t *testing.T) {
	if FilePath == "" {
		GetAbsolutePath()
	}
	err := SetPKSConfig(FilePath, "testOPKS", "testAPKS")
	if err != nil {
		fmt.Println(err.Error())
	}
}
func TestReadConfig(t *testing.T) {
	a := ConfigPKS("OPK_SECRET")
	fmt.Println(a)
	fmt.Println(ConfigPKS("APK_SECRET"))
}

func TestEditConfig(t *testing.T) {
	if FilePath == "" {
		GetAbsolutePath()
	}
	err := EditPKSConfig(FilePath, "OPK_SECRET", "editOPKS")
	if err != nil {
		fmt.Println(err.Error())
	}
}
