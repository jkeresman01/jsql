# jsql
Mini SQL shell working on an in memory DB — for learning how real database shells work, one command at a time.

## Currently supported ##

```shell
jsql> CREATE DATABASE shop;
Database 'shop' created.

jsql> \connect shop
Connected to database 'shop'.

jsql> INSERT INTO users VALUES (1, 'Milica');
1 row inserted.

jsql> \disconnect
Disconnected from database 'shop'.

jsql> SELECT * FROM users;
Error: no database selected.

jsql> \connect shop
Connected to database 'shop'.

jsql(shop)> SELECT * FROM users;
+-------+--------+
| COL1  | COL2   |
+-------+--------+
| 1     | Milica |
+-------+--------+
(1 row)

```
