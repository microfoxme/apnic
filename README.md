apnic
=====

apnic: IP Region Querier For Go Language


O(logN)的IP地域查询库程序。

1. 可以自动下载最新的IP分配国家列表；
2. 自动生成Go、JSON源代码，并为优化查询速度排序；
3. apnic.InPRC，apnic.IPInPRC用于查询某个IP是否属于大陆地区；

依赖包：
gawk  (sudo apt-get install gawk)
Golang (Go语言环境)
make

使用参见：
README

其他用途：

 $ make ipnet-cn.json

生成JSON格式的预排序文件。
