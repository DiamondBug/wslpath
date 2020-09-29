# wslpath
Convert Windows and WSL paths

### Install
Just download executable file ```wslpath``` and put into /usr/bin

### How to use
```shell
wslpath -p C:\\User\peter
wslpath -p /mnt/C/User/peter
```
Output:
```shell
/mnt/C/User/peter
C:\\User\peter
```
(It doesn't actually contain \n)
