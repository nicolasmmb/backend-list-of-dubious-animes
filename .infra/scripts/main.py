# http request call crate user with body


import requests
from httpx import AsyncClient

url = "http://localhost:8899/user"


async_request_pool = []

for i in range(1000000, 1010000):
    payload = {
        "name": f"NINO.{str(i)}",
        "email": f"NINO.{str(i)}@mail.com",
        "password": "1234",
        "avatar": "https://i.pinimg.com/564x/66/2b/c2/662bc2a219981b8c8c28c779e372aea2.jpg",
    }
    async_request_pool.append(payload)


async def main():
    async with AsyncClient() as client:
        for idx, payload in enumerate(async_request_pool):
            response = await client.post(url, json=payload)
            print(
                f"Request {idx}, status code: {response.status_code} ",
                response.json().get("id"),
            )


if __name__ == "__main__":
    import time

    start = time.time()
    import asyncio

    asyncio.run(main())
    print(f"Time: {time.time() - start}")
