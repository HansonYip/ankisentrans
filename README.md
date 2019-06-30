# anki-sentence-trans

本人积累英语词汇的方式主要是通过使用 Kindle 阅读英语文章并记录生词，使用 Kindle Mate 将生词导出后，再倒入 Anki，其中，卡片的模板使用实用且漂亮的 [Leaflyer](http://leaflyer.lofter.com/post/4798b6_a4492e3)。

在使用过程中发现存在两个小问题：

1. 在Kindle Mate 导出的数据中，单词的音标和释义在同一个字段；
2. 缺少例句的翻译。

因此，本人通过编写此工具来解决上述两问题。



### 安装
```bash
$ go get -u github.com/HansonYip/ankisentrans
$ ankisentrans
A tool for sentence translation for Anki.

Usage:
  ankisentrans [command]

Available Commands:
  help        Help about any command
  translate   

Flags:
  -h, --help   help for ankisentrans

Use "ankisentrans [command] --help" for more information about a command.
```


### 使用
```bash
$ ankisentrans translate <src_file> <dest_file>
```


