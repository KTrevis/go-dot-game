import json
from typing import Any
from websockets.asyncio.client import ClientConnection
from colored_print import log

async def validLogin(socket: ClientConnection):
    print("[VALID LOGIN]")
    message = {
            "username": "test",
            "password": "test"
            }
    await socket.send(json.dumps("LOGIN"))
    await socket.send(json.dumps(message))

    data: dict[str, Any] = json.loads(await socket.recv())
    assert "authenticated" in data
    assert data["authenticated"] == True
    log.success("[VALID LOGIN OK]")
