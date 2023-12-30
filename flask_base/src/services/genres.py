import json
import requests
from sqlalchemy import exc
from marshmallow import EXCLUDE
from flask_login import current_user

from schemas.genre import GenreSchema
from models.genre import Genre as GenreModel
from models.http_exceptions import *
import repositories.genres as genre_repository


genres_url = "http://localhost:8081/genres"  # URL de l'API genre (golang)


def get_genre(id):
    response = requests.request(method="GET", url=genres_url+id)
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
        genre_model.id = response.json()["id"]
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
