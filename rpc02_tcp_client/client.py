import socket
import json

request = {
    "id": 0,
    "params": [{"x":10, "y":20}],
    "method": "ServiceA.Add"
}

client = socket.create_connection(("localhost", 9091), 5)
client.sendall(json.dumps(request).encode("utf-8"))

response = client.recv(1024)
print(response.decode("utf-8"))
print(response)