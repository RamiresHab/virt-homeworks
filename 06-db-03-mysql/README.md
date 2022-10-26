# Домашнее задание к занятию "6.3. MySQL"

## Введение

Перед выполнением задания вы можете ознакомиться с 
[дополнительными материалами](https://github.com/netology-code/virt-homeworks/tree/master/additional/README.md).

## Задача 1

Используя docker поднимите инстанс MySQL (версию 8). Данные БД сохраните в volume.

Изучите [бэкап БД](https://github.com/netology-code/virt-homeworks/tree/master/06-db-03-mysql/test_data) и 
восстановитесь из него.

Перейдите в управляющую консоль `mysql` внутри контейнера.

Используя команду `\h` получите список управляющих команд.

Найдите команду для выдачи статуса БД и **приведите в ответе** из ее вывода версию сервера БД.

Подключитесь к восстановленной БД и получите список таблиц из этой БД.

**Приведите в ответе** количество записей с `price` > 300.

В следующих заданиях мы будем продолжать работу с данным контейнером.

Ответ:

```
vagrant@vagrant:~/06-db-03-mysql$ cat docker-compose.yaml
# Use root/example as user/password credentials
version: '3.1'

services:

  db:
    image: mysql:8
    command: --default-authentication-plugin=mysql_native_password
    restart: always
    environment:
      MYSQL_ROOT_PASSWORD: example
    volumes:
      - ./data:/var/lib/mysql
    ports:
      - 3306:3306

  adminer:
    image: adminer
    restart: always
    ports:
      - 8080:8080

CREATE database test_db #Создание базы

vagrant@vagrant:~/06-db-03-mysql$ sudo docker exec -it 06-db-03-mysql_db_1 bash
bash-4.4# mysql -p test_db < /var/lib/mysql/test_dump.sql #Восстановление базы из бэкапа
Enter password:

vagrant@vagrant:~/06-db-03-mysql$ sudo docker exec -it 06-db-03-mysql_db_1 mysql test_db -p #Переход в mysql внутри контейнера
Enter password:
Reading table information for completion of table and column names
You can turn off this feature to get a quicker startup with -A

Welcome to the MySQL monitor.  Commands end with ; or \g.
Your MySQL connection id is 22
Server version: 8.0.31 MySQL Community Server - GPL

Copyright (c) 2000, 2022, Oracle and/or its affiliates.

Oracle is a registered trademark of Oracle Corporation and/or its
affiliates. Other names may be trademarks of their respective
owners.

Type 'help;' or '\h' for help. Type '\c' to clear the current input statement.

mysql>
mysql> \h

For information about MySQL products and services, visit:
   http://www.mysql.com/
For developer information, including the MySQL Reference Manual, visit:
   http://dev.mysql.com/
To buy MySQL Enterprise support, training, or other products, visit:
   https://shop.mysql.com/

List of all MySQL commands:
Note that all text commands must be first on line and end with ';'
?         (\?) Synonym for `help'.
clear     (\c) Clear the current input statement.
connect   (\r) Reconnect to the server. Optional arguments are db and host.
delimiter (\d) Set statement delimiter.
edit      (\e) Edit command with $EDITOR.
ego       (\G) Send command to mysql server, display result vertically.
exit      (\q) Exit mysql. Same as quit.
go        (\g) Send command to mysql server.
help      (\h) Display this help.
nopager   (\n) Disable pager, print to stdout.
notee     (\t) Don't write into outfile.
pager     (\P) Set PAGER [to_pager]. Print the query results via PAGER.
print     (\p) Print current command.
prompt    (\R) Change your mysql prompt.
quit      (\q) Quit mysql.
rehash    (\#) Rebuild completion hash.
source    (\.) Execute an SQL script file. Takes a file name as an argument.
status    (\s) Get status information from the server.
system    (\!) Execute a system shell command.
tee       (\T) Set outfile [to_outfile]. Append everything into given outfile.
use       (\u) Use another database. Takes database name as argument.
charset   (\C) Switch to another charset. Might be needed for processing binlog with multi-byte charsets.
warnings  (\W) Show warnings after every statement.
nowarning (\w) Don't show warnings after every statement.
resetconnection(\x) Clean session context.
query_attributes Sets string parameters (name1 value1 name2 value2 ...) for the next query to pick up.
ssl_session_data_print Serializes the current SSL session data to stdout or file
For server side help, type 'help contents'

mysql  Ver 8.0.31 for Linux on x86_64 (MySQL Community Server - GPL) #версия сервера из вывода команды status

show tables from test_db
```

|Tables_in_test_db|
|-----------------|
|orders|

```
select count(*) from orders where price > 300
```

|count(*)|
|--------|
|1|


## Задача 2

Создайте пользователя test в БД c паролем test-pass, используя:
- плагин авторизации mysql_native_password
- срок истечения пароля - 180 дней 
- количество попыток авторизации - 3 
- максимальное количество запросов в час - 100
- аттрибуты пользователя:
    - Фамилия "Pretty"
    - Имя "James"

Предоставьте привелегии пользователю `test` на операции SELECT базы `test_db`.
    
Используя таблицу INFORMATION_SCHEMA.USER_ATTRIBUTES получите данные по пользователю `test` и 
**приведите в ответе к задаче**.

Ответ:

```
CREATE USER 'test'
  IDENTIFIED WITH mysql_native_password BY 'test-pass'
  PASSWORD EXPIRE INTERVAL 180 DAY
  FAILED_LOGIN_ATTEMPTS 3
  ATTRIBUTE '{"fname": "James", "lname": "Pretty"}'

alter USER 'test' with MAX_QUERIES_PER_HOUR 100

GRANT SELECT on test_db.* to 'test'

select * from INFORMATION_SCHEMA.USER_ATTRIBUTES where user='test'
```

|USER|HOST|ATTRIBUTE|
|----|----|---------|
|test|%|{"fname": "James", "lname": "Pretty"}|


## Задача 3

Установите профилирование `SET profiling = 1`.
Изучите вывод профилирования команд `SHOW PROFILES;`.

Исследуйте, какой `engine` используется в таблице БД `test_db` и **приведите в ответе**.

Измените `engine` и **приведите время выполнения и запрос на изменения из профайлера в ответе**:
- на `MyISAM`
- на `InnoDB`

Ответ:
```
SELECT TABLE_NAME, ENGINE
FROM   information_schema.TABLES
WHERE  TABLE_SCHEMA = 'test_db'
```

|TABLE_NAME|ENGINE|
|----------|------|
|orders|InnoDB|

|Query_ID|Duration|Query|
|--------|--------|-----|
|2|0.0006515|/* ApplicationName=DBeaver 22.2.3 - SQLEditor <Script-2.sql> */ SELECT DATABASE()|
|3|0.00076975|/* ApplicationName=DBeaver 22.2.3 - SQLEditor <Script-2.sql> */ SET SQL_SELECT_LIMIT=200|
|4|0.0006195|SHOW WARNINGS|
|5|0.00058175|/* ApplicationName=DBeaver 22.2.3 - SQLEditor <Script-2.sql> */ SET SQL_SELECT_LIMIT=DEFAULT|
|6|0.0004865|/* ApplicationName=DBeaver 22.2.3 - SQLEditor <Script-2.sql> */ SELECT DATABASE()|
|7|0.00033275|/* ApplicationName=DBeaver 22.2.3 - SQLEditor <Script-2.sql> */ SET SQL_SELECT_LIMIT=200|
|8|0.00039775|/* ApplicationName=DBeaver 22.2.3 - SQLEditor <Script-2.sql> */ SHOW ENGINE|
|9|0.0008255|/* ApplicationName=DBeaver 22.2.3 - SQLEditor <Script-2.sql> */ SHOW ENGINES|
|10|0.0007805|/* ApplicationName=DBeaver 22.2.3 - SQLEditor <Script-2.sql> */ SET SQL_SELECT_LIMIT=DEFAULT|
|11|0.00066675|/* ApplicationName=DBeaver 22.2.3 - SQLEditor <Script-2.sql> */ SELECT DATABASE()|
|12|0.00179725|/* ApplicationName=DBeaver 22.2.3 - SQLEditor <Script-2.sql> */ SELECT TABLE_NAME, ENGINE FROM information_schema.TABLES WHERE TABLE_SCHEMA = 'test_db' LIMIT 0, 200|
|13|0.00068575|/* ApplicationName=DBeaver 22.2.3 - SQLEditor <Script-2.sql> */ SELECT DATABASE()|
|14|0.00163675|/* ApplicationName=DBeaver 22.2.3 - SQLEditor <Script-2.sql> */ SELECT TABLE_NAME, ENGINE FROM information_schema.TABLES WHERE TABLE_SCHEMA = 'test_db' LIMIT 0, 200|
|15|0.00070775|/* ApplicationName=DBeaver 22.2.3 - SQLEditor <Script-2.sql> */ SELECT DATABASE()|
|16|0.000388|/* ApplicationName=DBeaver 22.2.3 - SQLEditor <Script-2.sql> */ SET SQL_SELECT_LIMIT=200|

alter table orders ENGINE='MyISAM'

|TABLE_NAME|ENGINE|
|----------|------|
|orders|MyISAM|

|Query_ID|Duration|Query|
|--------|--------|-----|
|17|0.00068725|SHOW WARNINGS|
|18|0.00060525|/* ApplicationName=DBeaver 22.2.3 - SQLEditor <Script-2.sql> */ SET SQL_SELECT_LIMIT=DEFAULT|
|19|0.00084425|/* ApplicationName=DBeaver 22.2.3 - SQLEditor <Script-2.sql> */ SELECT DATABASE()|
|20|0.00067025|SELECT @@session.transaction_read_only|
|21|0.0002875|/* ApplicationName=DBeaver 22.2.3 - SQLEditor <Script-2.sql> */ alter DATABASE test_db ENGINE = MyISAM|
|22|0.00081825|SELECT @@session.transaction_read_only|
|23|0.00057975|/* ApplicationName=DBeaver 22.2.3 - SQLEditor <Script-2.sql> */ alter DATABASE test_db ENGINE = 'MyISAM'|
|24|0.001026|SELECT @@session.transaction_read_only|
|25|0.00047|/* ApplicationName=DBeaver 22.2.3 - SQLEditor <Script-2.sql> */ alter DATABASE test_db ENGINE='MyISAM'|
|26|0.00030025|SELECT @@session.transaction_read_only|
|27|0.2089825|/* ApplicationName=DBeaver 22.2.3 - SQLEditor <Script-2.sql> */ alter table orders ENGINE='MyISAM'|
|28|0.00066375|/* ApplicationName=DBeaver 22.2.3 - SQLEditor <Script-2.sql> */ SELECT DATABASE()|
|29|0.001567|/* ApplicationName=DBeaver 22.2.3 - SQLEditor <Script-2.sql> */ SELECT TABLE_NAME, ENGINE FROM information_schema.TABLES WHERE TABLE_SCHEMA = 'test_db' LIMIT 0, 200|
|30|0.00026|/* ApplicationName=DBeaver 22.2.3 - SQLEditor <Script-2.sql> */ SELECT DATABASE()|
|31|0.00057975|/* ApplicationName=DBeaver 22.2.3 - SQLEditor <Script-2.sql> */ SET SQL_SELECT_LIMIT=200|



## Задача 4 

Изучите файл `my.cnf` в директории /etc/mysql.

Измените его согласно ТЗ (движок InnoDB):
- Скорость IO важнее сохранности данных
- Нужна компрессия таблиц для экономии места на диске
- Размер буффера с незакомиченными транзакциями 1 Мб
- Буффер кеширования 30% от ОЗУ
- Размер файла логов операций 100 Мб

Приведите в ответе измененный файл `my.cnf`.

Ответ:
```
bash-4.4# cat /etc/my.cnf
# For advice on how to change settings please see
# http://dev.mysql.com/doc/refman/8.0/en/server-configuration-defaults.html

[mysqld]
#
# Remove leading # and set to the amount of RAM for the most important data
# cache in MySQL. Start at 70% of total RAM for dedicated server, else 10%.
# innodb_buffer_pool_size = 128M
#
# Remove leading # to turn on a very important data integrity option: logging
# changes to the binary log between backups.
# log_bin
#
# Remove leading # to set options mainly useful for reporting servers.
# The server defaults are faster for transactions and fast SELECTs.
# Adjust sizes as needed, experiment to find the optimal values.
# join_buffer_size = 128M
# sort_buffer_size = 2M
# read_rnd_buffer_size = 2M

# Remove leading # to revert to previous value for default_authentication_plugin,
# this will increase compatibility with older clients. For background, see:
# https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_default_authentication_plugin
# default-authentication-plugin=mysql_native_password
skip-host-cache
skip-name-resolve
datadir=/var/lib/mysql
socket=/var/run/mysqld/mysqld.sock
secure-file-priv=/var/lib/mysql-files
user=mysql

pid-file=/var/run/mysqld/mysqld.pid
[client]
socket=/var/run/mysqld/mysqld.sock

!includedir /etc/mysql/conf.d/
innodb_flush_log_at_trx_commit=2
innodb_log_buffer_size=1M
innodb_buffer_pool_size=300M
innodb_log_file_size=100M
innodb_file_per_table=1
innodb_file_format=Barracuda
```
---

### Как оформить ДЗ?

Выполненное домашнее задание пришлите ссылкой на .md-файл в вашем репозитории.

---
