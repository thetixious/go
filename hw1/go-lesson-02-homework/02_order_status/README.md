# Задача 2. Анализ статусов заказов

**Самостоятельная работа · базовая сложность**

Реализуйте `CountStatuses(statuses []string) (newCount, paidCount, shippedCount, cancelledCount, unknownCount, activeCount int)`.

Статусы `new`, `paid`, `shipped`, `cancelled` учитываются отдельно. Все остальные строки (включая пустую) считаются неизвестными. **Сравнение регистрозависимое**. Активные заказы — `new`, `paid` и `shipped`.

Используйте `for range` и `switch`. Пустой список должен возвращать шесть нулей.

Пример: `[]string{"new","paid","shipped","cancelled","paid","unknown","new"}` → `(2,2,1,1,1,5)`.
