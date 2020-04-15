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


## 向量相似
召回算法: 根据用户向量匹配候选集的向量,搜索出相似度大的
消重: 根据文章内容聚类
