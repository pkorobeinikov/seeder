# Vault

В данном руководстве описана загрузка данных (сидирование) в `vault`.

## Настройка окружения

Получите рабочую копию репозитория `seeder`:

```shell
git clone https://github.com/pkorobeinikov/seeder.git
```

Перейдите в каталог `seeder-showcase/vault`:

```shell
cd seeder-showcase/vault
```

Разверните локальное окружение:

```shell
docker compose up [-d]
```

!!! note

    Если у вас не настроено расширение `compose`, самое время перейти на него
    или продолжить использовать устаревшую команду `docker-compose`.

## Изучение демонстрационных данных

Рассмотрите спецификацию описания сидов:

```shell
cat seeder.yaml
```

<!-- @formatter:off -->
```yaml title="seeder-showcase/vault/seeder.yaml"
seeder:
  state:
    - name: vault file seed
      type: vault
      config:
        - key: "secret/data/seed/file/json"
          file: seed/seed.json
        - key: "secret/data/seed/file/yaml"
          file: seed/seed.yaml
        - key: "secret/data/seed/file/yml"
          file: seed/seed.yml
```
<!-- @formatter:on -->

В спецификации описаны три правила загрузки данных в `vault`, сохранённых в
разных форматах в виде файла:

- `json`
- `yaml`
- `yml`

Все три формата эквивалентны и отличаются только правилами форматирования,
зависящими от синтаксиса языка.

## Загрузка данных

Для загрузки данных выполните команду:

```shell
export SEEDER_VAULT_ADDRESS=http://localhost:8200
export SEEDER_VAULT_TOKEN=secret

seeder
```

!!! note

    Если вы не используете `Docker Desktop`, вам, возможно, потребуется
    заменить `localhost` в строке подключения на правильное имя хоста
    или адрес машны с `dockerd`.

## Проверка результата

Выполните следующие команды для проверки загруженных данных в `vault`.

```shell
docker exec \
  -e VAULT_ADDR='http://127.0.0.1:8200' \
  -e VAULT_TOKEN=secret \
  vault vault kv get secret/seed/file/yaml

docker exec \
  -e VAULT_ADDR='http://127.0.0.1:8200' \
  -e VAULT_TOKEN=secret \
  vault vault kv get secret/seed/file/yml

docker exec \
  -e VAULT_ADDR='http://127.0.0.1:8200' \
  -e VAULT_TOKEN=secret \
  vault vault kv get secret/seed/file/json
```

## Резюме

В приведённом руководстве:

- Было развёрнуто локальное окружение с `vault` для загрузки демонстрационных
  данных (сидов);
- Изучен формат описания сидов для `vault`;
- Демонстрационные данные загружены в `vault`.
