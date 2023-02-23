# base

基础函数。

## If

三元表达式。

```go
fmt.Println(base.If(true, 1, 2))
fmt.Println(base.If(false, 1, 2))

// 1
// 2
```

[[play](https://go.dev/play/p/MQd_9pR-c_l)]

## Must

检查返回 error 是否为 nil，不是则 panic。

## Must1

检查返回 error，为 nil 则返回1个值。

```go
fmt.Println(base.Must1(1, nil))

// 1
```

[[play](https://go.dev/play/p/FbJFjqeOb5A)]

## Must2

检查返回 error，为 nil 则返回2个值。

```go
fmt.Println(base.Must2(1, 2, nil))

// 1 2
```

[[play](https://go.dev/play/p/0qqWYCzxT_c)]

## Must3

检查返回 error，为 nil 则返回3个值。

```go
fmt.Println(base.Must3(1, 2, 3, nil))

// 1 2 3
```

[[play](https://go.dev/play/p/9MXbOveJFvV)]
