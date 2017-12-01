package fs

import (
  "os"
  "io/ioutil"
  "time"
  "math/rand"
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

// generate random string
var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

func stringWithCharset(length int, charset string) string {
  b := make([]byte, length)
  for i := range b {
    b[i] = charset[seededRand.Intn(len(charset))]
  }
  return string(b)
}

func randomStr(length int) string {
  var (
    lower string = "abcdefghijklmnopqrstuvwxyz"
    upper string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
    num   string = "0123456789"
  )
  return stringWithCharset(length, lower+upper+num)
}

/**
create random temp dir
 */
func MkTemp(prefix string) (err error) {
  _, err = ioutil.TempDir(prefix+randomStr(6), prefix)
  return
}
