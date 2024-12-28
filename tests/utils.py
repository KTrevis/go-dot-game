import json
from typing import Any
from websockets.asyncio.client import ClientConnection

async def sendMessage(socket: ClientConnection, type: str, data: dict[str, Any]):
    message = type + "\r\n"
    message += json.dumps(data)
    await socket.send(message)
