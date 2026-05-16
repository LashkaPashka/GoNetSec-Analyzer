# 🛡️ GoNetSec Analyzer (Network Security Syslog Analyzer)

**GoNetSec Analyzer** — это легковесный, высокопроизводительный Syslog-сервер, написанный на Golang. Он предназначен для получения, парсинга и анализа логов с сетевого оборудования (Cisco, UserGate и др.) в реальном времени.

Проект создан как программное решение задач, изучаемых в рамках сертификации **CCNA Security**. Он автоматизирует обнаружение угроз (MAC-спуфинг, атаки на протоколы маршрутизации, Brute-force), превращая сырые логи роутеров в моментальные алерты.

## 🚀 Ключевые возможности

* **Высокая производительность:** Способен обрабатывать "шторм логов" при сетевых авариях без блокировки (non-blocking I/O).
* **Многопоточная архитектура (Pipeline):** Чтение сети, парсинг и отправка уведомлений разнесены по разным горутинам и общаются через буферизированные каналы.
* **Распознавание угроз "из коробки" (Security Signatures):**
  * 🛑 **L2 Security:** Обнаружение подмены MAC-адресов (`PORT_SECURITY-2-PSECURE_VIOLATION`).
  * 🛑 **L3 Security:** Мониторинг падения OSPF-соседства (`OSPF-5-ADJCHG`).
  * 🛑 **AAA / Management:** Детектирование попыток подбора паролей по SSH/Telnet (`SEC_LOGIN-4-LOGIN_FAILED`).
* **Асинхронные уведомления:** Вывод форматированных алертов в консоль и отправка сообщения в Telegram Bot.

## 🧠 Архитектура приложения

Проект построен по стандарту `Standard Go Project Layout` с жестким разделением обязанностей (Separation of Concerns).

```text
[Cisco Router] --(UDP/514)--> [Receiver] --> (RawLogChan) --> [Parser] --> (AlertChan) --> [Notifier]
```
1. **Parser (Бизнес-логика):** Читает канал с сырыми логами, применяет правила фильтрации (Regexp/Strings) и определяет, является ли лог угрозой.
2. **Notifier (Оповещения):** Читает канал с угрозами и занимается доставкой алертов администратору.

## 🛠️ Запуск проекта

### 1. Клонирование репозитория
```bash
git clone git@github.com:LashkaPashka/GoNetSec-Analyzer.git
cd GoNetSec-Analyzer
```

### 2. Настройка окружения
В файле `.env` в корне проекта укажите соотетствующие параметры:
```bash
UDPPort="5514"
BufferSize="2048"
TelegramToken=""
ChatID=""
WorkersNumber="5"
```

### 3. Запуск
```bash
go run cmd/main.go
```
*Сервер запустится и начнет слушать входящие UDP-пакеты.*

## 🧪 Тестирование (Имитация сетевых атак)

Мы можем сымитировать отправку логов с помощью утилиты `netcat` (`nc`) из другого терминала.

**1. Симуляция падения OSPF (L3 Угроза):**
```bash
echo "*Feb 17 21:24:45.309: %OSPF-5-ADJCHG: Process 1, Nbr 10.2.2.2 on Serial0/0/0 from FULL to DOWN" | nc -u -w1 127.0.0.1 5514
```

**2. Симуляция нарушения Port Security (L2 Угроза / MAC-спуфинг):**
```bash
echo "*Jan 14 01:34:39.750: %PORT_SECURITY-2-PSECURE_VIOLATION: Security violation occurred, caused by MAC address aaaa.bbbb.cccc on port FastEthernet0/5." | nc -u -w1 127.0.0.1 5514
```

**3. Симуляция Brute-force атаки (AAA):**
```bash
echo "*Mar  1 00:15:32.123: %SEC_LOGIN-4-LOGIN_FAILED: Login failed [user: admin] [Source: 192.168.1.100] [localport: 22]" | nc -u -w1 127.0.0.1 5514
```
