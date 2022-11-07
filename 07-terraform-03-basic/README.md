# Домашнее задание к занятию "7.3. Основы и принцип работы Терраформ"

## Задача 1. Создадим бэкэнд в S3 (необязательно, но крайне желательно).

Если в рамках предыдущего задания у вас уже есть аккаунт AWS, то давайте продолжим знакомство со взаимодействием
терраформа и aws. 

1. Создайте s3 бакет, iam роль и пользователя от которого будет работать терраформ. Можно создать отдельного пользователя,
а можно использовать созданного в рамках предыдущего задания, просто добавьте ему необходимы права, как описано 
[здесь](https://www.terraform.io/docs/backends/types/s3.html).
1. Зарегистрируйте бэкэнд в терраформ проекте как описано по ссылке выше. 

Ответ:
```
vagrant@vagrant:~/07-terraform-03-basic$ cat bucket.tf 
resource "yandex_iam_service_account" "sa" {
  name = "s3admin"
}

// Назначение роли сервисному аккаунту
resource "yandex_resourcemanager_folder_iam_member" "sa-editor" {
  folder_id = "b1<secret>bkm2gt1q"
  role      = "storage.editor"
  member    = "serviceAccount:${yandex_iam_service_account.sa.id}"
}

// Создание статического ключа доступа
resource "yandex_iam_service_account_static_access_key" "sa-static-key" {
  service_account_id = yandex_iam_service_account.sa.id
  description        = "static access key for object storage"
}

// Создание бакета с использованием ключа
resource "yandex_storage_bucket" "test" {
  access_key = yandex_iam_service_account_static_access_key.sa-static-key.access_key
  secret_key = yandex_iam_service_account_static_access_key.sa-static-key.secret_key
  bucket     = "netbucket"
  ```

<details>
<summary>Результаты выполнения terraform plan</summary>

```
vagrant@vagrant:~/07-terraform-03-basic$ terraform apply -auto-approve
yandex_iam_service_account.sa: Refreshing state... [id=aje60ggq35irh085pii6]
yandex_resourcemanager_folder_iam_member.sa-editor: Refreshing state... [id=b1gjbcdp4ij0bkm2gt1q/storage.editor/serviceAccount:aje60ggq35irh085pii6]
yandex_iam_service_account_static_access_key.sa-static-key: Refreshing state... [id=aje274rpqlr4glg59v1g]

Terraform used the selected providers to generate the following execution plan. Resource actions are indicated with the following symbols:
  + create

Terraform will perform the following actions:

  # yandex_storage_bucket.test will be created
  + resource "yandex_storage_bucket" "test" {
      + access_key            = "YCAJE_IqM5H1BMzVSUC51XLpR"
      + acl                   = "private"
      + bucket                = "netbucket"
      + bucket_domain_name    = (known after apply)
      + default_storage_class = (known after apply)
      + folder_id             = (known after apply)
      + force_destroy         = false
      + id                    = (known after apply)
      + secret_key            = (sensitive value)
      + website_domain        = (known after apply)
      + website_endpoint      = (known after apply)

      + anonymous_access_flags {
          + list = (known after apply)
          + read = (known after apply)
        }

      + versioning {
          + enabled = (known after apply)
        }
    }

Plan: 1 to add, 0 to change, 0 to destroy.
yandex_storage_bucket.test: Creating...
yandex_storage_bucket.test: Still creating... [10s elapsed]
yandex_storage_bucket.test: Still creating... [20s elapsed]
yandex_storage_bucket.test: Still creating... [30s elapsed]
yandex_storage_bucket.test: Still creating... [40s elapsed]
yandex_storage_bucket.test: Still creating... [50s elapsed]
yandex_storage_bucket.test: Still creating... [1m0s elapsed]
yandex_storage_bucket.test: Creation complete after 1m2s [id=netbucket]

Apply complete! Resources: 1 added, 0 changed, 0 destroyed.
```
  
</details>

## Задача 2. Инициализируем проект и создаем воркспейсы. 

1. Выполните `terraform init`:
    * если был создан бэкэнд в S3, то терраформ создат файл стейтов в S3 и запись в таблице 
dynamodb.
    * иначе будет создан локальный файл со стейтами.  
1. Создайте два воркспейса `stage` и `prod`.
1. В уже созданный `aws_instance` добавьте зависимость типа инстанса от вокспейса, что бы в разных ворскспейсах 
использовались разные `instance_type`.
1. Добавим `count`. Для `stage` должен создаться один экземпляр `ec2`, а для `prod` два. 
1. Создайте рядом еще один `aws_instance`, но теперь определите их количество при помощи `for_each`, а не `count`.
1. Что бы при изменении типа инстанса не возникло ситуации, когда не будет ни одного инстанса добавьте параметр
жизненного цикла `create_before_destroy = true` в один из рессурсов `aws_instance`.
1. При желании поэкспериментируйте с другими параметрами и рессурсами.

В виде результата работы пришлите:
* Вывод команды `terraform workspace list`.
* Вывод команды `terraform plan` для воркспейса `prod`.  

Ответ:


---

### Как cдавать задание

Выполненное домашнее задание пришлите ссылкой на .md-файл в вашем репозитории.

---
