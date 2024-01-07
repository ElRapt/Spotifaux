import json
import requests
from flask import jsonify
from marshmallow import EXCLUDE
from schemas.genre import GenreSchema
from models.http_exceptions import *

genres_url = "http://localhost:8081/genres/"  # URL de l'API genre (golang)



def get_genre(id):
    response = requests.request(method="GET", url=genres_url+id)
    if response.status_code != 200:
        return response.json(), response.status_code
    return response.json(), 200

def get_genres():
    response = requests.request(method="GET", url=genres_url)
    return response.json(), response.status_code

def create_genre(genre_register):
    genre_schema = GenreSchema().loads(json.dumps(genre_register), unknown=EXCLUDE)

    # on crée le genre côté API genre
    response = requests.request(method="POST", url=genres_url, json=genre_schema)
    if response.status_code != 201:
        return response.json(), response.status_code


def modify_genre(id, genre_update):
    genre_schema = GenreSchema().loads(json.dumps(genre_update), unknown=EXCLUDE)
    response = requests.request(method="PUT", url=genres_url+id, json=genre_schema)
    if response.status_code != 200:
        return response.json(), response.status_code
    return response.json(), 200
    
def delete_genre(id):
    response = requests.request(method="DELETE", url=genres_url+id)
    if response.status_code != 204:
        return response.json(), response.status_code
    return "", 204