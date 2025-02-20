# Домашнее задание к занятию "7.5. Основы golang"

С `golang` в рамках курса, мы будем работать не много, поэтому можно использовать любой IDE. 
Но рекомендуем ознакомиться с [GoLand](https://www.jetbrains.com/ru-ru/go/).  

## Задача 1. Установите golang.
1. Воспользуйтесь инструкций с официального сайта: [https://golang.org/](https://golang.org/).
2. Так же для тестирования кода можно использовать песочницу: [https://play.golang.org/](https://play.golang.org/).

Ответ:
```
ramires@Romans-MacBook-Air virt-homeworks % brew info go
==> go: stable 1.19.3 (bottled), HEAD
Open source programming language to build simple/reliable/efficient software
https://go.dev/
/opt/homebrew/Cellar/go/1.19.3 (12,444 files, 629MB) *
  Poured from bottle on 2022-11-12 at 14:23:02
```

## Задача 2. Знакомство с gotour.
У Golang есть обучающая интерактивная консоль [https://tour.golang.org/](https://tour.golang.org/). 
Рекомендуется изучить максимальное количество примеров. В консоли уже написан необходимый код, 
осталось только с ним ознакомиться и поэкспериментировать как написано в инструкции в левой части экрана.  

Ответ:
Обучение пройдено.

## Задача 3. Написание кода. 
Цель этого задания закрепить знания о базовом синтаксисе языка. Можно использовать редактор кода 
на своем компьютере, либо использовать песочницу: [https://play.golang.org/](https://play.golang.org/).

1. Напишите программу для перевода метров в футы (1 фут = 0.3048 метр). Можно запросить исходные данные 
у пользователя, а можно статически задать в коде.
    Для взаимодействия с пользователем можно использовать функцию `Scanf`:
    ```
    package main
    
    import "fmt"
    
    func main() {
        fmt.Print("Enter a number: ")
        var input float64
        fmt.Scanf("%f", &input)
    
        output := input * 2
    
        fmt.Println(output)    
    }
    ```
 
2. Напишите программу, которая найдет наименьший элемент в любом заданном списке, например:
    ```
    x := []int{48,96,86,68,57,82,63,70,37,34,83,27,19,97,9,17,}
    ```
3. Напишите программу, которая выводит числа от 1 до 100, которые делятся на 3. То есть `(3, 6, 9, …)`.

В виде решения ссылку на код или сам код.

Ответ:
1. Вариант решения с вводом значения [вот тут](https://github.com/RamiresHab/virt-homeworks/blob/master/07-terraform-05-golang/HW3/hw3_1a.go). Вариант с заданием значения переменной в коде [вот](https://github.com/RamiresHab/virt-homeworks/blob/master/07-terraform-05-golang/HW3/hw3_1b.go)
2. Решение задачи [тут](https://github.com/RamiresHab/virt-homeworks/blob/master/07-terraform-05-golang/HW3/hw3_3.go)
3. Решение задачи [тут](https://github.com/RamiresHab/virt-homeworks/blob/master/07-terraform-05-golang/HW3/hw3_2.go)

## Задача 4. Протестировать код (не обязательно).

Создайте тесты для функций из предыдущего задания. 

---

### Как cдавать задание

Выполненное домашнее задание пришлите ссылкой на .md-файл в вашем репозитории.

---

