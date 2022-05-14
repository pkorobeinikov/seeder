# Установка

Выберите наиболее подходящий вам вариант установки:

- Использование `docker`-образа;
- Установка с помощью `go install`;
- Сборка из исходного кода.

## Docker image

!!! Note

    Работа с docker-образом наиболее проста и предпочтительна в современных
    реалиях разработки.

Получите опубликованный образ:

```shell
docker pull ghcr.io/pkorobeinikov/seeder:latest
```

## С помощью `go install`

```shell
go install -v github.com/pkorobeinikov/seeder/seeder/cmd/seeder@latest
```

## Из исходного кода

Получите рабочую копию репозитория `seeder`:

```shell
git clone https://github.com/pkorobeinikov/seeder
```

Перейдите в каталог с исходным кодом проекта:

```shell
cd seeder/seeder
```

Выполните сборку:

```shell
go build -o /usr/local/bin/seeder cmd/seeder/main.go
```
