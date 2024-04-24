import requests
import json

# Definir la URL del servidor GraphQL
url = 'http://localhost:8080'

def menu():
    # Solicitar al usuario qué acción desea realizar
    action = input("¿Qué desea hacer? (agregar, editar, eliminar, listar, salir): ")

    if action == "agregar":
        # Solicitar al usuario que elija qué tipo de objeto agregar
        object_type = input("¿Qué desea agregar? (refrigerador, televisión, celular): ")

        # Definir la mutación y los datos requeridos de acuerdo al tipo de objeto elegido
        if object_type == "refrigerador":
            mutation_name = "addRefrigerator"
            mutation_input = {
                "name": input("Nombre del refrigerador: "),
                "brand": input("Marca del refrigerador: "),
                "color": input("Color del refrigerador: "),
                "category": "REFRIGERATOR"
            }
        elif object_type == "televisión":
            mutation_name = "addTV"
            mutation_input = {
                "name": input("Nombre de la televisión: "),
                "brand": input("Marca de la televisión: "),
                "color": input("Color de la televisión: "),
                "category": "TV"
            }
        elif object_type == "celular":
            mutation_name = "addCellphone"
            mutation_input = {
                "name": input("Nombre del celular: "),
                "brand": input("Marca del celular: "),
                "color": input("Color del celular: "),
                "category": "CELLPHONE"
            }
        else:
            print("Tipo de objeto no válido.")
            menu()

    elif action == "editar":
        # Solicitar al usuario que elija qué tipo de objeto editar
        object_type = input("¿Qué desea editar? (refrigerador, televisión, celular): ")

        # Definir la mutación y los datos requeridos de acuerdo al tipo de objeto elegido
        if object_type == "refrigerador":
            mutation_name = "editRefrigerator"
            object_id = input("ID del refrigerador a editar: ")
            mutation_input = {
                "id": object_id,
                "name": input("Nuevo nombre del refrigerador (deje en blanco para no cambiar): "),
                "brand": input("Nueva marca del refrigerador (deje en blanco para no cambiar): "),
                "color": input("Nuevo color del refrigerador (deje en blanco para no cambiar): "),
                "category": "REFRIGERATOR"
            }
        elif object_type == "televisión":
            mutation_name = "editTV"
            object_id = input("ID de la televisión a editar: ")
            mutation_input = {
                "id": object_id,
                "name": input("Nuevo nombre de la televisión (deje en blanco para no cambiar): "),
                "brand": input("Nueva marca de la televisión (deje en blanco para no cambiar): "),
                "color": input("Nuevo color de la televisión (deje en blanco para no cambiar): "),
                "category": "TV"
            }
        elif object_type == "celular":
            mutation_name = "editCellphone"
            object_id = input("ID del celular a editar: ")
            mutation_input = {
                "id": object_id,
                "name": input("Nuevo nombre del celular (deje en blanco para no cambiar): "),
                "brand": input("Nueva marca del celular (deje en blanco para no cambiar): "),
                "color": input("Nuevo color del celular (deje en blanco para no cambiar): "),
                "category": "CELLPHONE"
            }
        else:
            print("Tipo de objeto no válido.")
            menu()

    elif action == "eliminar":
        # Solicitar al usuario que elija qué tipo de objeto eliminar
        object_type = input("¿Qué desea eliminar? (refrigerador, televisión, celular): ")

        # Definir la mutación y los datos requeridos de acuerdo al tipo de objeto elegido
        if object_type == "refrigerador":
            mutation_name = "removeRefrigerator"
            object_id = input("ID del refrigerador a eliminar: ")
            mutation_input = object_id
        elif object_type == "televisión":
            mutation_name = "removeTV"
            object_id = input("ID de la televisión a eliminar: ")
            mutation_input = object_id
        elif object_type == "celular":
            mutation_name = "removeCellphone"
            object_id = input("ID del celular a eliminar: ")
            mutation_input = object_id
        else:
            print("Tipo de objeto no válido.")
            menu()
        
        # Definir el cuerpo de la solicitud GraphQL para la accion de eliminar
        query = '''
        mutation %s {
            %s(id: %s)
        }
        ''' % (mutation_name, mutation_name, '"'+mutation_input+'"')

        variables = {}

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
        menu()

    elif action == "listar":
        # Solicitar al usuario que elija qué tipo de objetos listar
        object_type = input("¿Qué desea listar? (refrigeradores, televisores, celulares): ")

        # Definir la query de acuerdo al tipo de objetos elegido
        if object_type == "refrigeradores":
            query_name = "refrigerators"
        elif object_type == "televisores":
            query_name = "televisions"
        elif object_type == "celulares":
            query_name = "cellphones"
        else:
            print("Tipo de objeto no válido.")
            menu()

        # Definir el cuerpo de la solicitud GraphQL
        query = '''
        query {
          %s {
            id
            name
            brand
            color
            category
          }
        }
        ''' % query_name

        variables = {}

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
        menu()

    elif action == "salir":
        exit()

    else:
        print("Acción no válida.")
        menu()

    # Definir el cuerpo de la solicitud GraphQL para las acciones de agregar y editar
    query = '''
    mutation %s($input: %sMutation!) {
      %s(input: $input) {
        id
        name
        brand
        color
        category
      }
    }
    ''' % (mutation_name, mutation_name[0].upper() + mutation_name[1:], mutation_name)

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
    menu()

# Ejecutar el menú
menu()
