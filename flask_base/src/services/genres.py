import json
import requests
import uuid
from marshmallow import EXCLUDE
from schemas.genre import GenreSchema
from models.genre import Genre as GenreModel
from models.http_exceptions import *
import repositories.genres as genre_repository


genres_url = "http://localhost:8081/genres"  # URL de l'API genre (golang)


def get_genre(id):
    try:
        response = requests.get(url=f"{genres_url}/{id}")
        # Check if the response status code indicates success
        if response.status_code == 200:
            return response.json(), response.status_code
        else:
            # Handle non-successful responses
            error_message = f"Error fetching genre with ID {id}: {response.text}"
            return {"error": error_message}, response.status_code
    except requests.RequestException as e:
        # Handle any requests-related exceptions
        error_message = f"Request failed: {str(e)}"
        return {"error": error_message}, 500


def get_genres():
    response = requests.request(method="GET", url=genres_url)
    return response.json(), response.status_code

def create_genre(genre_register):
    # on récupère le modèle utilisateur pour la BDD
    genre_model = GenreModel.from_dict(genre_register)
    # on récupère le schéma utilisateur pour la requête vers l'API genre
    genre_schema = GenreSchema().loads(json.dumps(genre_register), unknown=EXCLUDE)

    # on crée l'utilisateur côté API genre
    response = requests.request(method="POST", url=genres_url, json=genre_schema)
    if response.status_code != 201:
        return response.json(), response.status_code

    # on ajoute l'utilisateur dans la base de données
    # pour que les données entre API et BDD correspondent
    try:
        genre_repository.add_genre(genre_model)
    except Exception:
        raise SomethingWentWrong

    return response.json(), response.status_code



def modify_genre(id, genre_update):
    # on récupère le modèle utilisateur pour la BDD
    genre_model = GenreModel.from_dict(genre_update)
    # on récupère le schéma utilisateur pour la requête vers l'API genre
    genre_schema = GenreSchema().loads(json.dumps(genre_update), unknown=EXCLUDE)

    # on modifie l'utilisateur côté API genre
    response = requests.request(method="PUT", url=genres_url+id, json=genre_schema)
    if response.status_code != 200:
        return response.json(), response.status_code

    # on modifie l'utilisateur dans la base de données
    # pour que les données entre API et BDD correspondent
    try:
        genre_model.id = response.json()["id"]
        genre_repository.update_genre(genre_model)
    except Exception:
        raise SomethingWentWrong

    return response.json(), response.status_code
    


def get_genre_from_db(name):
    return genre_repository.get_genre(name)


def genre_exists(name):
    return get_genre_from_db(name) is not None

def delete_genre(id):
    try:
        response = requests.delete(url=f"{genres_url}/{id}")

        # Check if the response status code indicates success
        if response.status_code in [200, 204]:  # 204 is typically used for successful DELETE requests
            return response.json() if response.status_code == 200 else {}, response.status_code
        else:
            # Handle non-successful responses
            error_message = f"Error deleting genre with ID {id}: {response.text}"
            return {"error": error_message}, response.status_code
    except requests.RequestException as e:
        # Handle any requests-related exceptions
        error_message = f"Request failed: {str(e)}"
        return {"error": error_message}, 500
