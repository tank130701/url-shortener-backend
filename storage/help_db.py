import sqlite3
from tabulate import tabulate

# Подключение к базе данных SQLite
conn = sqlite3.connect('storage.db')
cursor = conn.cursor()

# Выполнение SQL-запроса для выборки всех записей из таблицы
cursor.execute("SELECT * FROM urls")

# Получение результатов запроса
data = cursor.fetchall()

# Получение заголовков столбцов (имен полей)
column_names = [description[0] for description in cursor.description]

# Используйте tabulate для вывода данных в виде таблицы
table = tabulate(data, headers=column_names, tablefmt='grid')

# Вывод таблицы на экран
print(table)

# Закрытие соединения с базой данных
conn.close()
