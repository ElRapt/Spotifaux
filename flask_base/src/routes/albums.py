import json
from flask import Blueprint, request
from flask_login import login_required
from marshmallow import ValidationError

from helpers.content_negotiation import content_negotiation
from models.http_exceptions import *
from schemas.album import AlbumUpdateSchema
from schemas.errors import *
import services.albums as albums_service

# from routes import albums
albums = Blueprint(name="albums", import_name=__name__)

@albums.route('/<id>', methods=['GET'])
@login_required
def get_album(id):
    """
    ---
    get:
      description: Getting a album
      parameters:
        - in: path
          name: id
          schema:
            type: uuidv4
          required: true
          description: UUID of album id
      responses:
        '200':
          description: Ok
          content:
            application/json:
              schema: album
            application/yaml:
              schema: album
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
          - albums
    """
    return content_negotiation(*albums_service.get_album(id))
  
@albums.route('/', methods=['GET'])
@login_required
def get_albums():
    """
    ---
    get:
      description: Getting all albums
      responses:
        '200':
          description: Ok
          content:
            application/json:
              schema: album
            application/yaml:
              schema: album
        '401':
          description: Unauthorized
          content:
            application/json:
              schema: Unauthorized
            application/yaml:
              schema: Unauthorized
      tags:
          - albums
    """
    return content_negotiation(*albums_service.get_albums())
  
@albums.route('/<id>', methods=['PUT'])
@login_required
def put_album(id):
    """
    ---
    put:
      description: Modify a album
      parameters:
        - in: path
          name: id
          schema:
            type: uuidv4
          required: true
          description: UUID of album id
      requestBody:
        required: true
        content:
          application/json:
            schema: AlbumUpdateSchema
          application/yaml:
            schema: AlbumUpdateSchema
      responses:
        '200':
          description: Ok
          content:
            application/json:
              schema: album
            application/yaml:
              schema: album
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
          - albums
    """
    try:
        album_update = AlbumUpdateSchema().loads(request.data)
    except ValidationError as err:
        error = UnprocessableEntitySchema().loads(json.dumps(err.messages))
        return content_negotiation(error, error.get("code"))

    return content_negotiation(*albums_service.update_album(id, album_update))
  

@albums.route('/', methods=['POST'])
@login_required
def post_album():
    """
    ---
    post:
      description: Create a album
      parameters:
        - in: path
          name: id
          schema:
            type: uuidv4
          required: true
          description: UUID of album id
      requestBody:
        required: true
        content:
          application/json:
            schema: AlbumUpdateSchema
          application/yaml:
            schema: AlbumUpdateSchema
      responses:
        '201':
          description: Created
          content:
            application/json:
              schema: album
            application/yaml:
              schema: album
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
          - albums
    """
    try:
        album_create = AlbumUpdateSchema().loads(request.data)
    except ValidationError as err:
        error = UnprocessableEntitySchema().loads(json.dumps(err.messages))
        return content_negotiation(error, error.get("code"))

    return content_negotiation(*albums_service.create_album(album_create))




