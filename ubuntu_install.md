## matebook13 2020安装Ubuntu双系统
matebook13 2020因硬件的问题,安装Ubuntu20之前的版本会因为内核和驱动问题使得WiFi,触摸板使用不了,只能安装Ubuntu20.

安装步骤
1. 修改电脑配置, 关闭"极速启动",这样才能进入bios,不然开机太快无法配置. 
2. 电脑硬盘是uefi模式,所以进BIOS关闭security boot即安全启动.
3. 制作U盘启动盘,用utraliso等工具将镜像写入到U盘.
4. 进入到安装阶段, 安装过程中选择"something else"安装方式,这样需要自动分区,不要选择第一个的与win7共存,有可能破坏已有的分区.
如果是重装或者升级Ubuntu,第一个选项会是reinstall.
5. 分区模式,分三个分区, 一个swap分区,与内存大小相同,逻辑分区, 一个是efi分区,1-2g,用于启动,逻辑分区,因为是uefi模式,不再需要以前的boot分区,最后一个是/,主分区.
