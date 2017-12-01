package fs

import (
  "os"
  "io/ioutil"
  "path"
  "io"
)

/**
copy a file or dir
 */
func Copy(src string, target string) (err error) {

  var (
    srcFile    *os.File
    targetFile *os.File
    fileInfo   os.FileInfo
    files      []os.FileInfo
  )

  if fileInfo, err = os.Stat(src); err != nil {
    return
  }

  if fileInfo.IsDir() {
    // read dir and copy one by one
    files, err = ioutil.ReadDir(src)
    if err != nil {
      return
    }

    if err = EnsureDir(target); err != nil {
      return
    }

    for _, file := range files {
      filename := file.Name()
      src = path.Join(src, filename)
      target = path.Join(target, filename)
      if err = Copy(src, target); err != nil {
        return err
      }
    }

  } else {
    srcFile, err = os.Open(src)

    if err != nil {
      return
    }

    targetFile, err = os.Create(target)

    defer func() {
      srcFile.Close()
      targetFile.Close()
    }()

    _, err = io.Copy(targetFile, srcFile)
  }
  return
}
