# jsql
Mini SQL shell working on an in memory DB — for learning how real database shells work, one command at a time.

## Currently supported ##

```shell
Welcome to jsql v0.1
Type \help for help, \exit to quit.

jsql> INSERT INTO users VALUES(1, 'Milica');
1 row inserted.
jsql> INSERT INTO users (VALUES(2, 'Krpotich');
Parse error: expected VALUES, got (
jsql> INSERT INTO users VALUES(2, 'Krmpotich');
1 row inserted.
jsql> SELECT * FROM users;
+--------------------------+
| USERS |
+--------------------------+
1 | Milica
2 | Krmpotich
(2 rows)
jsql> \help
Available commands:
  SQL-like: INSERT INTO table VALUES (...);
             SELECT * FROM table;
  Meta: \help, \exit
jsql> \exit
Goodbye!

```
