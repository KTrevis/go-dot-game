import asyncio
import json
from typing import Any
import websockets
from websockets.asyncio.client import ClientConnection
from ...utils import connect, login, read, sendMessage
from colorama import Fore

async def invalidType():
    socket = await connect()
    message = {
            "username": "test",
            "password": "sqdfllkj"
            }
    await sendMessage(socket, "CACA", message)

    try:
        _ = await read(socket)
        raise
    except: pass

    print(f"{Fore.GREEN}[INVALID TYPE OK]\n")
    await socket.close()

async def invalidData():
    socket = await connect()
    await sendMessage(socket, "LOGIN", "lkjqsdfqs")

    msgType, data = await read(socket)

    assert "LOGIN" in msgType
    assert "error" in data
    print(f"{Fore.GREEN}[INVALID DATA OK]\n")
    await socket.close()

async def invalidPassword():
    socket = await connect()
    await sendMessage(socket, "LOGIN", {
        "username": "test", "password": "qsdkjfhq"
        })
    msgType, data = await read(socket)

    assert "error" in data
    assert msgType == "LOGIN"
    print(f"{Fore.GREEN}[INVALID PASSWORD OK]\n")
    await socket.close()

async def invalidUsername():
    socket = await websockets.connect("ws://localhost:8080/websocket")
    message = {
            "username": "sqdfkhq",
            "password": "test"
            }
    await sendMessage(socket, "LOGIN", message)

    msgType, data = await read(socket)

    assert "error" in data
    assert msgType == "LOGIN"
    print(f"{Fore.GREEN}[INVALID USERNAME OK]\n")
    await socket.close()

async def invalidLogin():
    _ = await asyncio.gather(
            invalidPassword(),
            invalidUsername(),
            invalidData(),
            invalidType(),
            )
