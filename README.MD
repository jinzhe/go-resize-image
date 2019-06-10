# 这个是用来做什么的？

批量修改照片大小的命令行工具，只会处理jpg格式。

# 程序功能

- 支持批量修改照片宽度
- 支持深层多级文件夹
- 支持根据图片原始宽度百分比压缩
- 支持生成缩图
- 支持加入水印

# 命令行参数

- -p 要处理的文件夹，默认"./data/"，结尾必须以`/`结尾。
- -w 宽度可以写整数或者百分比，如：`1920` 或 `80%`。
- -q jpeg 文件压缩质量（1-100），如：`80`，默认`80`。
- -t 额外生成缩图，如：`200x200`。
- -m 水印文件名，如：`logo.png`,不填写默认。

```
./main -p ./data/ -w 1920 -q 80 -t 200x200 -m test.png  
```