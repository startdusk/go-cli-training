
初始化
https://github.com/spf13/cobra-cli

```bash
$ mkdir audiofile-cli 
$ cd audiofile-cli
$ go mod init audiofile-cli
# Initialize your Cobra CLI:
$ cobra-cli init
```

添加一个命令

```bash
$ cobra-cli add upload
```
会在cmd目录下添加一个```upload```的文件
