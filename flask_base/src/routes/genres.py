import json
from flask import Blueprint, request
from flask_login import login_required
from marshmallow import ValidationError

from models.http_exceptions import *
from schemas.genre import GenreUpdateSchema
from schemas.errors import *
import services.genres as genres_service

# from routes import genres
genres = Blueprint(name="genres", import_name=__name__)

@genres.route('/<id>', methods=['GET'])
@login_required
def get_genre(id):
    """
    ---
    get:
      description: Getting a genre
      parameters:
        - in: path
          name: id
          schema:
            type: uuidv4
          required: true
          description: UUID of genre id
      responses:
        '200':
          description: Ok
          content:
            application/json:
              schema: Genre
            application/yaml:
              schema: Genre
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
          - genres
    """
    return genres_service.get_genre(id)
  
@genres.route('/<id>', methods=['PUT'])
@login_required
def put_genre(id):
    """
    ---
    put:
      description: Modify a genre
      parameters:
        - in: path
          name: id
          schema:
            type: uuidv4
          required: true
          description: UUID of genre id
      requestBody:
        required: true
        content:
          application/json:
            schema: GenreUpdate
          application/yaml:
            schema: GenreUpdate
      responses:
        '200':
          description: Ok
          content:
            application/json:
              schema: Genre
            application/yaml:
              schema: Genre
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
          - genres
    """
    try:
        genre_update = GenreUpdateSchema().loads(request.data)
    except ValidationError as err:
        error = UnprocessableEntitySchema().loads(json.dumps(err.messages))
        return error, error.get("code")

    return genres_service.modify_genre(id, genre_update)
  
  
@genres.route('/', methods=['POST'])
@login_required
def post_genre():
    """
    ---
    post:
      description: Create a genre
      parameters:
        - in: path
          name: id
          schema:
            type: uuidv4
          required: true
          description: UUID of genre id
      requestBody:
        required: true
        content:
          application/json:
            schema: GenreUpdate
          application/yaml:
            schema: GenreUpdate
      responses:
        '201':
          description: Created
          content:
            application/json:
              schema: Genre
            application/yaml:
              schema: Genre
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
        '409':
          description: Conflict
          content:
            application/json:
              schema: Conflict
            application/yaml:
              schema: Conflict
      tags:
          - genres
    """
    try:
        genre_create = GenreUpdateSchema().loads(request.data)
    except ValidationError as err:
        error = UnprocessableEntitySchema().loads(json.dumps(err.messages))
        return error, error.get("code")

    return genres_service.create_genre(genre_create)




