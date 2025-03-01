# How To

## Basic Request
```json
{
    "requests": [
        {
            "name": "request_name",
            "url": "http://127.0.0.1:8000/posts/1",
            "method": "GET",
        }
    ]
}
```

## Send Json Resquest
```json
{
    "requests": [
        {
            "name": "request_name",
            "url": "http://127.0.0.1:8000/posts",
            "method": "POST",
            "json": {
                "name": "foo",
                "max_hp": 100,
                "is_active": true,
                "elements": ["fire", "ice"],
                "info": {
                    "birth_date": "1205-01-30",
                    "birth_place": "egypt",
                }
            }
        }
    ]
}
```
You can put any json valid syntax without "" on key json. Automaticaly set header (content-type: application/json).
You can only set json if body, form and form_multipart is not defined.

## Send Form
```json
{
    "requests": [
        {
            "name": "request_name",
            "url": "http://127.0.0.1:8000/posts",
            "method": "POST",
            "form": [
                {
                    "name": "username",
                    "value": "hello"
                },
                {
                    "name": "password",
                    "value": "some password"
                }
            ]
        }
    ]
}
```
Every object required two key "name" and "value".
Automaticaly set header (content-type: application/x-www-form-urlencoded).
You can only set form if body, json and form_multipart is not defined.

## Send Form Multipart
```json
{
    "requests": [
        {
            "name": "request_name",
            "url": "http://127.0.0.1:8000/posts",
            "method": "POST",
            "form_multipart": [
                {
                    "name": "username",
                    "type": "string",
                    "value": "hello"
                },
                {
                    "name": "file",
                    "type": "file",
                    "file_name": "README.md",
                    "file_path": "./README.md"
                }
            ]
        }
    ]
}
```
There are two type "string" or "file". 
If type equal "string", Key "value" is required.
if type equal "file", Key "file_name" and "file_path" is required.
You can send file using form_multipart.
You can only set form_multipart if body, json and form_multipart is not defined.

## Send Raw Request
```json
{
    "requests": [
        {
            "name": "request_name",
            "url": "http://127.0.0.1:8000/posts",
            "method": "POST",
            "headers": [
                {
                    "key": "content-type",
                    "value": "application/json"
                }
            ],
            "body": "{\"hello\":\"world\"}",
        }
    ]
}
```
## Add timeout
```json
{
    "requests": [
        {
            "name": "request_name",
            "url": "http://127.0.0.1:8000/posts/1",
            "method": "GET",
            "timeout": 1,
        }
    ]
}
```
Timeout defined in seconds. If not defined it means no timeout (wait until get response).
