import json
from typing import Any
from websockets.asyncio.client import ClientConnection

async def invalidPassword(socket: ClientConnection):
    print("[INVALID PASSWORD]")
    message = {
            "username": "test",
            "password": "sqdfllkj"
            }
    await socket.send(json.dumps("LOGIN"))
    await socket.send(json.dumps(message))

    data: dict[str, Any] = json.loads(await socket.recv())

    assert "error" in data

async def invalidUsername(socket: ClientConnection):
    print("[INVALID USERNAME]")
    message = {
            "username": "sqdfkhq",
            "password": "test"
            }
    await socket.send(json.dumps("LOGIN"))
    await socket.send(json.dumps(message))

    data: dict[str, Any] = json.loads(await socket.recv())

    assert "error" in data

async def invalidLogin(socket: ClientConnection):
    await invalidPassword(socket)
    await invalidUsername(socket)
