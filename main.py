# http request call crate user with body


import requests

url = "http://localhost:8899/user"


for i in range(160000, 200000):
    payload = {
        "name": f"NINO.{str(i)}",
        "email": f"NINO.{str(i)}@mail.com",
        "password": "1234",
        "avatar": "https://i.pinimg.com/564x/66/2b/c2/662bc2a219981b8c8c28c779e372aea2.jpg",
    }

    response = requests.request("POST", url, json=payload)
