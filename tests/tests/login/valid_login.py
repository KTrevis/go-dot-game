import json
from typing import Any
import websockets
from colorama import Fore

from ...utils import connect, login, sendMessage

async def loginDifferentSocket():
    socket = await connect()
    msgType, data = await login(socket)

    assert msgType == "LOGIN"
    assert "error" in data
    print(f"{Fore.GREEN}[LOGIN TO SAME ACCOUNT DIFFERENT SOCKET OK]\n")
    await socket.close()

async def validLogin():
    socket = await connect()
    msgType, data = await login(socket)

    assert msgType == "LOGIN"
    assert "authenticated" in data
    print(f"{Fore.GREEN}[VALID LOGIN OK]\n")
    await loginDifferentSocket()
    await socket.close()
