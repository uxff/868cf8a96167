# gohive

使用golang链接hiveserver2的库。

对应hive版本是2.1。

go通过hive的thrift rpc接口链接hiveserver2。

实现了同步查询，异步查询等接口，并简化调用，优化性能，客户端不用处理复杂的thrift接口传输对象。

go get -insecure "git.dev.acewill.net/rpc/Gohive"

当前使用thrift-0.9.3版本，与hive的thrift对应。

如果需要更新，则删除inf，将inf.ctx重命名为inf，进行go build会报错，将对应的thrift接口第一个参数增加context.Background(), 即可。

