## log, metric, trace 区别    
- Logging - 全量采集、时效性低、数据完整性要求高、结构化数据 or 非结构化；持久化到数据仓库、离线 hive or spark 任务、支撑报表 和 数据挖掘等上层应用；同时，数据完备性 可用于解决 用户反馈的 case 追查、冷备数据存储 (eg. 全量关系链、user-profile 等数据表的扫描上报)
- Metrics - 监控打点、指标聚合上报、时效性高、用于异常报警 和 监控分析。
- Tracing - 抽样采集、结构化数据、记录调用依赖 和 时间线 (以及扩展的 tag 信息)；常用于 全链路调用展示、服务理解、可复现问题的根源定位、测试分析、性能分析、链路优化 等


## git
目前常用的操作是
```
    git clone xx    直接基于master分支开发
    git commit 
    git push origin master:dev_zhg
    git pull       其他人有新的commit之后, 会自动merge,但会显得比较乱,推进按git pull --rebase
    git push xx
    
```

可以优化为
```
    git clone xx    直接基于master分支开发
    git checkout -b dev_zhg 创建一个自己的分支
    git rebase master   合并新的修改到本地自己的分支,rebase会使得开发过程更整洁
    git push origin dev_zhg   提交到仓库,并创建mr, review通过后对commits进行合并,将多个commit合并为一个,使得提交记录更清晰
    git rebase -i commitID
    git push origin dev_zhg
    
```

## zsh
zsh插件更丰富

zsh历史命令自动提示的安装, bash没对应的插件,不支持
```
cd ~/.oh-my-zsh/custom/plugins/
git clone https://github.com/zsh-users/zsh-autosuggestions   # 这个仓库还有其他插件,根据需要使用
vi ~/.zshrc 文件,将plugins改为plugins=(git zsh-autosuggestions), 主题根据需要选择,可以使用robbyrussel
vi ~/.oh-my-zsh/themes/robbyrussell.zsh-theme   改为  PROMPT='${ret_status} %{$fg[magenta]%}%d%{$reset_color%} $(git_prompt_info)'
source ~/.zshrc
```

oh-my-zsh 的 Zsh 扩展工具,是用于配置zsh的，需要先安装zsh,源自GitHub 的一个个人开源项目，它内置丰富的自定义主题（指 Zsh 显示风格），自带上百个功能各异的插件，比如显示git分支.
参考: https://martinguo.github.io/blog/2016/06/06/Your-Zsh/
robbyrussel主题有一个很令人难受的痛点，在于它并不能显示全路径，只能显示当前所在的文件夹名，直接修改.zsh-theme文件, 



ubuntu 配置
```
sudo apt-get install zsh 
chsh -s /bin/zsh   # 默认shell切换为zsh,如果提示sent invalidate(passwd) request则执行 sudo apt-get remove unscd
重启服务器   echo $SHELL为zsh则生效

```


## 存储
redis存在大key或热key问题,通过对value拆分,可以解决大key的存储问题,但仍然存在读写放大的问题, 就是要取其中的少数据数据仍然要把所有数据都取出来.
若用set类型, 过大也会导致redis卡顿.

分布式存储分为块存储,文件存储,对象存储,主要是所提供的接口的区别,另外文件存储和块存储的使用上没有区别,但块存储不能共享,硬盘只能挂载在一个主机上,文件系统是可以共享,路径可以被
多个人同时挂载使用. 块存储实际也能被共享挂载,但不支持同时写.  都属于网络存储(nas)，因为都不是本地直接访问，要通过网络协议(http, nfs)等远程读取.
- 对象存储，可以理解为键值存储，对外提供HTTP接口,支持的操作是GET、PUT、DEL
- 块存储，接口通常以QEMU Driver或者Kernel Module的方式存在，这种接口需要实现Linux的Block Device的接口或者QEMU提供的Block Driver接口，具有代表性的系统有阿里云的盘古系统，还有Ceph的RBD。
- 文件存储，支持POSIX接口，它跟传统的文件系统如Ext4是一个类型的，但区别在于分布式存储提供了并行化的能力，具有代表性的系统有Ceph的CephFS。


hbase的使用,利用其大宽表的优势,每列为一天的数据,存储n天.raw key需要做散列,避免数据聚集.


