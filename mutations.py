import requests
import json

# Definir la URL del servidor GraphQL
url = 'http://localhost:8080'

# Solicitar al usuario que elija qué tipo de objeto agregar
object_type = input("¿Qué desea agregar? (refrigerador, televisión, celular): ")

# Definir la mutación y los datos requeridos de acuerdo al tipo de objeto elegido
if object_type == "Refrigerador":
    mutation_name = "addRefrigerator"
    mutation_input = {
        "name": input("Nombre del refrigerador: "),
        "brand": input("Marca del refrigerador: "),
        "color": input("Color del refrigerador: "),
        "category": "REFRIGERATOR"
    }
elif object_type == "Television":
    mutation_name = "addTV"
    mutation_input = {
        "name": input("Nombre de la televisión: "),
        "brand": input("Marca de la televisión: "),
        "color": input("Color de la televisión: "),
        "category": "TV"
    }
elif object_type == "Celular":
    mutation_name = "addCellphone"
    mutation_input = {
        "name": input("Nombre del celular: "),
        "brand": input("Marca del celular: "),
        "color": input("Color del celular: "),
        "category": "CELLPHONE"
    }
else:
    print("Tipo de objeto no válido.")
    exit()

# Definir el cuerpo de la solicitud GraphQL
query = '''
mutation AddObject($input: %sMutation!) {
  %s(input: $input) {
    id
    name
    brand
    color
    category
  }
}
''' % (mutation_name[0].upper() + mutation_name[1:], mutation_name)

variables = {
    "input": mutation_input
}

data = {
    'query': query,
    'variables': variables
}

# Establecer los encabezados de la solicitud
headers = {
    'Content-Type': 'application/json'
}

# Realizar la solicitud POST
response = requests.post(url, headers=headers, json=data)

# Mostrar la respuesta
print(response.json())
