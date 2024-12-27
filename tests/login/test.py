from .valid_login import validLogin
from .invalid_login import invalidLogin
from websockets.asyncio.client import ClientConnection

async def testLogin(socket: ClientConnection):
    print("[LOGIN TESTS]")
    await invalidLogin(socket)
    await validLogin(socket)
