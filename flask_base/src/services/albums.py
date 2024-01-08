import json
import requests
from flask import jsonify
from marshmallow import EXCLUDE
from schemas.album import AlbumSchema
from models.http_exceptions import *

albums_url = "http://localhost:8081/albums/"  # URL de l'API album (golang)



def get_album(id):
    response = requests.request(method="GET", url=albums_url+id)
    if response.status_code != 200:
        return response.json(), response.status_code
    return response.json(), 200


def get_albums():
    response = requests.request(method="GET", url=albums_url)
    return response.json(), response.status_code

def create_album(album_register):
    album_schema = AlbumSchema().loads(json.dumps(album_register), unknown=EXCLUDE)

    # on crée le album côté API album
    response = requests.request(method="POST", url=albums_url, json=album_schema)
    if response.status_code != 201:
        return response.json(), response.status_code
    return response.json(), 201


def modify_album(id, album_update):
    album_schema = AlbumSchema().loads(json.dumps(album_update), unknown=EXCLUDE)
    response = requests.request(method="PUT", url=albums_url+id, json=album_schema)
    if response.status_code != 200:
        return response.json(), response.status_code
    return response.json(), 200
    
def delete_album(id):
    response = requests.request(method="DELETE", url=albums_url+id)
    if response.status_code != 204:
        return response.json(), response.status_code
    return "", 204