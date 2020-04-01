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
    git pull       其他人有新的commit之后
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