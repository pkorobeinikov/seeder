# Быстрый старт

Создайте конфигурационный файл `seeder.yaml` в корне вашего проекта:

```shell
touch seeder.yaml
```

Добавьте в него "скелет" конфигурации:

<!-- @formatter:off -->
```yaml title="seeder.yaml"
seeder:
  state:
    - name: my first seed
      type: ...
      config:
        - ...
```
<!-- @formatter:on -->

Определите тип хранилища, в примере это будет `postgres`:

<!-- @formatter:off -->
```yaml title="seeder.yaml"
seeder:
  state:
    - name: my first seed
      type: postgres
      config:
        - ...
```
<!-- @formatter:on -->

Опишите конфигурацию:

<!-- @formatter:off -->
```yaml title="seeder.yaml"
seeder:
  state:
    - name: my first seed
      type: postgres
      config:
        - file: my_first_seed.sql
```
<!-- @formatter:on -->

Создайте файл `my_first_seed.sql` и опишите данные в нём:

```shell
touch my_first_seed.sql
```

```sql title="my_first_seed.sql"
insert into person (id, name)
values (gen_random_uuid(), 'Some name');
```

Загрузите данные в `postgres`:

```shell
export SEEDER_PG_CONNSTR="postgres://postgres:secret@localhost:5432/db"

seeder
```

Более подробные инструкции по конфигурации разных типов хранилищ вы найдёте в
разделе "Туториал".
