import asyncio
import websockets

from .login.test import testLogin

async def main():
    socket = await websockets.connect("ws://localhost:8080/websocket")
    await testLogin(socket)

asyncio.run(main())
