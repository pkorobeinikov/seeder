# Postgres

В данном руководстве описана загрузка данных (сидирование) в `postgres`.

## Настройка окружения

Получите рабочую копию репозитория `seeder`:

```shell
git clone https://github.com/pkorobeinikov/seeder.git
```

Перейдите в каталог `seeder-showcase/postgres`:

```shell
cd seeder-showcase/postgres
```

Разверните локальное окружение:

```shell
docker compose up [-d]
```

!!! note

    Если у вас не настроено расширение `compose`, самое время перейти на него
    или продолжить использовать устаревшую команду `docker-compose`.

## Подготовка базы данных

Войдите в контейнер с `postgres`:

```shell
docker exec -it --user postgres postgres /bin/bash
```

Вы окажетесь в смонтированном в `/showcase` текущем каталоге.

Создайте структуру таблиц для загрузки данных:

```shell
psql -d seeder -f schema/setup.sql
```

Убедитесь, что схема данных создана успешно.

Из контейнера подключитесь к базе данных `seeder`:

```shell
psql -d seeder
```

В приглашении `psql` выполните команды:

```shell
\dn
\dt app.*
```

Если всё было сделано правильно, в результате выполнения описанных шагов в базе
данных `seeder` появятся схема `app` и таблица `person`:

```
seeder=# \dn
  List of schemas
  Name  |  Owner
--------+----------
 app    | postgres
 public | postgres
(2 rows)

seeder=# \dt app.*
         List of relations
 Schema |  Name  | Type  |  Owner
--------+--------+-------+----------
 app    | person | table | postgres
(1 row)
```

После подготовки базы данных вернитесь в рабочий
каталог `seeder-showcase/postgres`.

## Изучение демонстрационных данных

Рассмотрите спецификацию описания сидов:

```shell
cat seeder.yaml
```

<!-- @formatter:off -->
```yaml title="seeder-showcase/postgres/seeder.yaml"
seeder:
  state:
    - name: postgres file seed
      type: postgres # (1)
      config:
        - file: seed/seed.sql # (2)

    - name: postgres plpgsql file seed
      type: postgres # (3)
      config:
        - file: seed/seed_plpgsql.sql # (4)
```
<!-- @formatter:on -->

1. Тип целевого хранилища
2. Путь до файла с данными
3. Тип целевого хранилища
4. Путь до файла с данными

В спецификации описано два сида: `seed/seed.sql` и `seed/seed_plpgsql.sql` —
первый использует обычные команды `sql`, второй — блок кода на `plpgsql`,
обёрнутый в оператор `do`.

Принциписальной разницы между примерами нет, однако блок кода на `plpgsql`
обладает большей гибкостью, чем обычный `sql`.

## Загрузка данных

Для загрузки данных выполните команду:

```shell
export SEEDER_PG_CONNSTR="postgres://postgres:secret@localhost:5432/seeder"

seeder
```

!!! note

    Если вы не используете `Docker Desktop`, вам, возможно, потребуется
    заменить `localhost` в строке подключения на правильное имя хоста
    или адрес машины с `dockerd`.

## Проверка результата

Вернитесь в контейнер с базой данных или подключитесь к ней удобным для вас
клиентом:

```shell
docker exec -it --user postgres postgres psql -d seeder
```

Выполните выборку по таблице `app.person`:

```
table app.person \g
```

Если всё прошло успешно, вы увидите набор сидированных данных:

```
seeder=# table app.person \g
                  id                  |   name    |  surname
--------------------------------------+-----------+------------
 ce4b777a-0260-4138-a691-91459cf24879 | John      | Doe
 d11f3af4-75f1-4f7d-9201-064584d51766 | Kelvin    | Houston
 1aa51043-0e5c-4104-a9f0-1da15b351d8b | Brett     | Vaught
 94997eb1-c39c-4602-a526-b5a2f5479663 | Jane      | Doe
 5e533e3d-d06b-4b81-ad01-e0312e85c465 | Kathryn   | Fee
 b7fe5a7d-2704-4491-ab58-7db2dc216207 | Margaret  | Martinez
 a0ca87cc-1804-4ca3-a847-e01fb5e3812b | Frederick | Simpson
 774ac047-485d-48ad-b1ad-3d7158a10858 | James     | Anderson
 422e02d7-e80d-4d19-8c71-4885bd6f00b5 | Timothy   | Pennington
 4de86230-dd75-4946-bf85-d597b4ff2951 | Dorothy   | Martinez
 9fdca959-671a-4613-9df4-2f16c8733f83 | Pamela    | Lawson
 f8f7bafd-93c4-4975-b699-8093dca10d00 | Barbara   | Roth
```

## Резюме

В приведённом руководстве:

- Было развёрнуто локальное окружение с `postgres` для загрузки демонстрационных
  данных (сидов);
- Изучен формат описания сидов для `postgres`;
- Демонстрационные данные загружены в `postgres`.
