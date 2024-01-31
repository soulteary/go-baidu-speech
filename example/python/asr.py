import requests
files = {
    'file': open('1706690806.wav', 'rb'),
}
response = requests.post('http://127.0.0.1:8080/asr', files=files)
print(response.text)