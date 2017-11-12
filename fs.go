package fs

import (
  "os"
  "io"
  "io/ioutil"
  "time"
  "math/rand"
)

/**
create a dir
 */
func Mkdir(dir string) (err error) {
  err = os.Mkdir(dir, os.ModePerm)
  return
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
ensure the file exist
 */
func EnsureFile(filepath string) (err error) {
  var (
    file *os.File
  )
  if _, err = os.Stat(filepath); os.IsNotExist(err) {
    file, err = os.Create(filepath)
    defer func() {
      file.Close()
    }()
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

  defer func() {
    srcFile.Close()
    targetFile.Close()
  }()

  return io.Copy(targetFile, srcFile)
}

/*
move a file
 */
func Move(src string, target string) (err error) {
  return os.Rename(src, target)
}

/**
remove a file or a dir
 */
func Remove(name string) (err error) {
  return os.RemoveAll(name)
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

// =========== permission ===========

/**
change the file permission
 */
func Chmod(filename string, mode os.FileMode) error {
  return os.Chmod(filename, mode)
}

func Lchmod(path string, uid int, gid int) (err error) {
  err = os.Lchown(path, uid, gid)
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
var seededRand *rand.Rand = rand.New(
  rand.NewSource(time.Now().UnixNano()))

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
func Mkdtemp(prefix string) (err error) {
  _, err = ioutil.TempDir(prefix+randomStr(6), prefix)
  return
}
