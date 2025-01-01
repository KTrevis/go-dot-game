import colorama
from ...utils import connect, login, read, register, sendMessage


async def teleportTest():
    username = "teleport"
    password = "qsdjkfqskkfsq"

    socket = await connect()
    _ = await register(username, password)
    _ = await login(socket, username, password)
    await sendMessage(socket, "CREATE_CHARACTER", {"name": username, "class": "Mage"})
    await sendMessage(socket, "ENTER_WORLD", {"character": username})
    await sendMessage(socket, "UPDATE_PLAYER_POSITION", {"position": {"x": 420, "y": 69}})

    try:
        _ = await read(socket)
        raise
    except: pass
    await socket.close()
    print(f"{colorama.Fore.GREEN}[TELEPORT TEST OK]\n")
