# MergeMinds | Go backend

## Как это запускать?

### Пояснения

1. Dragonfly - это Key-Value база данных, которая имеет полную совместимость с API редиски, но не является её форком. А ещё довольно производительная, многопоточная и вообще...
2. **Если вы работаете на Linux, вероятно потребуется ввести sudo перед командами, которые содержат в себе docker.**

### Нулевой шаг - установить Just

Just - это command runner. Команды берутся из файла с названием justfile. Список всех доступных команд можно получить, запустив: `just -l`. Вы можете использовать эту команду, как ещё одну документацию для запуска проекта.

### Docker Compose

**Данный Docker Compose не рекомендуется использовать для продакшена, так как он не устанавливает пароли для БД и вообще возможно сыроват. Зато фронтендерам будет удобно (наверно).**

```sh
just full-run-compose
```

Данная команда **очищает всю базу данных**, если в ней что-то было, после чего заново её инициализурует, после чего уже запускает само приложение вместе с необходимыми сервисами.

Если вы хотите сами контроллировать, когда инициализировать и очищать базу данных, вы можете использовать команды `just initdb-compose`, `just dropdb-compose` и `just run-compose`.

Также вы можете запустить приложение в режиме разработки. В таком случае оно запуститься в Docker Compose через Air, которые автоматически перекомплириует проект при изменении файлов. Для этого вам также нужно *инициализировать базу данных* при помощи `just initdb-compose` и запустить проект при помощи `just dev`.

Готово! Ваш бекенд готов к изнурительной работе на господина-фронтенда.

### Хочу сам севрер локально запустить, а всё остальное через Docker

#### Первый шаг - запускаем PostgreSQL и Dragonfly

```sh
just rundb-docker
just runfly-docker
```

#### Второй шаг - копируем .env.example в .env и меняем по вкусу

```sh
cp .env.example .env
nvim .env
```

#### Третий шаг - инициализируем БД

```sh
just initdb-host
```

#### Четвёртый шаг - запускаем проект

```sh
just run-host
```

## Я хочу сюда комитить, чтобы всё было красиво

Тогда выполни эту команду (предварительно установив в свою ОС pre-commit):

```sh
just precommit-install
```

## Что ещё?

1. Во-первых, всё это чудо, как вы могли догадаться, написано на Go. Просто потому что он хайповый, производительный, компилируется быстро, он простой, все дела.
2. В качестве либы для PostgreSQL используется pgx, просто потому что он не deprecated, в отличие от некоторых.
3. Для дракоши используется go-redis. Дракономух идеально умеет мимикрировать под редиску, поэтому и используем эту либу.
4. Для логов используем Zap. Быстрый, есть настройка уровня логгирования, написан Uber'ом. Why not, как говорится.
5. В качестве веб-фреймворка используем Gin. Поддерживает валидацию JSON, хайповый, а больше ничего и не надо.

Также, если хотите, вы можете почистить БД от ваших шалостей. Для этого вы можете выполнить одну из приведённых команд (в зависимости от вашего способа запуска):

```sh
just dropdb-compose
```

или

```sh
just dropdb-host
```

## Всё!

Теперь мы можем в полной мере наслаждаться бекендом на Go :)
