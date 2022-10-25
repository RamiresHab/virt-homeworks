# Домашнее задание к занятию "6.2. SQL"

## Введение

Перед выполнением задания вы можете ознакомиться с 
[дополнительными материалами](https://github.com/netology-code/virt-homeworks/tree/master/additional/README.md).

## Задача 1

Используя docker поднимите инстанс PostgreSQL (версию 12) c 2 volume, 
в который будут складываться данные БД и бэкапы.

Приведите получившуюся команду или docker-compose манифест.

Ответ:
Я воспользовался рекомендациями команды разработчиков PostgreSQL и добавил к нему Volumes и проброс портов для db, чтобы было удобнее работать через Dbeaver, а не только через предлагаемый adminer. В итоге получился такой конфиг:
```
$ cat docker-compose.yaml
# docker-compose.yml
version: '3.1'
services:
  db:
    image: postgres:12
    restart: always
    volumes:
      - ./data:/var/lib/psql/data
      - ./backup:/backup
    ports:
      - 5432:5432
    environment:
      POSTGRES_PASSWORD: example

  adminer:
    image: adminer
    restart: always
    ports:
      - 8080:8080
```

## Задача 2

В БД из задачи 1: 
- создайте пользователя test-admin-user и БД test_db
- в БД test_db создайте таблицу orders и clients (спeцификация таблиц ниже)
- предоставьте привилегии на все операции пользователю test-admin-user на таблицы БД test_db
- создайте пользователя test-simple-user  
- предоставьте пользователю test-simple-user права на SELECT/INSERT/UPDATE/DELETE данных таблиц БД test_db

Таблица orders:
- id (serial primary key)
- наименование (string)
- цена (integer)

Таблица clients:
- id (serial primary key)
- фамилия (string)
- страна проживания (string, index)
- заказ (foreign key orders)

Приведите:
- итоговый список БД после выполнения пунктов выше,
- описание таблиц (describe)
- SQL-запрос для выдачи списка пользователей с правами над таблицами test_db
- список пользователей с правами над таблицами test_db

Ответ:
1. select * from pg_catalog.pg_database 

|oid|datname|datdba|encoding|datcollate|datctype|datistemplate|datallowconn|
|---|-------|------|--------|----------|--------|-------------|------------|
|13458|postgres|10|6|en_US.utf8|en_US.utf8|false|true|
|16384|test_db|10|6|en_US.utf8|en_US.utf8|false|true|
|1|template1|10|6|en_US.utf8|en_US.utf8|true|true|
|13457|template0|10|6|en_US.utf8|en_US.utf8|true|false|

2.
```
SELECT 
   table_name, 
   column_name, 
   data_type 
FROM 
   information_schema.columns
WHERE 
   table_name = 'orders'
 ```
 
|table_name|column_name|data_type|
|----------|-----------|---------|
|orders|id|integer|
|orders|наименование|text|
|orders|цена|integer|

```
SELECT 
   table_name, 
   column_name, 
   data_type 
FROM 
   information_schema.columns
WHERE 
   table_name = 'clients'
```
 
|table_name|column_name|data_type|
|----------|-----------|---------|
|clients|id|integer|
|clients|фамилия|text|
|clients|страна проживания|text|
|clients|заказ|integer|

3.
select distinct grantee from information_schema.table_privileges where table_catalog = 'test_db'

|grantee|
|-------|
|PUBLIC|
|postgres|

   
## Задача 3

Используя SQL синтаксис - наполните таблицы следующими тестовыми данными:

Таблица orders

|Наименование|цена|
|------------|----|
|Шоколад| 10 |
|Принтер| 3000 |
|Книга| 500 |
|Монитор| 7000|
|Гитара| 4000|

Таблица clients

|ФИО|Страна проживания|
|------------|----|
|Иванов Иван Иванович| USA |
|Петров Петр Петрович| Canada |
|Иоганн Себастьян Бах| Japan |
|Ронни Джеймс Дио| Russia|
|Ritchie Blackmore| Russia|

Используя SQL синтаксис:
- вычислите количество записей для каждой таблицы 
- приведите в ответе:
    - запросы 
    - результаты их выполнения.

Ответ:
```
insert into orders (наименование, цена) values
	('Шоколад', 10),
	('Принтер', 3000),
	('Книга', 500),
	('Монитор', 7000),
	('Гитара', 4000)
insert into clients (фамилия, "страна проживания") values
	('Иванов Иван Иванович', 'USA'),
	('Петров Петр Петрович', 'Canada'),
	('Иоганн Себастьян Бах', 'Japan'),
	('Ронни Джеймс Дио', 'Russia'),
	('Ritchie Blackmore', 'Russia') 
```
select count(*) from orders

|count|
|-----|
|5|

select count(*) from clients

|count|
|-----|
|5|


## Задача 4

Часть пользователей из таблицы clients решили оформить заказы из таблицы orders.

Используя foreign keys свяжите записи из таблиц, согласно таблице:

|ФИО|Заказ|
|------------|----|
|Иванов Иван Иванович| Книга |
|Петров Петр Петрович| Монитор |
|Иоганн Себастьян Бах| Гитара |

Приведите SQL-запросы для выполнения данных операций.

Приведите SQL-запрос для выдачи всех пользователей, которые совершили заказ, а также вывод данного запроса.
 
Подсказк - используйте директиву `UPDATE`.

Ответ:
```
update clients set заказ = (select id from orders where "наименование" = 'Книга') where "фамилия" = 'Иванов Иван Иванович'
update clients set заказ = (select id from orders where "наименование" = 'Монитор') where "фамилия" = 'Петров Петр Петрович'
update clients set заказ = (select id from orders where "наименование" = 'Гитара') where "фамилия" = 'Иоганн Себастьян Бах'
select c."фамилия", o."наименование" from clients c join orders o on c."заказ"=o.id 
```

|фамилия|наименование|
|-------|------------|
|Иванов Иван Иванович|Книга|
|Петров Петр Петрович|Монитор|
|Иоганн Себастьян Бах|Гитара|

## Задача 5

Получите полную информацию по выполнению запроса выдачи всех пользователей из задачи 4 
(используя директиву EXPLAIN).

Приведите получившийся результат и объясните что значат полученные значения.

Ответ:
explain select c."фамилия", o."наименование" from clients c join orders o on c."заказ"=o.id 

|QUERY PLAN|
|----------|
|Hash Join  (cost=37.00..57.24 rows=810 width=64)|
|  Hash Cond: (c."заказ" = o.id)|
|  ->  Seq Scan on clients c  (cost=0.00..18.10 rows=810 width=36)|
|  ->  Hash  (cost=22.00..22.00 rows=1200 width=36)|
|        ->  Seq Scan on orders o  (cost=0.00..22.00 rows=1200 width=36)|

Explain говорит нам, что планировщик сначала выполняет сканирование по таблице orders, потом сканирование по таблице clients, потом JOIN. Цифры после cost показывают время ожидаемого начала и завершения операции, ROWS - оценка количества строк, которые должна просмотреть СУБД, WIDTH - ожидаемый размер каждой строки в байтах.

## Задача 6

Создайте бэкап БД test_db и поместите его в volume, предназначенный для бэкапов (см. Задачу 1).

Остановите контейнер с PostgreSQL (но не удаляйте volumes).

Поднимите новый пустой контейнер с PostgreSQL.

Восстановите БД test_db в новом контейнере.

Приведите список операций, который вы применяли для бэкапа данных и восстановления. 

---

### Как cдавать задание

Выполненное домашнее задание пришлите ссылкой на .md-файл в вашем репозитории.

---
