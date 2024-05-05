# http request call crate user with body


import asyncio
import logging
import time
import uuid

from httpx import AsyncClient

url = "http://localhost:8899/user"
logger = logging.getLogger(__name__)


def generate_payload(total: int):
    payload_tasks = []
    for _ in range(total):
        _id = str(uuid.uuid4())
        payload = {
            "name": f"NINO.{_id}",
            "email": f"NINO.{_id}@mail.com",
            "password": "1234",
            "avatar": "https://i.pinimg.com/564x/66/2b/c2/662bc2a219981b8c8c28c779e372aea2.jpg",
        }
        payload_tasks.append(payload)
    return payload_tasks


async def send_request(client: AsyncClient, payload: dict):
    try:
        response = await client.post(url, json=payload)
        x_trace_id = response.headers.get("x-trace-id")
        msg = f"Trace ID: {x_trace_id} - Request ID: {response.json().get('id')} - Status: {response.status_code}"
        logger.warning(msg)
    except Exception as e:
        print(e)


async def send_request_async(payload: list[dict]):
    all_tasks = []
    async with AsyncClient() as client:
        all_tasks.extend(send_request(client, payload) for payload in payload)
        await asyncio.gather(*all_tasks)


if __name__ == "__main__":
    start = time.time()

    for idx in range(122):
        print(f"Request: {idx}")
        asyncio.run(send_request_async(generate_payload(1000)))
        time.sleep(2)

    print(f"Time: {time.time() - start}")
