import json
from flask import Blueprint, request
from flask_login import login_required
from marshmallow import ValidationError

from helpers.content_negotiation import content_negotiation
from models.http_exceptions import *
from schemas.music import MusicUpdateSchema
from schemas.errors import *
import services.musics as musics_service

# from routes import musics
musics = Blueprint(name="musics", import_name=__name__)


@musics.route('/', methods=['GET'])  
@login_required
def get_musics():
    """
    ---
    get:
      description: Getting all musics
      responses:
        '200':
          description: Ok
          content:
            application/json:
              schema: MusicSchema
            application/yaml:
              schema: MusicSchema
        '401':
          description: Unauthorized
          content:
            application/json:
              schema: Unauthorized
            application/yaml:
              schema: Unauthorized
      tags:
          - musics
    """
    return content_negotiation(*musics_service.get_musics())

@musics.route('/<id>', methods=['GET'])
@login_required
def get_music(id):
    """
    ---
    get:
      description: Getting a music
      parameters:
        - in: path
          name: id
          schema:
            type: uuidv4
          required: true
          description: UUID of music id
      responses:
        '200':
          description: Ok
          content:
            application/json:
              schema: MusicSchema
            application/yaml:
              schema: MusicSchema
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
          - musics
    """
    return content_negotiation(*musics_service.get_music(id)) 



@musics.route('/<id>', methods=['PUT'])
@login_required
def put_music(id):
    """
    ---
    put:
      description: Modify a music
      parameters:
        - in: path
          name: id
          schema:
            type: uuidv4
          required: true
          description: UUID of music id
      requestBody:
        required: true
        content:
            application/json:
                schema: MusicUpdateSchema
      responses:
        '200':
          description: Ok
          content:
            application/json:
              schema: Music
            application/yaml:
              schema: Music
        '401':
          description: Unauthorized
          content:
            application/json:
              schema: Unauthorized
            application/yaml:
              schema: Unauthorized
        '403':
          description: Forbidden
          content:
            application/json:
              schema: Forbidden
            application/yaml:
              schema: Forbidden
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
      tags:
          - musics
    """
    # parser le body
    try:
        music_update = MusicUpdateSchema().loads(request.data)
    except ValidationError as err:
        error = UnprocessableEntitySchema().loads(json.dumps(err.messages))
        return content_negotiation(error, error.get("code"))

    return content_negotiation(*musics_service.modify_music(id, music_update))
  
@musics.route('/<id>', methods=['DELETE'])
@login_required
def delete_music(id):
    """
    ---
    delete:
      description: Delete a music
      parameters:
        - in: path
          name: id
          schema:
            type: uuidv4
          required: true
          description: UUID of music id
      responses:
        '204':
          description: No content
        '401':
          description: Unauthorized
          content:
            application/json:
              schema: Unauthorized
            application/yaml:
              schema: Unauthorized
        '403':
          description: Forbidden
          content:
            application/json:
              schema: Forbidden
            application/yaml:
              schema: Forbidden
        '404':
          description: Not found
          content:
            application/json:
              schema: NotFound
            application/yaml:
              schema: NotFound
      tags:
          - musics
    """
    return content_negotiation(*musics_service.delete_music(id))
  
@musics.route('/', methods=['POST'])
@login_required
def post_music():
    """
    ---
    post:
      description: Create a music
      requestBody:
        required: true
        content:
            application/json:
                schema: MusicUpdateSchema
      responses:
        '201':
          description: Created
          content:
            application/json:
              schema: NewMusicSchema
            application/yaml:
              schema: NewMusicSchema
        '401':
          description: Unauthorized
          content:
            application/json:
              schema: Unauthorized
            application/yaml:
              schema: Unauthorized
        '403':
          description: Forbidden
          content:
            application/json:
              schema: Forbidden
            application/yaml:
              schema: Forbidden
        '409':
          description: Conflict
          content:
            application/json:
              schema: Conflict
            application/yaml:
              schema: Conflict
        '422':
          description: Unprocessable entity
          content:
            application/json:
              schema: UnprocessableEntity
            application/yaml:
              schema: UnprocessableEntity
      tags:
          - musics
    """
    try:
        music_register = MusicUpdateSchema().loads(request.data)
    except ValidationError as err:
        error = UnprocessableEntitySchema().loads(json.dumps(err.messages))
        return content_negotiation(error, error.get("code"))

    return content_negotiation(*musics_service.create_music(music_register))