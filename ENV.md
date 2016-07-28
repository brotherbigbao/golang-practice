##环境变量配置
> ~/.profile
```
export GOPATH=/home/liuyibao/Golang
if [ -d "$GOPATH/bin" ] ; then
    PATH="$GOPATH/bin:$PATH"
fi
```
