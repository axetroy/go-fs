package fs

import (
  "os"
  "io"
  "io/ioutil"
)

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
ensure the file exist
 */
func EnsureFile(filepath string) (err error) {
  if _, err = os.Stat(filepath); os.IsNotExist(err) {
    _, err = os.Create(filepath)
  }
  return
}

/**
check a path is exist or not
 */
func PathExists(path string) (isExist bool) {
  if _, err := os.Stat(path); os.IsNotExist(err) {
    return false
  }
  return true
}

/**
stat a file
 */
func Stat(name string) (info os.FileInfo, err error) {
  return os.Stat(name)
}

/**
stat a file
 */
func LStat(name string) (info os.FileInfo, err error) {
  return os.Lstat(name)
}

/**
copy a file
 */
func Copy(src string, target string) (written int64, err error) {
  var (
    srcFile    *os.File
    targetFile *os.File
  )
  srcFile, err = os.Open(src)

  if err != nil {
    return
  }

  targetFile, err = os.Create(target)

  return io.Copy(targetFile, srcFile)
}

/*
move a file
 */
func Move(src string, target string) (err error) {
  return os.Rename(src, target)
}

/**
remove a file
 */
func Remove(name string) (err error) {
  return os.Remove(name)
}

/*
write a file
 */
func WriteFile(filename string, data []byte, perm os.FileMode) error {
  return ioutil.WriteFile(filename, data, perm)
}

/**
read a file
 */
func ReadFile(filename string) ([]byte, error) {
  return ioutil.ReadFile(filename)
}

/**
change the file permission
 */
func Chmod(filename string, mode os.FileMode) error {
  return os.Chmod(filename, mode)
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
