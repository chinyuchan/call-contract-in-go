# Go语言调用智能合约
合约中pure和view方法，不需要改变状态，不需要签名发送交易，只需读全节点上的数据即可。要改变区块链状态的，才需签名发送交易。

## 视频教程
* BiliBili
  * [01-自己编码调用参数](https://www.bilibili.com/video/BV1XU4y1N7Hw/?spm_id_from=333.999.0.0&vd_source=79484a601afa1e7d36a00ef527669e7e)
  * [02-使用abi对象](https://www.bilibili.com/video/BV1RL41177Nv/?spm_id_from=333.999.0.0&vd_source=79484a601afa1e7d36a00ef527669e7e)
  * [03-使用abigen生成代码](https://www.bilibili.com/video/BV1mQ4y1e7Ju/?spm_id_from=333.999.0.0&vd_source=79484a601afa1e7d36a00ef527669e7e)
* YouTube
  * [01-自己编码调用参数](https://www.youtube.com/watch?v=Z9NWXeKxjQk&list=PL9aoThVN5PLn7_FiUoaqHnttwimueR-3F&index=1)
  * [02-使用abi对象](https://www.youtube.com/watch?v=w_aMws2nRwA&list=PL9aoThVN5PLn7_FiUoaqHnttwimueR-3F&index=2)
  * [03-使用abigen生成代码](https://www.youtube.com/watch?v=a9Pbn6u82Jo&list=PL9aoThVN5PLn7_FiUoaqHnttwimueR-3F&index=3)

## 流程
调用合约的本质是把函数名、参数编码后，放到交易的data里面，发送到节点处理。因此，主要工作就是如何进行编码。

## 代码
* **自己编码**  
  `call1`知道合约的方法名、参数列表，调用以太坊的库进行编码。
* **借助abi编码**
  * **使用abi对象**  
    `call2`读abi文件，构造出abi对象，进行编码。
  * **使用合约对象**  
    `cdd + main.go`借助abigen工具，生成合约对象，不需要编码（编码在内部帮你做了），直接调用方法。

## 总结
自己编码是最通用的，任何时候都能用。其他两种方式都要借助abi，如果没有abi（未开源），就不适用。在有abi的前提下，推荐使用第三种方式，最方便。

