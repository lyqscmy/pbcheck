# NAME

pbcheck - lint pb.go unsafe useage

Table of Contents
=================

* [NAME](#name)
* [DESCRIPTION](#description)
* [PREREQUISITES](#prerequisites)
* [INSTALLATION](#installation)
* [AUTHORS](#authors)

# DESCRIPTION

proto-gen-go生成的file.pb.go中的结构体每个字段都有Getter方法，方法处理了receiver为nil的情况，所以是推荐使用。但是在编码时程序员可能图方便直接foo.bar使用字段。

pbcheck检查在读的场景没有使用Getter方法，而是直接读指针的用法。

# PREREQUISITES

TODO

[Back to TOC](#table-of-contents)

# INSTALLATION
    make
    sudo make install

[Back to TOC](#table-of-contents)

# AUTHORS

- Li Yongqiang (李勇強) `<lyqscmy@gmail.com>`

[Back to TOC](#table-of-contents)

