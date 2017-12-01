package fs

import (
  "os"
  "io/ioutil"
  "path"
  "fmt"
)

/*
move a file or dir
 */
func Move(src string, target string) (err error) {

  var (
    fileInfo os.FileInfo
    files    []os.FileInfo
  )

  if fileInfo, err = os.Stat(src); err != nil {
    return
  }

  if fileInfo.IsDir() {

    // read dir and move one by one
    files, err = ioutil.ReadDir(src)
    if err != nil {
      return
    }

    if err = EnsureDir(target); err != nil {
      return
    }

    for _, file := range files {
      filename := file.Name()
      srcFile := path.Join(src, filename)
      targetFile := path.Join(target, filename)
      if err = Move(srcFile, targetFile); err != nil {
        return err
      }
    }

    fmt.Println("remove", src)

    // copy all done, should remove the src dir
    err = os.RemoveAll(src)

    return
  } else {
    return os.Rename(src, target)
  }
}
