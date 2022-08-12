# Mycoll API

[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](https://go.dev/)

This is my pet project. Main idea is create knowledge database with links. This is like [pocket](https://www.mozilla.org/en-US/firefox/pocket/).
Usage packages:

1. http://github.com/gorilla/mux
2. http://github.com/sirupsen/logrus
3. https://pkg.go.dev/golang.org/x/crypto/bcrypt
4. ![MongoDB](https://img.shields.io/badge/MongoDB-%234ea94b.svg?style=for-the-badge&logo=mongodb&logoColor=white) https://github.com/mongodb/mongo-go-driver
5. ![JWT](https://img.shields.io/badge/JWT-black?style=for-the-badge&logo=JSON%20web%20tokens) https://github.com/golang-jwt/jwt
6. https://github.com/BurntSushi/toml

# How to install

1. Clone this repo: `$ git clone https://github.com/arimatakao/mycoll-api.git`
2. Create config.toml file in `/configs` directory. Use teplate_config.toml for create your own config.
3. Execute make: `$ make`
4. Open: `http://YOURIP:YOURPORT/api/v1`

# Docs

| Method     | Route               | Request body                                                 | Example response                                             | Error Response                                               |
| ---------- | ------------------- | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
|            | /api/v1             |                                                              | {<br/>    "api_name": "mycoll",<br/>    "version": "v1",<br/>    "count_users": 123,<br/>    "count_links": 123<br/>} | 1. Always when request body is wrong<br />{<br />"error": "Wrong json"<br />}<br />2. Always when error in server side<br />{<br />"error": "Something wrong"<br />} |
| **POST**   | /api/v1/signup      | {<br/>    "name": "yourname",<br/>    "password": "yourpassword"<br/>} | {<br/>    "message": "Created new user"<br/>}                | {<br/>    "error": "User is already exist"<br/>}             |
| **POST**   | /api/v1/signin      | {<br/>    "name": "yourname",<br/>    "password": "yourpassword"<br/>} | {<br/>    "accessToken": "adc.adc.adc"<br/>}                 | {<br />"error": "Wrong username or password"<br />}          |
| **DELETE** | /api/v1/deleteme    |                                                              | {<br/>    "count_deleted_links": 123,<br/>    "message": "Success"<br/>} |                                                              |
| **POST**   | /api/v1/createlink  | {<br/>    "title": "YourTitle",<br/>    "tags": [<br/>        "tag1",<br/>        "tag2"<br/>    ],<br/>    "description": "YourDescription",<br/>    "links": [<br/>        "test.com",<br/>        "example.com"<br/>    ]<br/>} | {<br/>    "message": "123abc"<br/>}                          |                                                              |
| **POST**   | /api/v1/findlinks   |                                                              | {<br/>    "links": [<br/>        {<br/>            "_id": "62da8938f370fd0fd9075381",<br/>            "id_owner": "62da8920f370fd0fd9075380",<br/>            "title": "YourTitle",<br/>            "tags": [<br/>                "tag1",<br/>                "tag2"<br/>            ],<br/>            "description": "TestDescription",<br/>            "links": [<br/>                "test.com",<br/>                "example.com"<br/>            ]<br/>        }<br/>    ]<br/>} | {<br/>    "links": null<br />}                               |
| **PUT**    | /api/v1/updatelinks | {<br/>    "_id": "123abc",<br/>    "title": "YourTitle1",<br/>    "tags": [<br/>        "tag11111",<br/>        "tag22222"<br/>    ],<br/>    "description": "TestDescription111111",<br/>    "links": [<br/>        "test.com",<br/>        "google.com"<br/>    ]<br/>} | {<br/>    "count_updated": 1,<br/>    "message": "Update success"<br/>} |                                                              |
| **DELETE** | /api/v1/deletelinks | {<br/>    "_id": "123abc",<br/>}                             | {<br/>    "count_deleted": 1<br />}                          |                                                              |

**When you get accessToken in /api/v1/signin save him and put to authorization header:** `Authorization: Bearer <accessToken>`
