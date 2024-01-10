import json
from flask import Blueprint, request
from flask_login import login_required
from marshmallow import ValidationError

from helpers.content_negotiation import content_negotiation
from models.http_exceptions import *
from schemas.artist import ArtistUpdateSchema
from schemas.errors import *
import services.artists as artists_service

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
              schema: ArtistSchema
            application/yaml:
              schema: ArtistSchema
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
    return content_negotiation(*artists_service.get_artist(id))
  
@artists.route('/', methods=['GET'])
@login_required
def get_artists():
    """
    ---
    get:
      description: Getting all artists
      responses:
        '200':
          description: Ok
          content:
            application/json:
              schema: ArtistSchema
            application/yaml:
              schema: ArtistSchema
        '401':
          description: Unauthorized
          content:
            application/json:
              schema: Unauthorized
            application/yaml:
              schema: Unauthorized
      tags:
          - artists
    """
    return content_negotiation(*artists_service.get_artists())
  
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
              schema: ArtistSchema
            application/yaml:
              schema: ArtistSchema
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
        return content_negotiation(error, error.get("code"))

    return content_negotiation(*artists_service.modify_artist(id, artist_update))
  

@artists.route('/', methods=['POST'])
@login_required
def post_artist():
    """
    ---
    post:
      description: Create a artist
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
            schema: NewArtistSchema
          application/yaml:
            schema: NewArtistSchema
      responses:
        '201':
          description: Created
          content:
            application/json:
              schema: artist
            application/yaml:
              schema: artist
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
          - artists
    """
    try:
        artist_create = ArtistUpdateSchema().loads(request.data)
    except ValidationError as err:
        error = UnprocessableEntitySchema().loads(json.dumps(err.messages))
        return content_negotiation(error, error.get("code"))

    return content_negotiation(*artists_service.create_artist(artist_create))




