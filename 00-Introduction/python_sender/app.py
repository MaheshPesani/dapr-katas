import time
import requests
import os

n = 0
while True:
    n = (n + 1) % 1000000
    message = {"message": "Hello, World" + str(n)}
    try:
        resp = requests.post(
            """http://localhost:3500/v1.0/invoke/nodereceiver/method/greeting""",
            json=message)
    except Exception as e:
        print(e)
    time.sleep(5)