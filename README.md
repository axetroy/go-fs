## File Accessor For go-lang Inspired by NodeJs fs system

[![Build Status](https://travis-ci.org/axetroy/go-fs.svg?branch=master)](https://travis-ci.org/axetroy/go-fs)
![License](https://img.shields.io/badge/license-Apache-green.svg)
![Size](https://github-size-badge.herokuapp.com/gpmer/gpm.js.svg)

API almost like NodeJs [fs module](http://nodejs.cn/api/fs.html)

## Usage

```bash
go get https://github.com/axetroy/go-fs.git
```

```go

import fs "github.com/axetroy/go-fs"

func main(){
  if err := fs.EnsureFile("./testFile.txt");err !=nill {
    
  }
  if err := fs.Copy("./testFile.txt", "./newTestFile.txt");err !=nil {
  
  }
}

```

## Contributing

```bash
go get https://github.com/axetroy/go-fs.git
cd $GOPATH/src/github.com/axetroy/go-fs
go test -v
```

[Contributing Guid](https://github.com/axetroy/Github/blob/master/CONTRIBUTING.md)

## Test

```bash
go test -v
```

## Contributors

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
| [<img src="https://avatars1.githubusercontent.com/u/9758711?v=3" width="100px;"/><br /><sub>Axetroy</sub>](http://axetroy.github.io)<br />[💻](https://github.com/axetroyanti-redirect/go-fs/commits?author=axetroy) [🐛](https://github.com/axetroy/go-fs/issues?q=author%3Aaxetroy) 🎨 |
| :---: |
<!-- ALL-CONTRIBUTORS-LIST:END -->

## License

[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Faxetroy%2Fgo-fs.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Faxetroy%2Fgo-fs?ref=badge_large)