# draw

Golang 实现的带权重的抽奖算法。  

安装：

```sh
go get github.com/duiying/draw
```

使用：

```go
a := map[int]int{
    1: 1, // 1 的权重占 10%
    2: 3, // 2 的权重占 30%
    3: 6, // 3 的权重占 60%,
}
k := draw.WeightRand(a)
```






