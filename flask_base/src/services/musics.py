import json
import requests
from sqlalchemy import exc
from marshmallow import EXCLUDE
from flask_login import current_user

from schemas.music import MusicSchema
from models.http_exceptions import *

musics_url = "http://localhost:8081/musics/"  # URL de l'API music (golang)

def music_exists(music_id):
    response = requests.request(method="GET", url=musics_url+music_id)
    if response.status_code == 200:
        return True
    elif response.status_code == 422:
        raise UnprocessableEntity
    return False

def get_music(id):
    from src.services.ratings import get_ratings

    response = requests.request(method="GET", url=musics_url+id)
    if response.status_code != 200:
        return response.json(), response.status_code

    body = response.json()
    body["ratings"], _ = get_ratings(id)

    return body, response.status_code

def get_musics():
    from src.services.ratings import get_ratings
    response = requests.request(method="GET", url=musics_url)
    if response.status_code != 200:
        return response.json(), response.status_code

    body = response.json()

    for music in body:
        music["ratings"], _ = get_ratings(music["id"])

    return body, response.status_code



def create_music(music_register):
    music_schema = MusicSchema().loads(json.dumps(music_register), unknown=EXCLUDE)

    # on crée la musique côté API music
    response = requests.request(method="POST", url=musics_url, json=music_schema)
    if response.status_code != 201:
        return response.json(), response.status_code
    return response.json(), 201



def modify_music(id, music_update):
    music_schema = MusicSchema().loads(json.dumps(music_update), unknown=EXCLUDE)

    # on modifie la musique côté API music
    response = requests.request(method="PUT", url=musics_url+id, json=music_schema)
    if response.status_code != 200:
        return response.json(), response.status_code


def delete_music(id):
    response = requests.request(method="DELETE", url=musics_url+id)
    if response.status_code != 204:
        return response.json(), response.status_code
    return "", 204
