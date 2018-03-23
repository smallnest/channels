# Channel 应用模式



## Lock/TryLock模式

### 最高效的TryLock

[trylock]()

### 使用Channel实现TryLock

[trylock_channel]()

### 使用Channel实现Timeout功能的TryLock

[trylock_timeout]()

## or 信号模式
从多个channel读取一个信号, 一旦读取到一个信号，则不再读取。

比如向多个服务器发送相同的http request,每个请求的结果放在单独的一个channel中， 只要其中一个服务器返回结果，则其它请求就被忽略。

### or channel by  goroutine

[or_channel_go]()

最简单的方式就是为每个channel启动一个goroutine, 每个goroutine读取自己负责的channel，一旦读取到一个信号，就关闭返回的channel。
显然，为每个channel启动一个goroutine太浪费了，虽然goroutine是一种轻量级的实现，但是如果数量巨大的情况下也会导致资源的大量占用以及调度上的性能低下。


### or channel (递归)

[or_channel]()
基于递归的方式实现， 使用依次递归的方式

### or channel (递归)

[or_channel_rec]()
基于递归的方式实现, 使用分而治之的方式

### or channel reflect

[or_channel_rec]()
基于反射的方式

## or_done_channel模式

与上面的or 信号模式不同， `or done channel`模式是从一个channel中读取数据，只有当channel被关闭，或者done 信号channel被关闭的时候，读取操作才退出。

[or_done_channel]()

如果将`done`这个信号channel换成 `context`,则可以依靠 `context.WithCancel` 来cancel读取，和这个模式类似。

## flat 模式

[flat]()
将多个channels平展成一个channels。 与`Fan In`不同的是，输入的channels是从一个channel中读取出来的，而`Fan In`模式中的channels是一个channel slice。

## map/reduce 模式

[map]()
[reduce]()


## Fan In 扇入模式

将多个channel合并成一个channel

## Fan Out 扇出模式

将一个Channel分成多个Channel。 有两种情况， 一种是每个channel都包含同样的数据(复制模式)， 另一种将原数据均匀分布到各输出channel中(分布模式)

## Tee Channel

类似linux的tee命令，是`Fan Out`的一种特例



## References
1. https://github.com/kat-co/concurrency-in-go-src
2. https://github.com/campoy/justforfunc/tree/master/27-merging-chans
3. https://github.com/eapache/channels
4. https://github.com/LK4D4/trylock