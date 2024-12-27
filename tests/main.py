import asyncio
import websockets
from colored_print import log

from .login.test import testLogin

async def main():
    socket = await websockets.connect("ws://localhost:8080/websocket")
    await testLogin(socket)
    log.success("[ALL TESTS OK]\n")

asyncio.run(main())
