import asyncio
from typing import Any
import websockets
from ...utils import sendMessage
from .valid_login import validLogin
from .invalid_login import invalidLogin
import json
from colorama import Fore

async def doubleLogin():
    socket = await websockets.connect("ws://localhost:8080/websocket")
    msg = {
        "username": "test",
        "password": "test"
            }
    await sendMessage(socket, "LOGIN", msg)
    await sendMessage(socket, "LOGIN", msg)

    data: dict[str, Any] = json.loads(await socket.recv())
    data = json.loads(await socket.recv())
    assert "error" in data
    assert data["error"] == "you are already authenticated"
    await socket.close()
    print(f"{Fore.GREEN}[DOUBLE LOGIN OK]\n")

async def testLogin():
    _ = await validLogin(),
    _ = await asyncio.gather(
            invalidLogin(),
            doubleLogin()
            )
