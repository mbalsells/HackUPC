# Make sure you have installed the required packages:
#   pip install requests
'''
import requests

url = 'https://api-us.restb.ai/vision/v2/predict'
payload = {
    # Add your client key
    'client_key': '3ad1b09d045d5efa10fa715c690247d43f378113b6f300f9c07d74616d014b17',
    'model_id': 'real_estate_global_v2',
    # Add the image URL you want to classify
    'image_url': 'https://images-na.ssl-images-amazon.com/images/I/91SaGONDneL._SX355_.jpg'
}

# Make the classify request
response = requests.get(url, params=payload)
print(response)
# The response is formatted in JSON
json_response = response.json()
print(json_response)
'''


import requests

url = 'http://localhost:8080/sampleEndpoint'
payload = {
    'name': 'Javier',
    'email': 'javi@gmail.com',
    'username': 'javierlc2000',
    'key': 'pasnataga'
}


# Make the classify request
response = requests.get(url, params=payload)
print(response)
# The response is formatted in JSON
json_response = response.json()
print(json_response)