package fs

import (
  "os"
  "io/ioutil"
)

/**
create a dir
 */
func Mkdir(dir string) (err error) {
  return os.Mkdir(dir, os.ModePerm)
}

/**
ensure the dir exist
 */
func EnsureDir(dir string) (err error) {
  if _, err = os.Stat(dir); os.IsNotExist(err) {
    err = os.Mkdir(dir, os.ModePerm)
  }
  return
}

/**
read dir and get file list
 */
func Readdir(dir string) (files []string, err error) {
  var (
    fileInfos []os.FileInfo
  )
  fileInfos, err = ioutil.ReadDir(dir)

  for _, f := range fileInfos {
    files = append(files, f.Name())
  }
  return
}

/**
create random temp dir
 */
func Mktemp(dir string, prefix string) (string, error) {
  return ioutil.TempDir(dir, prefix)
}

func Rmdir(path string) (error) {
  return os.RemoveAll(path)
}
