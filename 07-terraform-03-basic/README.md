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
      + access_key            = "YC<secret>MzVSUC51XLpR"
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
```
vagrant@vagrant:~/07-terraform-03-basic$ terraform workspace select prod
Switched to workspace "prod".
vagrant@vagrant:~/07-terraform-03-basic$ terraform workspace list
  default
* prod
  stage

vagrant@vagrant:~/07-terraform-03-basic$ terraform plan

Terraform used the selected providers to generate the following execution plan. Resource actions are indicated with the following symbols:
  + create

Terraform will perform the following actions:

  # yandex_compute_instance.instance[0] will be created
  + resource "yandex_compute_instance" "instance" {
      + created_at                = (known after apply)
      + folder_id                 = (known after apply)
      + fqdn                      = (known after apply)
      + hostname                  = "-1"
      + id                        = (known after apply)
      + metadata                  = {
          + "ssh-keys" = <<-EOT
                ubuntu:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQDGzPUNN7E4SZLMzBtu9HQ9q3pvqYH7FvLs8qDRTGoW52bwDp/MSmfMuQxCv96ZRphBDzTk/plAWEean3294fbj4uuJTDAV7oIWsyj+Bgae71dukCM/jxZ6N/u2xgkOnxQ5gVDLuu1zpSTCoFRzlbof7mRHextWQzqfAci33Zv0N4031iTS5ZURkdYrwqcy5qrXUIpDc1bTnlA4szSDBcfOdRZJfEGErlhkI4P1FPOu9bKNkL4OaM+tSdnDcv5gFHF9Rfv3wQt23IkSIotKS4k0g9rndiGtJo70Wd1BLt8PIHRdY9x5bHiIm5pBZuGhQn5zCVO6DTpgwszSj/nWJr/IMOxBX2h3jZzt2sKWv4PucVwwc641sH53g2hIyjOOgyxq5R75/IuaEi8K7JVje1QiCoZuYcnNkTog7jzlFZiGRukfbvrx2zupXjsTLjI2kUWwpv/aHg1jB/KWi2FwKJxH5e+Wx9m2/htSK8oC+GDkvnMga7V2ZRpU6XSyY7fbZcs= vagrant@vagrant
            EOT
        }
      + name                      = "vm-1"
      + network_acceleration_type = "standard"
      + platform_id               = "standard-v1"
      + service_account_id        = (known after apply)
      + status                    = (known after apply)
      + zone                      = (known after apply)

      + boot_disk {
          + auto_delete = true
          + device_name = (known after apply)
          + disk_id     = (known after apply)
          + mode        = (known after apply)

          + initialize_params {
              + block_size  = (known after apply)
              + description = (known after apply)
              + image_id    = "fd8haecqq3rn9ch89eua"
              + name        = (known after apply)
              + size        = (known after apply)
              + snapshot_id = (known after apply)
              + type        = "network-hdd"
            }
        }

      + network_interface {
          + index              = (known after apply)
          + ip_address         = (known after apply)
          + ipv4               = true
          + ipv6               = (known after apply)
          + ipv6_address       = (known after apply)
          + mac_address        = (known after apply)
          + nat                = true
          + nat_ip_address     = (known after apply)
          + nat_ip_version     = (known after apply)
          + security_group_ids = (known after apply)
          + subnet_id          = (known after apply)
        }

      + placement_policy {
          + host_affinity_rules = (known after apply)
          + placement_group_id  = (known after apply)
        }

      + resources {
          + core_fraction = 100
          + cores         = 2
          + memory        = 4
        }

      + scheduling_policy {
          + preemptible = (known after apply)
        }
    }

  # yandex_compute_instance.instance[1] will be created
  + resource "yandex_compute_instance" "instance" {
      + created_at                = (known after apply)
      + folder_id                 = (known after apply)
      + fqdn                      = (known after apply)
      + hostname                  = "-2"
      + id                        = (known after apply)
      + metadata                  = {
          + "ssh-keys" = <<-EOT
                ubuntu:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQDGzPUNN7E4SZLMzBtu9HQ9q3pvqYH7FvLs8qDRTGoW52bwDp/MSmfMuQxCv96ZRphBDzTk/plAWEean3294fbj4uuJTDAV7oIWsyj+Bgae71dukCM/jxZ6N/u2xgkOnxQ5gVDLuu1zpSTCoFRzlbof7mRHextWQzqfAci33Zv0N4031iTS5ZURkdYrwqcy5qrXUIpDc1bTnlA4szSDBcfOdRZJfEGErlhkI4P1FPOu9bKNkL4OaM+tSdnDcv5gFHF9Rfv3wQt23IkSIotKS4k0g9rndiGtJo70Wd1BLt8PIHRdY9x5bHiIm5pBZuGhQn5zCVO6DTpgwszSj/nWJr/IMOxBX2h3jZzt2sKWv4PucVwwc641sH53g2hIyjOOgyxq5R75/IuaEi8K7JVje1QiCoZuYcnNkTog7jzlFZiGRukfbvrx2zupXjsTLjI2kUWwpv/aHg1jB/KWi2FwKJxH5e+Wx9m2/htSK8oC+GDkvnMga7V2ZRpU6XSyY7fbZcs= vagrant@vagrant
            EOT
        }
      + name                      = "vm-2"
      + network_acceleration_type = "standard"
      + platform_id               = "standard-v1"
      + service_account_id        = (known after apply)
      + status                    = (known after apply)
      + zone                      = (known after apply)

      + boot_disk {
          + auto_delete = true
          + device_name = (known after apply)
          + disk_id     = (known after apply)
          + mode        = (known after apply)

          + initialize_params {
              + block_size  = (known after apply)
              + description = (known after apply)
              + image_id    = "fd8haecqq3rn9ch89eua"
              + name        = (known after apply)
              + size        = (known after apply)
              + snapshot_id = (known after apply)
              + type        = "network-hdd"
            }
        }

      + network_interface {
          + index              = (known after apply)
          + ip_address         = (known after apply)
          + ipv4               = true
          + ipv6               = (known after apply)
          + ipv6_address       = (known after apply)
          + mac_address        = (known after apply)
          + nat                = true
          + nat_ip_address     = (known after apply)
          + nat_ip_version     = (known after apply)
          + security_group_ids = (known after apply)
          + subnet_id          = (known after apply)
        }

      + placement_policy {
          + host_affinity_rules = (known after apply)
          + placement_group_id  = (known after apply)
        }

      + resources {
          + core_fraction = 100
          + cores         = 2
          + memory        = 4
        }

      + scheduling_policy {
          + preemptible = (known after apply)
        }
    }

  # yandex_vpc_network.network-1 will be created
  + resource "yandex_vpc_network" "network-1" {
      + created_at                = (known after apply)
      + default_security_group_id = (known after apply)
      + folder_id                 = (known after apply)
      + id                        = (known after apply)
      + labels                    = (known after apply)
      + name                      = "network1"
      + subnet_ids                = (known after apply)
    }

  # yandex_vpc_subnet.subnet-1 will be created
  + resource "yandex_vpc_subnet" "subnet-1" {
      + created_at     = (known after apply)
      + folder_id      = (known after apply)
      + id             = (known after apply)
      + labels         = (known after apply)
      + name           = "subnet1"
      + network_id     = (known after apply)
      + v4_cidr_blocks = [
          + "192.168.10.0/24",
        ]
      + v6_cidr_blocks = (known after apply)
      + zone           = "ru-central1-a"
    }

Plan: 4 to add, 0 to change, 0 to destroy.

─────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────

Note: You didn't use the -out option to save this plan, so Terraform can't guarantee to take exactly these actions if you run "terraform apply" now.
```


---

### Как cдавать задание

Выполненное домашнее задание пришлите ссылкой на .md-файл в вашем репозитории.

---
