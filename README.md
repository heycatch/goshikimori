## [EN](https://github.com/heycatch/goshikimori/blob/master/README_en.md) | RU

### О проекте
Небольшая библиотека для взаимодействия с шикимори, написанная на языке golang.
* Работа с API происходит только через `OAuth2`.
И начать нужно с ознакомления документации [первые шаги](https://github.com/heycatch/goshikimori/blob/master/examples/first_steps/README.md).
* Никаких зависимостей от других библиотек.
* Для тестов и сборки используется утилита [GNU make](https://www.gnu.org/software/make/manual/make.html).

### Установка
```bash
go get github.com/heycatch/goshikimori
```

### Готовые примеры
* [Нажать сюда](https://github.com/heycatch/goshikimori/tree/master/examples)

### Документация GODOC / GOPKG
Также доступна поддержка **godoc**.

Или вы можете использовать страницу с официального
[сайта](https://pkg.go.dev/github.com/heycatch/goshikimori) Go pkg.\
**P.S.** документация обновляется с опозданием.
```bash
# Способ #1: Используя докер.
git clone git@github.com:heycatch/goshikimori.git && cd goshikimori
make docker-build
make docker-start

# Открыть в браузере.
http://localhost:1337/pkg/github.com/heycatch/goshikimori
```
```bash
# Способ #2(Linux): Установка godoc.
go install -v golang.org/x/tools/cmd/godoc@latest
# Добавить 'экспорт' в файл /home/$USER/.profile и перезагружаемся.
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
# Проверяем работоспособность.
godoc --help

# После установки или если 'godoc' уже установлен.
git clone git@github.com:heycatch/goshikimori.git && cd goshikimori
make doc

# Открыть в браузере.
http://localhost:1337/pkg/github.com/heycatch/goshikimori
```

### Документация шикимори
* [GraphQL](https://shikimori.one/api/doc/graphql)
* [API v1](https://shikimori.one/api/doc/1.0)
* [API v2](https://shikimori.one/api/doc/2.0)

## Обратная связь
* Написать в личные сообщения на [сайте](https://shikimori.one/arctica).
* Открыть [проблему](https://github.com/heycatch/goshikimori/issues).
```bash
# Текущие задачи и проблемы в коде можно посмотреть в терминале с помощью команды.
git grep TODO
git grep FIXME
git grep NOTES
```
