## 项目结构
```shell
├───application # 服务目录
│   └───fileserver
│       ├───api
│       ├───model
│       └───rpc
├───common  # 公共模块
└───deploy  # 部署
    ├───compose  # compose 部署
    └───k8s      # k8s 部署
```
## 中间件部署
### mysql

k8s 集群部署方式，mysql 采用一主两从的方式部署，主节点向外暴露3306端口，从节点只读向外暴露3307端口。因此可以采用读写分离的模式，访问3306端口写，3307端口读。

#### 技术细节

mysql 清单，master_init.sh、slave_init.sh 脚本用于主节点的初始化，放在 /docker-entrypoint-initdb.d/ 目录下面，mysql 容器的 entrypoint 脚本执行完之后会执行这里面的脚本，而且容器只会在第一次初始化的时候执行。

### redis

redis 采用一主两从，三哨兵的方式部署，使用的时候通过哨兵的地址连接。通过 ip:16379 端口连接哨兵，哨兵会返回当前主节点的地址。
#### 技术细节
在这里有个细节就是，redis 三个 pod 用 30040、30041、30042 三个不同的端口向外暴露服务，这样做的好处是避免 NodePort 负载均衡机制将哨兵返回的主节点地址负载到从节点上。

### minio

### kafka