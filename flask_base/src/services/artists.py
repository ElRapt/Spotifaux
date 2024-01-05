import json
import requests
from sqlalchemy import exc
from marshmallow import EXCLUDE
from flask_login import current_user

from schemas.artist import ArtistSchema
from models.artist import Artist as ArtistModel
from models.http_exceptions import *
import repositories.artists as artist_repository


artists_url = "http://localhost:8081/artists"  # URL de l'API artist (golang)


def get_artist(id):
    response = requests.request(method="GET", url=artists_url+id)
    return response.json(), response.status_code

def get_artists():
    response = requests.request(method="GET", url=artists_url)
    return response.json(), response.status_code

def create_artist(artist_register):
    # on récupère le modèle utilisateur pour la BDD
    artist_model = ArtistModel.from_dict(artist_register)
    # on récupère le schéma utilisateur pour la requête vers l'API artist
    artist_schema = ArtistSchema().loads(json.dumps(artist_register), unknown=EXCLUDE)

    # on crée l'utilisateur côté API artist
    response = requests.request(method="POST", url=artists_url, json=artist_schema)
    if response.status_code != 201:
        return response.json(), response.status_code

    # on ajoute l'utilisateur dans la base de données
    # pour que les données entre API et BDD correspondent
    try:
        artist_model.id = response.json()["id"]
        artist_repository.add_artist(artist_model)
    except Exception:
        raise SomethingWentWrong

    return response.json(), response.status_code



def modify_artist(id, artist_update):
    # on récupère le modèle utilisateur pour la BDD
    artist_model = ArtistModel.from_dict(artist_update)
    # on récupère le schéma utilisateur pour la requête vers l'API artist
    artist_schema = ArtistSchema().loads(json.dumps(artist_update), unknown=EXCLUDE)

    # on modifie l'utilisateur côté API artist
    response = requests.request(method="PUT", url=artists_url+id, json=artist_schema)
    if response.status_code != 200:
        return response.json(), response.status_code

    # on modifie l'utilisateur dans la base de données
    # pour que les données entre API et BDD correspondent
    try:
        artist_model.id = response.json()["id"]
        artist_repository.update_artist(artist_model)
    except Exception:
        raise SomethingWentWrong

    return response.json(), response.status_code
    


def get_artist_from_db(name):
    return artist_repository.get_artist(name)


def artist_exists(name):
    return get_artist_from_db(name) is not None
