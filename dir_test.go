package fs

import (
  "testing"
  "path"
  "os"
)

const TestDir = ".temp"

func init() {
  err := EnsureDir(TestDir)
  if err != nil {
    panic(err)
  }
}

/**
Test mkdir
 */
func Test_MkDir(t *testing.T) {
  var (
    dirPath = path.Join(TestDir, "mkdir_test_dir")
    err     error
  )

  if err = Mkdir(dirPath); err != nil {
    t.Errorf("Create mkdir fial")
    return
  }

  defer func() {
    Remove(dirPath)
  }()

  if _, err := os.Stat(dirPath); os.IsNotExist(err) {
    t.Errorf("Test Mkdir fail")
    return
  }
}

/**
Ensure a dir
 */
func Test_EnsureDir(t *testing.T) {
  var (
    dirPath = path.Join(TestDir, "ensure_test_dir")
    err     error
  )

  if err = EnsureDir(dirPath); err != nil {
    t.Errorf("ensure dir fail")
  }

  defer func() {
    Remove(dirPath)
  }()

  if _, err := os.Stat(dirPath); os.IsNotExist(err) {
    t.Errorf("ensure dir fail")
    return
  }
}

/**
Read dir
 */
func Test_ReadDir(t *testing.T) {
  var (
    distDir = path.Join(".github")
  )

  if files, err := Readdir(distDir); err != nil {
    t.Error("Readdir Fail.")
  } else {
    if len(files) != 4 {
      t.Error("Readdir Fail.")
      return
    }
  }
}
