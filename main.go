package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"

	"gopkg.in/ini.v1"
)

func GetInI(section string, key string) (*ini.Key, error) {
	conf, err := ini.Load("conf.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		return nil, err
	}
	return conf.Section(section).GetKey(key)
}

func IsFileExist(filePath string) bool {
	_, err := os.Stat(filePath)
	return err == nil || os.IsExist(err)
}

func IsDir(FilePath string) bool {
	FileInfo, _ := os.Stat(FilePath)
	return FileInfo.IsDir()
}

func CopyFile(filePath string, distFilePath string) {
	copyFile, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Fail to Open file: %v", err)
		return
	}
	defer copyFile.Close()

	distFile, err := os.Create(distFilePath)
	if err != nil {
		fmt.Printf("Fail to Create file: %v", err)
		return
	}
	defer distFile.Close()

	io.Copy(distFile, copyFile)
}

func CopyDirectory(CopyDirectoryPath string, OutDirectoryPath string) {
	if !IsFileExist(OutDirectoryPath) {
		os.Mkdir(OutDirectoryPath, os.ModePerm)
	}

	Files, err := ioutil.ReadDir(CopyDirectoryPath)
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		return
	}

	for _, File := range Files {
		CopyFilePath := path.Join(CopyDirectoryPath, File.Name())
		OutCopyFilePath := path.Join(OutDirectoryPath, File.Name())
		if IsDir(CopyFilePath) {
			CopyDirectory(CopyFilePath, OutCopyFilePath)
			continue
		}
		CopyFile(CopyFilePath, OutCopyFilePath)
	}
}

func Copy(FileName []string) {
	CopyDirectoryPath, _ := GetInI("paths","copy_directory_path")
	OutDirectoryName, _ := GetInI("","out_directory_name")
	OutDirectoryPath, _ := GetInI("paths","out_directory_path")

	OutPath := path.Join(OutDirectoryPath.String(), OutDirectoryName.String())

	if !IsFileExist(OutPath) {
		os.Mkdir(OutPath, os.ModePerm)
	}

	for _, value := range FileName {
		FilePath := path.Join(CopyDirectoryPath.String(), value)
		OutCopyDirectoryPath := path.Join(OutPath, value)
		if !IsFileExist(FilePath) { continue }
		if IsDir(FilePath) {
			CopyDirectory(FilePath, OutCopyDirectoryPath)
			continue
		}
		CopyFile(FilePath, OutCopyDirectoryPath)
	}
}

func main() {
	FileName := []string{"locales", "swiftshader", "89.0.4335.0.manifest", "chrome.dll", "chrome.exe",
		"chrome_100_percent.pak", "chrome_200_percent.pak", "chrome_elf.dll", "headless_lib.pak", "icudtl.dat",
		"libEGL.dll", "libGLESv2.dll", "resources.pak", "snapshot_blob.bin", "v8_context_snapshot.bin"}
	Copy(FileName)
}
