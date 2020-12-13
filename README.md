

这是一个用来练习 InnoDB 表空间的简单命令，通过命令可以查看表空间里的页信息，目前支持查看`FSP_HDR`、`INODE`、`XDES`、`INDEX`、`IBUFBITMAP`和`DATA DICT`类型的页数据。

PS：灵感来自于《MySQL技术内幕-InnoDB存储引擎》

参数：

| Option | Parameters        | Description          |
| ------ | ----------------- | -------------------- |
| -f     | <filename>        | 指定表空间           |
| -p     | <page offset>     | 指定页码             |
| -w     | <output filename> | 将内容输出到指定文件 |

子命令：

| Command  | Description |
| -------   |  ----------- |
| all     | 简单列出所有页的类型和页码，并且根据页类型统计数量，，支持`-f`、`-w`参数。 |
| simple | 简单列出页面的页头和页尾信息，支持`-f`、`-p`参数。           |
| detail | 详细列出页面的页头、页体和页尾信息，支持`-f`、`-p`，`-w`参数。 |

