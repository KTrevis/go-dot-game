import json
from typing import Any
from websockets.asyncio.client import ClientConnection

async def sendMessage(socket: ClientConnection, type: str, data: dict[str, Any]):
    await socket.send(json.dumps(type))
    await socket.send(json.dumps(data))
