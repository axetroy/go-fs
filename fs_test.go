package fs

import (
  "testing"
  "path"
  "os"
)

const testDir string = ".temp"

func init() {
  err := EnsureDir(testDir)
  if err != nil {
    panic(err)
  }
}

func Test_CreateDir(t *testing.T) {
  var (
    dirPath = path.Join(testDir, "mkdir_test")
  )

  Mkdir(dirPath)

  defer func() {
    Remove(dirPath)
  }()

  isExist := PathExists(dirPath)

  if isExist == false {
    t.Error("Test Mkdir fail...")
  }

  t.Log("Test Mkdir success...")
}

func Test_EnsureFile(t *testing.T) {
  var (
    newFilePath = path.Join(testDir, "test_ensure.md")
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
  EnsureDir(newDirPath)

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
    filePath    = path.Join(testDir, "test_copy.md")
    newFilePath = path.Join(testDir, "test_copy_new.md")
  )
  EnsureFile(filePath)

  // remove all test file
  defer func() {
    if err := Remove(filePath); err != nil {
      panic(err)
    }
    if err := Remove(newFilePath); err != nil {
      panic(err)
    }
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
    filePath    = path.Join(testDir, "test_move.md")
    newFilePath = path.Join(testDir, "test_mode_new.md")
  )
  EnsureFile(filePath)

  // remove dist file
  defer func() {
    Remove(newFilePath)
    Remove(filePath)
  }()

  if err := Move(filePath, newFilePath); err != nil {
    panic(err)
    t.Error("Move file fail...")
    return
  }

  // old file should not exist
  if isExist := PathExists(filePath); isExist == true {
    t.Error("Move file Fail.")
    return
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
    newFilePath = path.Join(testDir, "test_chmod.md")
  )
  EnsureFile(newFilePath)

  defer func() {
    Remove(newFilePath)
  }()

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
