import json
from typing import Any
import websockets
from websockets.asyncio.client import ClientConnection
import requests

async def register(username: str, password: str) -> bool:
    res = requests.post("http://localhost:8080/api/register", {"username": username, "password": password, "confirmPassword": password})
    return res.status_code == 200

async def login(socket: ClientConnection, username: str = "test", password: str = "test") -> tuple[str, dict[str, Any]]:
    await sendMessage(socket, "LOGIN", {"username": username, "password": password})

    msgType, data = await read(socket)
    assert msgType == "LOGIN"
    return msgType, data

async def sendMessage(socket: ClientConnection, type: str, data: dict[str, Any]):
    await socket.send(f"{type}\r\n{json.dumps(data)}")

async def read(socket: ClientConnection) -> tuple[str, dict[str, Any]]:
    msg = str(await socket.recv())
    split: list[str] = msg.split("\r\n")
    data: dict[str, Any] = json.loads(split[1])
    return split[0], data

async def connect() -> ClientConnection:
    return await websockets.connect("ws://localhost:8080/websocket")
