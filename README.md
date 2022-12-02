## Приложение погоды на API Яндекс.Погоды

#### Координаты места для определения погоды задаются в градусах по широте и долготе


Для настройки необходимо переименовать файл `.test.env` в `.env` в директории `etc` со следующими параметрами:
- token - тип строка (можно получить в личном кабинете разработчика на яндекс аккаунте)
- latitude - тип число, широта
- longitude - тип число, долгота

Приложение выводит результат запроса в консоль со следующими значениями:
- уточнённая широта и долгота территории по которой выданы данные о фактической погоде
- Фактическая температура воздуха, температура воздуха по ощющениям.
- Влажность воздуха.
- Скорость ветра, направление ветра.
- Давление воздуха в паскалях и миллиметрах ртутного столба.
- Описание погоды (солнечно, дождь, гроза и т.п.)