# Surl Cli Command

## Help
show help message
### Usage
```sh
surl help
```
### Output
```sh
SURL v0.0.1

Commands:
  list       show list of requests
  run <name> run http request by name
...
```
### Flags
- `surl <command> -h` show help for specific command

## List
for list all request on surl.json
### Usage
```sh
surl list
```
### Output
```sh
NAME   METHOD  URL                                         
first  GET     https://jsonplaceholder.typicode.com/todos/2
second POST    https://jsonplaceholder.typicode.com/posts 
```
### Flags
- `-p <path>` overide json path (default: ./surl.json)
- `-h` show help message

## Run
execute http request on surl.json by name
### Usage
```sh
surl run <request name>
```
### Output
```sh
{
    "id": 1,
    "name": "first todo",
    "is_done": true
}
```
### Flags
-  `-p <path>` overide json path (default: ./surl.json)
-  `-v` run verbosely
-  `-h` show help message
