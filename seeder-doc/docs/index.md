# Введение

> The only tool for reproducible seeding volumes and storages (including
> databases as well).

`seeder` предназначен для наполнения воспроизводимыми данными локальных и
тестовых окружений.

Благодаря строгому и простому файлу конфигурации обеспечивается единый и
переносимый опыт использования инструмента от проекта к проекту:

<!-- @formatter:off -->
```yaml
seeder:
  state:
    # seeding postgres data
    - name: postgres file seed
      type: postgres
      config:
        - file: seed.sql

    # seeding s3 data
    - name: s3 plain text file seed
      type: s3
      config:
        - bucket: "bucket"
          object-name: "seeded/file/seed.txt"
          option:
            content-type: text/plain
            content-encoding: utf8
          file: seed.txt

    # seeding vault secrets
    - name: vault file seed
      type: vault
      config:
        - key: "secret/data/seed/file/json"
          file: seed.json
```
<!-- @formatter:on -->

`seeder` поддерживает различные типы хранилищ и баз данных:

- `postgres`
- `vault`
- `s3`
- `kafka`

По мере развития проекта будут добавлены новые типы.

Особую ценность подход с передачей воспроизводимых данных приобретает при
командной работе:

- Разработчики видят один и тот же набор данных в процессе разработки;
- Они могут добавлять новые данные или изменять уже существующие, используя
  декларативный конфигурационный файл — это существенно облегчает процесс
  передачи тестовых данных между членами команды разработки;
- Новые разработчики наполняют свои локальные базы данных заранее определённым
  набором, что снижает порог входа в систему;
- Тестировщики могут передать состояние системы любому разработчику в
  декларативном виде — это даёт возможность проще и быстрее воспроизводить
  ошибки и исправлять их.
