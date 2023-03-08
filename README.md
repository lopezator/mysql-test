# mysql-test

Insert works because we are removing the setting `NO_ZERO_DATE`.

If we comment (or remove) the line:

```go
dsn.Params["sql_mode"] = strings.Replace(dsn.Params["sql_mode"], "NO_ZERO_DATE,", "", -1)
```

The insert fails as expected

```
2023/03/08 17:15:55 impossible insert teacher: Error 1292 (22007): Incorrect datetime value: '0000-00-00' for column 'datetime' at row 1
```