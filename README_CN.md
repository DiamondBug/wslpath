# wslpath
方便的在wsl中转换windows和wsl路径

### 安装
下载 ```wslpath``` 并放到 ```/usr/bin``` 下即可直接使用

### 如何使用
```shell
wslpath -p "C:\\User\peter"
wslpath -p /mnt/C/User/peter
```
输出:
```shell
/mnt/C/User/peter
C:\\User\peter
```
(It doesn't actually contain \n)

### 技巧
```
some/command $(wslpath -p "C:\\User\peter")
```
