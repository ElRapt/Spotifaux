import json
from flask import Blueprint, request
from flask_login import login_required
from marshmallow import ValidationError

from models.http_exceptions import *
from schemas.artist import ArtistUpdateSchema
from schemas.errors import *
import services.artists as artist_service

# from routes import artists
artists = Blueprint(name="artists", import_name=__name__)

@artists.route('/<id>', methods=['GET'])
@login_required
def get_artist(id):
    """
    ---
    get:
      description: Getting a artist
      parameters:
        - in: path
          name: id
          schema:
            type: uuidv4
          required: true
          description: UUID of artist id
      responses:
        '200':
          description: Ok
          content:
            application/json:
              schema: Artist
            application/yaml:
              schema: Artist
        '401':
          description: Unauthorized
          content:
            application/json:
              schema: Unauthorized
            application/yaml:
              schema: Unauthorized
        '404':
          description: Not found
          content:
            application/json:
              schema: NotFound
            application/yaml:
              schema: NotFound
      tags:
          - artists
    """
    return artist_service.get_artist(id)
  
@artists.route('/<id>', methods=['PUT'])
@login_required
def put_artist(id):
    """
    ---
    put:
      description: Modify a artist
      parameters:
        - in: path
          name: id
          schema:
            type: uuidv4
          required: true
          description: UUID of artist id
      requestBody:
        required: true
        content:
          application/json:
            schema: ArtistUpdateSchema
          application/yaml:
            schema: ArtistUpdateSchema
      responses:
        '200':
          description: Ok
          content:
            application/json:
              schema: Artist
            application/yaml:
              schema: Artist
        '400':
          description: Bad request
          content:
            application/json:
              schema: BadRequest
            application/yaml:
              schema: BadRequest
        '401':
          description: Unauthorized
          content:
            application/json:
              schema: Unauthorized
            application/yaml:
              schema: Unauthorized
        '404':
          description: Not found
          content:
            application/json:
              schema: NotFound
            application/yaml:
              schema: NotFound
      tags:
          - artists
    """
    try:
        artist_update = ArtistUpdateSchema().loads(request.data)
    except ValidationError as err:
        error = UnprocessableEntitySchema().loads(json.dumps(err.messages))
        return error, error.get("code")

    return artist_service.modify_artist(id, artist_update)