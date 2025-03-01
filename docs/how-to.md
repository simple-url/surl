# How To

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
You can put any json valid syntax without "" on key json. Automaticaly set json header (content-type: application/json).
You can only set json if body, form and form_multipart is not defined.


## Send Raw Json Resquest
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
