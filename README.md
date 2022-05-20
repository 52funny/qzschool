# 趣智校园

趣智校园 Go Api

欢迎提交 PR

### 登陆

```go
q := qzschool.NewQzSchool("username", "password")
err := q.Login()
if err != nil {
    panic(err)
}
```

### Custom 登陆

```go
q := qzschool.NewDetailQzSchool("username", "password", "loginCode", "projectId", "version")
```

### 获取余额

```go
q := qzschool.NewQzSchool("username", "password")
err := q.Login()
if err != nil {
    panic(err)
}
m, err := q.GetAccountMoney()
if err != nil {
    panic(err)
}
fmt.Println(m)
```

### 获取账单

```go
q := qzschool.NewQzSchool("username", "password")
err := q.Login()
if err != nil {
    panic(err)
}
bill, err := q.GetAllMonthBill(BillAllType, "2022-05")
if err != nil {
    panic(err)
}
for i, it := range bill {
    fmt.Println(i+1, it)
}
```
