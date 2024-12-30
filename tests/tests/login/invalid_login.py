import asyncio
import json
from typing import Any
import websockets
from websockets.asyncio.client import ClientConnection
from ...utils import sendMessage
from colorama import Fore

async def invalidType():
    socket = await websockets.connect("ws://localhost:8080/websocket")
    message = {
            "username": "test",
            "password": "sqdfllkj"
            }
    await sendMessage(socket, "CACA", message)

    try:
        data: dict[str, Any] = json.loads(await socket.recv())
        raise
    except: pass

    print(f"{Fore.GREEN}[INVALID TYPE OK]\n")
    await socket.close()

async def invalidData():
    socket = await websockets.connect("ws://localhost:8080/websocket")
    await sendMessage(socket, "LOGIN", "lkjqsdfqs")

    data: dict[str, Any] = json.loads(await socket.recv())

    assert "error" in data
    print(f"{Fore.GREEN}[INVALID DATA OK]\n")
    await socket.close()

async def invalidPassword():
    socket = await websockets.connect("ws://localhost:8080/websocket")
    message = {
            "username": "test",
            "password": "sqdfllkj"
            }
    await sendMessage(socket, "LOGIN", message)

    data: dict[str, Any] = json.loads(await socket.recv())

    assert "error" in data
    print(f"{Fore.GREEN}[INVALID PASSWORD OK]\n")
    await socket.close()

async def invalidUsername():
    socket = await websockets.connect("ws://localhost:8080/websocket")
    message = {
            "username": "sqdfkhq",
            "password": "test"
            }
    await sendMessage(socket, "LOGIN", message)

    data: dict[str, Any] = json.loads(await socket.recv())

    assert "error" in data
    print(f"{Fore.GREEN}[INVALID USERNAME OK]\n")
    await socket.close()

async def invalidLogin():
    _ = await asyncio.gather(
            invalidPassword(),
            invalidUsername(),
            invalidData(),
            invalidType(),
            )
