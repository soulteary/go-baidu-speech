import requests
headers = {
    'Content-Type': 'application/x-www-form-urlencoded',
}
data = 'text=阳光彩虹小白马'.encode()
response = requests.post('http://127.0.0.1:8080/tts', headers=headers, data=data)
print(response.text)