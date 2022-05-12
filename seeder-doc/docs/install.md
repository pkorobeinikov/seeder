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

Запустите полученный образ:

```shell
docker run --rm --name seeder --workdir /seed -v $(pwd):/seed ghcr.io/pkorobeinikov/seeder
```

## С помощью `go install`

```shell
go install github.com/pkorobeinikov/seeder/seeder/cmd/seeder@latest
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
go build -o seeder seeder/cmd/seeder/main.go
```
