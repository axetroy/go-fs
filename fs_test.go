package fs

import (
  "testing"
  "fmt"
  "path"
  "os"
)

const testDir string = "./.temp"

func init() {
  err := EnsureDir(testDir)
  if err != nil {
    fmt.Println(err)
  }
}

func Test_EnsureFile(t *testing.T) {
  var (
    newFilePath = path.Join(testDir, "testfile.md")
  )
  EnsureFile(newFilePath)

  isExist := PathExists(newFilePath)

  defer func() {
    Remove(newFilePath)
  }()

  if isExist == true {
    t.Log("Ensure File Success...")
  } else {
    t.Error("Ensure File Fail.")
  }
}

func Test_EnsureDir(t *testing.T) {
  var (
    newDirPath = path.Join(testDir, "testdir")
  )
  EnsureFile(newDirPath)

  defer func() {
    Remove(newDirPath)
  }()

  isExist := PathExists(newDirPath)

  if isExist == true {
    t.Log("Ensure dir Success...")
  } else {
    t.Error("Ensure dir Fail.")
  }
}

func Test_Copy(t *testing.T) {
  var (
    filePath    = path.Join(testDir, "testFile.md")
    newFilePath = path.Join(testDir, "testFile-new.md")
  )
  EnsureFile(filePath)

  // remove all test file
  defer func() {
    Remove(filePath)
    Remove(newFilePath)
  }()

  Copy(filePath, newFilePath)

  // old file should exist
  if isExist := PathExists(filePath); isExist == false {
    t.Error("Copy file Fail.")
  }

  isExist := PathExists(newFilePath)

  if isExist == true {
    t.Log("Copy file Success...")
  } else {
    t.Error("Copy file Fail.")
  }
}

func Test_Move(t *testing.T) {
  var (
    filePath    = path.Join(testDir, "testFile.md")
    newFilePath = path.Join(testDir, "testFile-new.md")
  )
  EnsureFile(filePath)

  // remove dist file
  defer func() {
    Remove(newFilePath)
  }()

  Move(filePath, newFilePath)

  // old file should not exist
  if isExist := PathExists(filePath); isExist == true {
    t.Error("Move file Fail.")
  }

  isExist := PathExists(newFilePath)

  if isExist == true {
    t.Log("Move file Success...")
  } else {
    t.Error("Move file Fail.")
  }
}

func Test_Readdir(t *testing.T) {
  var (
    dir = path.Join(".github")
  )

  files, err := Readdir(dir)

  if err != nil {
    panic(err)
    t.Error("Readdir Fail.")
    return
  }

  if len(files) != 4 {
    t.Error("Readdir Fail.")
    return
  }

  t.Log("Readdir Success.")
}

func Test_Chmod(t *testing.T) {
  var (
    newFilePath = path.Join(testDir, "testfile.md")
  )
  EnsureFile(newFilePath)

  Chmod(newFilePath, 0666)

  stat, err := os.Stat(newFilePath)

  if err != nil {
    panic(err)
    t.Error(err)
  }

  if stat.Mode() != 0666 {
    t.Log("Chmod fail!")
  }

  t.Log("Chmod success!")

}
