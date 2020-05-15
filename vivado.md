## vivado
xilinx的fpga开发ide最开始是ise,但因性能和支持的fpga的限制,后面推出了vivado. ise最后的版本是14.1,在官网可以下载.

Vivado是Xilinx公司于2012年推出的新一代集成设计环境。推出Vivado是为了提高设计者的效率，他能显著增加Xilinx的28nm工艺的可编程逻辑期间的设计、综合于实现效率。
即随着FPGA进入28nm时代，ISE工具有些“不合时宜”了，硬件提升了，软件也需要提升。当下已经是2019.2版本,分为hlx版本,lab等版本, hlx需要license.
各个版本的比较https://china.xilinx.com/products/design-tools/vivado/vivado-webpack.html.


Vivado HLS是Xilinx推出的高层次综合工具，可以实现直接使用 C，C++ 以及 System C 语言规范对赛灵思可编程器件进行编程，无需手动创建 RTL，从而可加速 IP 创建。参考:https://www.uisrc.com/portal.php?mod=view&aid=102


## 计算机组成
zwj: 1s钟能够执行多少条指令? 
因为缓存特性: 
(1)数组顺序访问为什么会比链表访问要快
(2)遍历一个数组,按行遍历和按列遍历的时间是不同的. 当数据量比较大时,按行遍历是按列遍历的25倍.(数据来源CSAPP第二版) 