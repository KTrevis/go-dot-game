import json
from typing import Any
import websockets
from colorama import Fore

from ...utils import sendMessage

async def loginDifferentSocket():
    socket = await websockets.connect("ws://localhost:8080/websocket")
    msg = {
        "username": "test",
        "password": "test"
            }
    await sendMessage(socket, "LOGIN", msg)
    data: dict[str, Any] = json.loads(await socket.recv())

    assert "error" in data
    assert data["error"] == "this account is already logged in"
    print(f"{Fore.GREEN}[LOGIN TO SAME ACCOUNT DIFFERENT SOCKET OK]\n")
    await socket.close()

async def validLogin():
    socket = await websockets.connect("ws://localhost:8080/websocket")
    message = {
            "username": "test",
            "password": "test"
            }
    await sendMessage(socket, "LOGIN", message)

    data: dict[str, Any] = json.loads(await socket.recv())
    assert "authenticated" in data
    assert data["authenticated"] == True
    print(f"{Fore.GREEN}[VALID LOGIN OK]\n")
    await loginDifferentSocket()
    await socket.close()
