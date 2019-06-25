Leaf
====
这是一个[name5566/leaf](https://github.com/name5566/leaf) 的并发版本。
leaf是一款非常简单、稳定的游戏框架，拥有良好的封装，非常适合快速搭建游戏服务端程序。
我在实际项目中使用leaf，并针对我们的情况做了一些修改。

Features
---------

* 网络请求并发路由
* 集成了[natefinch/lumberjack](https://github.com/natefinch/lumberjack)日志库
* 暴露TCP超时接口，方便检测心跳
* 把userData和一些计数器改成原子操作，保证并发安全性
* Msg处理函数列表从list改成map，从而可以使用非连续msgID

请求并发路由
---------

原leaf为了降低网络编程的复杂性，把网络事件通过chan做了串行处理，并配合leaf/go模块实现并发。
对于希望高并发的游戏、并且msg较多的情况下，在每个handler中嵌入go略显繁琐，尤其是对DB的操作比较耗时，串行路由可能因此整体事件积压，导致高延迟。
因此我们从框架层去掉chan，直接在每个conn读数据的goroutine来进行事件处理。这样能保证单个conn的事件是串行处理的，而且耗时操作不会影响服务器对其他conn的响应。

Licensing
---------

Leaf is licensed under the Apache License, Version 2.0. See [LICENSE](https://github.com/name5566/leaf/blob/master/LICENSE) for the full license text.
