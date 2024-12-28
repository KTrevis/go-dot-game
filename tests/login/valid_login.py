import json
from typing import Any
from websockets.asyncio.client import ClientConnection
from colored_print import log

from tests.utils import sendMessage

async def validLogin(socket: ClientConnection):
    message = {
            "username": "test",
            "password": "test"
            }
    await sendMessage(socket, "LOGIN", message)

    data: dict[str, Any] = json.loads(await socket.recv())
    assert "authenticated" in data
    assert data["authenticated"] == True
    log.success("[VALID LOGIN OK]\n")
