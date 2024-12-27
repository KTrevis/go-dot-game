from typing import Any
from .valid_login import validLogin
from .invalid_login import invalidLogin
from websockets.asyncio.client import ClientConnection
import json
from colored_print import log

async def doubleLogin(socket: ClientConnection):
    print("[DOUBLE LOGIN]")
    msg = {
        "username": "test",
        "password": "test"
            }
    await socket.send(json.dumps("LOGIN"))
    await socket.send(json.dumps(msg))

    data: dict[str, Any] = json.loads(await socket.recv())
    log.success("[DOUBLE LOGIN OK]\n")

async def testLogin(socket: ClientConnection):
    print("[LOGIN TESTS]\n")
    await invalidLogin(socket)
    await validLogin(socket)
    await doubleLogin(socket)
