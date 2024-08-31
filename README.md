# The Gilded Rose Kata

This 'go starter' is copied from the following address (with a couple _minor_ tweaks): 
https://github.com/emilybache/GildedRose-Refactoring-Kata/tree/main/go



# GO Starter

- Run :

```shell
go run texttest_fixture.go [<number-of-days>; default: 2]
```

- Run tests :

```shell
go test ./...
```

- Run tests and coverage :

```shell
go test ./... -coverprofile=coverage.out

go tool cover -html=coverage.out
```