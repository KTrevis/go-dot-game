import asyncio
from typing import Any
import websockets
from ...utils import connect, login, read, sendMessage
from .valid_login import validLogin
from .invalid_login import invalidLogin
import json
from colorama import Fore

async def doubleLogin():
    socket = await connect()
    _ = await login(socket)
    _, data = await login(socket)

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
