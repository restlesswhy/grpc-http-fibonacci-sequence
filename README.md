# Последовательность Фибоначчи

### Требования

Реализовать сервис, возвращающий срез последовательности чисел из ряда Фибоначчи.

Сервис должен отвечать на запросы и возвращать ответ. В ответе должны быть перечислены все числа, последовательности Фибоначчи с порядковыми номерами от x до y.

1. Требуется реализовать два протокола: HTTP REST и GRPC
2. Код должен быть выложен в репозиторий с возможность предоставления доступа (например github.com, bitbucker.org, gitlab.com). Решение предоставить ссылкой на этот репозиторий.
3. Необходимо продумать и описать в readme развертку сервиса на другом компьютер
4. (Опционально) Кэширование. Сервис не должен повторно вычислять числа из ряда Фибоначчи. Значения необходимо сохранить в Redis или Memcache.
5. (Опционально) Код должен быть покрыт тестами.


### Для запуска приложения:

```
make run
```

### Unit-тесты

```
make test
```

### Интеграционные тесты 

```
make test.integration
```

Опция кэширования настраивается в конфиге.
