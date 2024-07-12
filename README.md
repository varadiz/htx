# htx Html-To-teXt

tool for parsing text from html

takes input from stdin

## build

```sh
go build -o bin/htx
```

## how to use

```sh
cat sample.html | bin/htx
```

```sh
wget -q -O - https://example.com/ | bin/htx
```
