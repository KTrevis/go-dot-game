import json
from typing import Any
from websockets.asyncio.client import ClientConnection

async def login(socket: ClientConnection):
    await sendMessage(socket, "LOGIN", {"username": "test", "password": "test"})

    data: dict[str, Any] = json.loads(await socket.recv())
    assert "authenticated" in data
    assert data["authenticated"] == True

async def sendMessage(socket: ClientConnection, type: str, data: dict[str, Any]):
    message = type + "\r\n"
    message += json.dumps(data)
    await socket.send(message)

async def sendMessageAndRead(socket: ClientConnection, type: str, data: dict[str, Any]):
    await sendMessage(socket, type, data)
    data = json.loads(await socket.recv())
    return data
