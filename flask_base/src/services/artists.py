import json
import requests
from sqlalchemy import exc
from marshmallow import EXCLUDE
from flask_login import current_user

from schemas.artist import ArtistSchema
from models.http_exceptions import *



artists_url = "http://localhost:8081/artists/"  # URL de l'API artist (golang)


def get_artist(id):
    response = requests.request(method="GET", url=artists_url+id)
    if response.status_code != 200:
        return response.json(), response.status_code
    return response.json(), 200

def get_artists():
    response = requests.request(method="GET", url=artists_url)
    return response.json(), response.status_code

def create_artist(artist_create):
    artist_schema = ArtistSchema().loads(json.dumps(artist_create), unknown=EXCLUDE)
    # on crée l'utilisateur côté API artist
    response = requests.request(method="POST", url=artists_url, json=artist_schema)
    if response.status_code != 201:
        return response.json(), response.status_code



def modify_artist(id, artist_update):
    artist_schema = ArtistSchema().loads(json.dumps(artist_update), unknown=EXCLUDE)

    # on modifie l'utilisateur côté API artist
    response = requests.request(method="PUT", url=artists_url+id, json=artist_schema)
    if response.status_code != 200:
        return response.json(), response.status_code

    
