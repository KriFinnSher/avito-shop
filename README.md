# Тестовое задание Avito (winter 2025)

## Краткое описание сервиса
Сервис avito-shop предназначен для внутреннего использования сотрудниками Avito. Сервис позволяет покупать товары за монеты (coins),
а также отправлять их другим сотрудникам  в знак благодарности или как подарок.

## Инструкция по запуску
1. Клонируйте данный репозиторий:
```
git clone https://github.com/KriFinnSher/avito-shop.git
```
2. В главной директории проекта запустите команду для сборки:
```
docker-compose up
```
Готово! Теперь сервис доступен на порту `:8080`.

## Вопросы, с которыми я столкнулся и логика их решения
### Транзкции (перемещения монеток)
Они бывают двух типов: покупка (purchase) и передача (transfer). Поскольку на логическом уровне оба типа транзакций
практически идентичны по смыслу, было принято решение использовать для обоих единую модель данных `Transaction` со следующей
структурой:
```
type Transaction struct {
	ID     uuid.UUID
	From   string
	Type   string  // "transfer" or "purchase"
	Amount uint64
	To     string  // if "transfer" then it stores reciever's name, otherwise NULL
	Item   string  // if "purchase" then it stores item's name, otherwise NULL
	Date   time.Time
}
```
## Нагрузочное тестирование
В качестве основного инструмента для проведения нагрузочного тестирования был выбран Apache JMeter.
Полученные результаты тестирования отображены на скриншотах ниже:
![image](https://github.com/user-attachments/assets/c322e2ca-f2b6-4ba3-870f-b8a270d08f47)
![image](https://github.com/user-attachments/assets/9f4266a9-f180-4370-a21e-dff5f5afb9e6)


