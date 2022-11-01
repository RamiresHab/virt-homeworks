# Домашнее задание к занятию "6.5. Elasticsearch"

## Задача 1

В этом задании вы потренируетесь в:
- установке elasticsearch
- первоначальном конфигурировании elastcisearch
- запуске elasticsearch в docker

Используя докер образ [elasticsearch:7](https://hub.docker.com/_/elasticsearch) как базовый:

- составьте Dockerfile-манифест для elasticsearch
- соберите docker-образ и сделайте `push` в ваш docker.io репозиторий
- запустите контейнер из получившегося образа и выполните запрос пути `/` c хост-машины

Требования к `elasticsearch.yml`:
- данные `path` должны сохраняться в `/var/lib` 
- имя ноды должно быть `netology_test`

В ответе приведите:
- текст Dockerfile манифеста
- ссылку на образ в репозитории dockerhub
- ответ `elasticsearch` на запрос пути `/` в json виде

Подсказки:
- при сетевых проблемах внимательно изучите кластерные и сетевые настройки в elasticsearch.yml
- при некоторых проблемах вам поможет docker директива ulimit
- elasticsearch в логах обычно описывает проблему и пути ее решения
- обратите внимание на настройки безопасности такие как `xpack.security.enabled` 
- если докер образ не запускается и падает с ошибкой 137 в этом случае может помочь настройка `-e ES_HEAP_SIZE`
- при настройке `path` возможно потребуется настройка прав доступа на директорию

Далее мы будем работать с данным экземпляром elasticsearch.

Ответ:
Я не смог загрузить контейнер с тэгом :7, поэтому воспользовался последним релизом седьмой версии :7.17.7
```
vagrant@vagrant:~/06-db-05-elasticsearch$ cat Dockerfile 
FROM elasticsearch:7.17.7
RUN mkdir /var/lib/elasticsearch && chown elasticsearch:elasticsearch /var/lib/elasticsearch
COPY --chown=elasticsearch:elasticsearch elasticsearch.yml /usr/share/elasticsearch/config/

vagrant@vagrant:~/06-db-05-elasticsearch$ sudo docker ps
CONTAINER ID   IMAGE                                     COMMAND                  CREATED              STATUS              PORTS                                                                                  NAMES
37f0f1ac5a88   ramireshab/elasticsearch-netology:1.0.0   "/bin/tini -- /usr/l…"   About a minute ago   Up About a minute   0.0.0.0:9200->9200/tcp, :::9200->9200/tcp, 0.0.0.0:9300->9300/tcp, :::9300->9300/tcp   elasticsearch

vagrant@vagrant:~/06-db-05-elasticsearch$ curl localhost:9200/
{
  "name" : "netology_test",
  "cluster_name" : "docker-cluster",
  "cluster_uuid" : "xYcG22SQTfiIzSAIpleKAQ",
  "version" : {
    "number" : "7.17.7",
    "build_flavor" : "default",
    "build_type" : "docker",
    "build_hash" : "78dcaaa8cee33438b91eca7f5c7f56a70fec9e80",
    "build_date" : "2022-10-17T15:29:54.167373105Z",
    "build_snapshot" : false,
    "lucene_version" : "8.11.1",
    "minimum_wire_compatibility_version" : "6.8.0",
    "minimum_index_compatibility_version" : "6.0.0-beta1"
  },
  "tagline" : "You Know, for Search"
}
```

Вот ссылка на dockerhub: https://hub.docker.com/repository/docker/ramireshab/elasticsearch-netology


## Задача 2

В этом задании вы научитесь:
- создавать и удалять индексы
- изучать состояние кластера
- обосновывать причину деградации доступности данных

Ознакомтесь с [документацией](https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-create-index.html) 
и добавьте в `elasticsearch` 3 индекса, в соответствии со таблицей:

| Имя | Количество реплик | Количество шард |
|-----|-------------------|-----------------|
| ind-1| 0 | 1 |
| ind-2 | 1 | 2 |
| ind-3 | 2 | 4 |

Получите список индексов и их статусов, используя API и **приведите в ответе** на задание.

Получите состояние кластера `elasticsearch`, используя API.

Как вы думаете, почему часть индексов и кластер находится в состоянии yellow?

Удалите все индексы.

**Важно**

При проектировании кластера elasticsearch нужно корректно рассчитывать количество реплик и шард,
иначе возможна потеря данных индексов, вплоть до полной, при деградации системы.

Ответ:
Список индексов и их статусов:
```
vagrant@vagrant:~/06-db-05-elasticsearch$ curl -X GET "localhost:9200/_cat/indices?pretty"
green  open .geoip_databases 5DdTMuHMT-mgWoBc9-11nA 1 0 40 0 38.4mb 38.4mb
green  open ind-1            DrVdqJxZSNSRslQXbnQUEA 1 0  0 0   226b   226b
yellow open ind-3            gRxb_vcuQwmR4lH2RjlaxA 4 2  0 0   904b   904b
yellow open ind-2            3Q0HmrTVQ1iqD9yolGdJQw 2 1  0 0   452b   452b
```

Состояние кластера:
```
vagrant@vagrant:~/06-db-05-elasticsearch$ curl -X GET "localhost:9200/_cluster/health?pretty"
{
  "cluster_name" : "docker-cluster",
  "status" : "yellow",
  "timed_out" : false,
  "number_of_nodes" : 1,
  "number_of_data_nodes" : 1,
  "active_primary_shards" : 10,
  "active_shards" : 10,
  "relocating_shards" : 0,
  "initializing_shards" : 0,
  "unassigned_shards" : 10,
  "delayed_unassigned_shards" : 0,
  "number_of_pending_tasks" : 0,
  "number_of_in_flight_fetch" : 0,
  "task_max_waiting_in_queue_millis" : 0,
  "active_shards_percent_as_number" : 50.0
}
```
Часть индексов и кластеров находится в состоянии yellow потому, что есть unattached реплики. А реплики в этом состоянии потому, что у нас всего одна нода, их некуда приаттачить.

## Задача 3

В данном задании вы научитесь:
- создавать бэкапы данных
- восстанавливать индексы из бэкапов

Создайте директорию `{путь до корневой директории с elasticsearch в образе}/snapshots`.

Используя API [зарегистрируйте](https://www.elastic.co/guide/en/elasticsearch/reference/current/snapshots-register-repository.html#snapshots-register-repository) 
данную директорию как `snapshot repository` c именем `netology_backup`.

**Приведите в ответе** запрос API и результат вызова API для создания репозитория.

Создайте индекс `test` с 0 реплик и 1 шардом и **приведите в ответе** список индексов.

[Создайте `snapshot`](https://www.elastic.co/guide/en/elasticsearch/reference/current/snapshots-take-snapshot.html) 
состояния кластера `elasticsearch`.

**Приведите в ответе** список файлов в директории со `snapshot`ами.

Удалите индекс `test` и создайте индекс `test-2`. **Приведите в ответе** список индексов.

[Восстановите](https://www.elastic.co/guide/en/elasticsearch/reference/current/snapshots-restore-snapshot.html) состояние
кластера `elasticsearch` из `snapshot`, созданного ранее. 

**Приведите в ответе** запрос к API восстановления и итоговый список индексов.

Подсказки:
- возможно вам понадобится доработать `elasticsearch.yml` в части директивы `path.repo` и перезапустить `elasticsearch`

Ответ:
```
vagrant@vagrant:~/06-db-05-elasticsearch$ curl -X PUT "localhost:9200/_snapshot/netology_backup?pretty" -H 'Content-Type: application/json' -d'
{              
  "type": "fs",
  "settings": {                                     
    "location": "/usr/share/elasticsearch/snapshots"
  }
}
'
{
  "acknowledged" : true
}

vagrant@vagrant:~/06-db-05-elasticsearch$ curl -X GET "localhost:9200/_cat/indices?pretty"
green open .geoip_databases hyKhhw4BQEqSit0bsJhMTw 1 0 40 0 38.4mb 38.4mb
green open test             YniCNgQXSnWyDLYyTzgwGQ 1 0  0 0   226b   226b

vagrant@vagrant:~/06-db-05-elasticsearch$ sudo docker exec -it 8c39bdf78af8 bash -c "ls snapshots" 
index-0       indices                          snap-gcxYnu7LTnGkh0AXYdSU3w.dat
index.latest  meta-gcxYnu7LTnGkh0AXYdSU3w.dat

vagrant@vagrant:~/06-db-05-elasticsearch$ curl -X GET "localhost:9200/_cat/indices?pretty"
green open test-2           JqGvs19GTT6Q-rsQQg03ug 1 0  0 0   226b   226b
green open .geoip_databases hyKhhw4BQEqSit0bsJhMTw 1 0 40 0 38.4mb 38.4mb

vagrant@vagrant:~/06-db-05-elasticsearch$ curl -X POST "localhost:9200/_snapshot/netology_backup/my_snapshot_2022.11.01/_restore?pretty" -H 'Content-Type: application/json' -d'
> {
>   "indices": "test"
> }
> '
{
  "accepted" : true
}

vagrant@vagrant:~/06-db-05-elasticsearch$ curl -X GET "localhost:9200/_cat/indices?pretty"
green open test-2           JqGvs19GTT6Q-rsQQg03ug 1 0  0 0   226b   226b
green open .geoip_databases hyKhhw4BQEqSit0bsJhMTw 1 0 40 0 38.4mb 38.4mb
green open test             _w-2XezYSzWYzypsovF0HA 1 0  0 0   226b   226b
```

---

### Как cдавать задание

Выполненное домашнее задание пришлите ссылкой на .md-файл в вашем репозитории.

---
