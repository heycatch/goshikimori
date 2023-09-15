### [EN](https://github.com/heycatch/goshikimori/blob/master/README.md) | RU

### О проекте
Небольшая библиотека для взаимодействия с шикимори, написанная на языке golang.
* Работа с API происходит только через `OAuth2`.
* Никаких зависимостей от других библиотек.
* Для тестов и сборки используется утилита GNU [make](https://www.gnu.org/software/make/manual/make.html).

### Установка
```bash
go get github.com/heycatch/goshikimori
```

### Готовые примеры
* [Нажать сюда](https://github.com/heycatch/goshikimori/tree/master/examples)

### Документация
В настоящее время доступно большое количество функций.
Для удобства лучше использовать ***godoc***.
```bash
# Способ #1: Используя докер.
git clone git@github.com:heycatch/goshikimori.git && cd goshikimori
make docker-build
make docker-start

# Открыть в браузере.
http://localhost:1337/pkg/github.com/heycatch/goshikimori/
```
```bash
# Способ #2(Linux): Установка godoc.
go install -v golang.org/x/tools/cmd/godoc@latest
# Экспортируем путь(GOPATH).
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
# Проверка работоспособности.
godoc --help

# После установки godoc.
git clone git@github.com:heycatch/goshikimori.git && cd goshikimori
make doc

# Открыть в браузере.
http://localhost:1337/pkg/github.com/heycatch/goshikimori/
```

### Документация шикимори
* [GraphQL](https://shikimori.me/api/doc/graphql)
* [API v1](https://shikimori.me/api/doc/1.0)
* [API v2](https://shikimori.me/api/doc/2.0)
* [OAuth2](https://shikimori.me/oauth)
