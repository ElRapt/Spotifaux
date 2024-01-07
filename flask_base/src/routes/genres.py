import json
from flask import Blueprint, request
from flask_login import login_required
from marshmallow import ValidationError

from helpers.content_negotiation import content_negotiation
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
          description: UUID of Genre id
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
    return content_negotiation(*genres_service.get_genre(id))
  

@genres.route('/', methods=['GET'])  
@login_required
def get_genres():
    """
    ---
    get:
      description: Getting all genres
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
        '401':
          description: Unauthorized
          content:
            application/json:
              schema: Unauthorized
            application/yaml:
              schema: Unauthorized
      tags:
          - genres
    """
    return content_negotiation(*genres_service.get_genres())

@genres.route('/<id>', methods=['PUT'])
@login_required
def put_genre(id):
    """
    ---
    put:
      description: Updating a genre
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
                schema: GenreUpdateSchema
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
        '422':
          description: Unprocessable entity
          content:
            application/json:
              schema: UnprocessableEntity
            application/yaml:
              schema: UnprocessableEntity
        '500':
          description: Something went wrong
          content:
            application/json:
              schema: SomethingWentWrong
            application/yaml:
              schema: SomethingWentWrong
      tags:
          - genres
    """
    # parser le body
    try:
        genre_update = GenreUpdateSchema().loads(json_data=request.data.decode('utf-8'))
    except ValidationError as e:
        error = UnprocessableEntitySchema().loads(json.dumps({"message": e.messages.__str__()}))
        return content_negotiation(error, error.get("code"))
    
    try:
        return content_negotiation(*genres_service.update_genre(id, genre_update))
    except Exception:
        error = SomethingWentWrongSchema().loads("{}")
        return content_negotiation(error, error.get("code"))
  
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
        return content_negotiation(error, error.get("code"))

    return content_negotiation(*genres_service.create_genre(genre_create))



@genres.route('/<id>', methods=['DELETE'])
@login_required
def delete_genre(id):
    """
    ---
    delete:
      description: Delete a genre
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
    return content_negotiation(*genres_service.delete_genre(id))
  
  
  