一般说的在线存储，指的是对延迟，可用性要求相对较高的一个存储系统，对读写的p99一般都在10ms内。比如推荐，消重的存储服务。这类服务对存储的性能要求很高，存储有一点波动，对业务就会有很大的影响。所以对存储的稳定性，可靠性，安全性都有很大的要求。 

## 向量相似
召回: 候选集是n*m(n个,m为特征长度),输入是为长度为m的向量(用户特征),取topk.可以先离线根据聚类算法将文章分类,减少候选量,然后再匹配.
消重: 精确和非精确,非精确部分可以用布隆过滤器.

同城多机房一般指的IDC之间的延迟在2ms以内,主备在不同的机房,主备都提供写,但是备机房的写会route到主机房,不需要其他数据同步策略,并且延迟短不影响体验.
异地多活指的是IDC的传输延迟在10ms以上,跨州的机房延迟更大,所以异地多活采用的方案是各自写到各自的机房,通过数据binlog同步等方式进行异步同步,
,需要互相同步,进行数据融合,因为不多活的话会影响体验. 需要解决数据同步过程的数据顺序性和数据冲突,数据顺序性需要一个统一的时间轴,数据冲突跟业务有关.

一般应用是双机房,但是一来的zk可以是三机房,使得在一个机房挂掉的时候zk能够正常提供服务. 切主有自动切和手动切两种,自动切主是通过watch zk即通过zk提供的选主服务.
切主有正常切主比较容灾演练,机房调整, 也有异常切主,比如机房断电.  切主的时候服务的各个角色都要执行相应的动作,存在失败的可能,需要有应对措施,甚至要考虑zk挂掉的场景.

## 监控
metric监控主要是store, counter, timer三类,分别对应的场景是内存使用量,qps,接口延时.对于counter类型,存储的是累计值,但实际关注的是增量.
存储累计值只是手段,计算增量进而计算变化率才是目的. 每个实例客户端将数据每隔一段时间flush到服务端,服务端每隔一个采样周期进行上报.客户端发送的是增量,
服务端会保存累计值一段时间,一般是几个小时,然后加上增量数据发送到kafka,再进行最终的处理. 时序数据库opentsdb存的也是累计值, 比如每30s有100个点,那么存储就是
0:0, 30:100,60:200,90:300. counter只能用于计算增量数据,不能用于计算累加量, 因为会因为服务升级,agent重启导致数据中断,取的累加值没有意义.而且累加值不可能保存下来,
数据太大肯定会超限.   那么对于服务升级这种场景,计算出来的qps曲线为啥又没出现明显的断层, 是因为有一个trick就是去掉导数的负值. 所以即使累计值下降了,在这短时间之类的数据
有点丢失也不影响整体趋势. 但理论上间隔变大,但看曲线也没变大, 那么更深层次的处理是只计算新起的pod和尚未升级的pod打点变化,而已经kill掉的pod的数据并不计算在内,数据存储的时候
是区分host的.qps计算的准确度跟采样周期有关,30s是够了的. 因为一些服务短期的流量高峰至少会持续这么久,如果持续的再短对于系统的影响没那么大. 


熔断是调用方探测到下游响应慢,超时等主动发起的不对下游调用,自动触发. 可能有多个调用方,但只有部分调用方会熔断.
降级被调用方服务降级,可以是全部降级也可以是部分降级,可以是自动也可以是手动,降级后会返回一个mock的假数据.比如可以是读接口全部降级, 只支持写.
可以是丢弃部分请求,或者部分请求返回mock数据.   但调用方主动对请求比例丢弃也属于降级范畴,而非熔断.
限流是qps或连接数达到了限制,需要开启过载保护,虽然负载过高,但系统响应不一定慢,不一定会触发熔断.
这几个服务治理措施可以在整体的service mesh实现,也可以是rpc框架做,降级和限流还可以是接入层网关做.
service mesh对rpc,http, mysql等协议都是支持的,也要能支持物理机的服务.

降级也还有跟业务紧密相关的策略,比如请求返回的文章数, 如果返回100条服务压力太大, 可以降级为只返回10条.

## go
在go里面，panic的问题往往比较棘手，协程的panic会导致整个服务进程挂掉.在main函数里面开启协程，而协程panic的话，会导致进程挂掉。main函数的recover实际不起作用.
