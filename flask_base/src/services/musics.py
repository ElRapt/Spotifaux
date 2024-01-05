import json
import requests
from sqlalchemy import exc
from marshmallow import EXCLUDE
from flask_login import current_user

from schemas.music import MusicSchema
from models.music import Music as MusicModel
from models.http_exceptions import *
import repositories.musics as music_repository


musics_url = "http://localhost:8081/musics"  # URL de l'API music (golang)


def get_music(id):
    response = requests.request(method="GET", url=musics_url+id)
    return response.json(), response.status_code

def get_musics():
    response = requests.request(method="GET", url=musics_url)
    return response.json(), response.status_code


def create_music(music_register):
    # on récupère le modèle utilisateur pour la BDD
    music_model = MusicModel.from_dict(music_register)
    # on récupère le schéma utilisateur pour la requête vers l'API music
    music_schema = MusicSchema().loads(json.dumps(music_register), unknown=EXCLUDE)

    # on crée l'utilisateur côté API music
    response = requests.request(method="POST", url=musics_url, json=music_schema)
    if response.status_code != 201:
        return response.json(), response.status_code

    # on ajoute l'utilisateur dans la base de données
    # pour que les données entre API et BDD correspondent
    try:
        music_model.id = response.json()["id"]
        music_repository.add_music(music_model)
    except Exception:
        raise SomethingWentWrong

    return response.json(), response.status_code



def modify_music(id, music_update):
    # on récupère le modèle utilisateur pour la BDD
    music_model = MusicModel.from_dict(music_update)
    # on récupère le schéma utilisateur pour la requête vers l'API music
    music_schema = MusicSchema().loads(json.dumps(music_update), unknown=EXCLUDE)

    # on modifie l'utilisateur côté API music
    response = requests.request(method="PUT", url=musics_url+id, json=music_schema)
    if response.status_code != 200:
        return response.json(), response.status_code

    # on modifie l'utilisateur dans la base de données
    # pour que les données entre API et BDD correspondent
    try:
        music_model.id = response.json()["id"]
        music_repository.update_music(music_model)
    except Exception:
        raise SomethingWentWrong

    return response.json(), response.status_code
    


def get_user_from_db(username):
    return music_repository.get_user(username)


def user_exists(username):
    return get_user_from_db(username) is not None
