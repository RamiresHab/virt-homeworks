# Домашнее задание к занятию "5.5. Оркестрация кластером Docker контейнеров на примере Docker Swarm"

## Задача 1

Дайте письменые ответы на следующие вопросы:

- В чём отличие режимов работы сервисов в Docker Swarm кластере: replication и global?
- Какой алгоритм выбора лидера используется в Docker Swarm кластере?
- Что такое Overlay Network?

Ответ:
1. Режим работы global подразумевает репликацию сервиса на каждую ноду (как manager, так и worker), режим работы replicated подразумевает репликацию сервиса на заданное число worker нод. по умолчанию, если не задано иного, то сервисы создаются в replicated режиме.
2. Для выбора лидера используется алгоритм распределённого консенсуса RAFT. У каждой ноды есть рандомный таймер, который обнуляется при получение запроса от лидера. Если лидера нет, нода становится кандидатом и рассылает запрос на лидера всем другим нодам. Если большинство отвечает согласием, то нода становится лидером. Дальше лидер постоянно рассылает запросы другим нодам, которые обнуляют их таймеры.
3. Overlay - это тип сети для docker контейнеров, который соединяет их в единую логическую сеть. Контейнеры могут находиться на одном хосте или на разных, для сетевой связности между разными хостами одной overlay сети надо открыть порт UDP 4789.

## Задача 2

Создать ваш первый Docker Swarm кластер в Яндекс.Облаке

Для получения зачета, вам необходимо предоставить скриншот из терминала (консоли), с выводом команды:
```
docker node ls
```

Ответ:
```
[centos@node01 ~]$ sudo docker node ls
ID                            HOSTNAME             STATUS    AVAILABILITY   MANAGER STATUS   ENGINE VERSION
rnnkplc0f3n1pe01plaqzguwu *   node01.netology.yc   Ready     Active         Leader           20.10.20
pgzp8ph96vbgp27p3tzp7hw2x     node02.netology.yc   Ready     Active         Reachable        20.10.20
3i80fivg2oyapl7p8zv8m6p9a     node03.netology.yc   Ready     Active         Reachable        20.10.20
y8wcat15exstoc4bt0vmy6bf6     node04.netology.yc   Ready     Active                          20.10.20
fc68kwxng4jpolrw624vyukxs     node05.netology.yc   Ready     Active                          20.10.20
9jjo5hxpk5r4x60ic6dp6zwwf     node06.netology.yc   Ready     Active                          20.10.20
```

## Задача 3

Создать ваш первый, готовый к боевой эксплуатации кластер мониторинга, состоящий из стека микросервисов.

Для получения зачета, вам необходимо предоставить скриншот из терминала (консоли), с выводом команды:
```
docker service ls
```

Ответ:
```
[centos@node01 ~]$ sudo docker service ls
ID             NAME                                MODE         REPLICAS   IMAGE                                          PORTS
cfyvhtk90oxi   swarm_monitoring_alertmanager       replicated   1/1        stefanprodan/swarmprom-alertmanager:v0.14.0
3sz8qkahjj0t   swarm_monitoring_caddy              replicated   1/1        stefanprodan/caddy:latest                      *:3000->3000/tcp, *:9090->9090/tcp, *:9093-9094->9093-9094/tcp
543f9v0k9tb0   swarm_monitoring_cadvisor           global       6/6        google/cadvisor:latest
xe44cgm8es7r   swarm_monitoring_dockerd-exporter   global       6/6        stefanprodan/caddy:latest
s1c67skrd65h   swarm_monitoring_grafana            replicated   1/1        stefanprodan/swarmprom-grafana:5.3.4
qi2wwrh0kx8w   swarm_monitoring_node-exporter      global       6/6        stefanprodan/swarmprom-node-exporter:v0.16.0
tdlu5wepx4l5   swarm_monitoring_prometheus         replicated   1/1        stefanprodan/swarmprom-prometheus:v2.5.0
okt2j8ubm7e8   swarm_monitoring_unsee              replicated   1/1        cloudflare/unsee:v0.8.0
```

## Задача 4 (*)

Выполнить на лидере Docker Swarm кластера команду (указанную ниже) и дать письменное описание её функционала, что она делает и зачем она нужна:
```
# см.документацию: https://docs.docker.com/engine/swarm/swarm_manager_locking/
docker swarm update --autolock=true
```
Ответ:
```
[centos@node01 ~]$ sudo docker swarm update --autolock=true
Swarm updated.
To unlock a swarm manager after it restarts, run the `docker swarm unlock`
command and provide the following key:

    SWMKEY-1-A7bBo5QSfPBFitsraV8SNXI8RVe355lgIz+we1soN1o

Please remember to store this key in a password manager, since without it you
will not be able to restart the manager.
```

Таким образом я залочил TLS ключи для связи между нодами и ключи для расшифрофки RAFT логов. Теперь после перезагрузки докера мне нужно будет ввести выданный мне ключ, чтобы docker swarm начал работать. Это нужно для того, чтобы злоумышленники не похитили ключи, хранящиеся на ноде. 

---

### Как cдавать задание

Выполненное домашнее задание пришлите ссылкой на .md-файл в вашем репозитории.

---
