import json

from flask import Blueprint, request
from flask_login import login_required
from marshmallow import ValidationError

from helpers.content_negotiation import content_negotiation
from models.http_exceptions import *
from schemas.errors import *
import services.ratings as ratings_service
from schemas.rating import NewRatingSchema, RatingUpdateSchema

ratings = Blueprint(name="ratings", import_name=__name__)


@ratings.route('/', methods=['GET'])
@login_required
def get_ratings(music_id):
    """
    ---
    get:
      description: Getting ratings of a music
      parameters:
        - in: path
          name: music_id
          schema:
            type: uuidv4
          required: true
          description: UUID of music id
      responses:
        '200':
          description: Ok
          content:
            application/json:
              schema: Rating
            application/yaml:
              schema: Rating
        '401':
          description: Unauthorized
          content:
            application/json:
              schema: Unauthorized
            application/yaml:
              schema: Unauthorized
        '404':
          description: Music not found
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
          - musics
          - ratings
    """
    try:
        return ratings_service.get_ratings(music_id)
    except NotFound:
        error = NotFoundSchema().loads("{}")
        return content_negotiation(error, error.get("code"))
    except UnprocessableEntity:
        error = UnprocessableEntitySchema().loads("{}")
        return content_negotiation(error, error.get("code"))
    except Exception:
        error = SomethingWentWrongSchema().loads("{}")
        return content_negotiation(error, error.get("code"))

@ratings.route('/', methods=['POST'])
@login_required
def post_rating(music_id):
    """
    ---
    post:
      description: Posting a rating
      parameters:
        - in: path
          name: music_id
          schema:
            type: uuidv4
          required: true
          description: UUID of music id
      requestBody:
        required: true
        content:
            application/json:
                schema: NewRating
      responses:
        '201':
          description: Created
          content:
            application/json:
              schema: Rating
            application/yaml:
              schema: Rating
        '401':
          description: Unauthorized
          content:
            application/json:
              schema: Unauthorized
            application/yaml:
              schema: Unauthorized
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
          - musics
          - ratings
    """

    try:
        new_rating = NewRatingSchema().loads(json_data=request.data.decode('utf-8'))
    except ValidationError as e:
        error = UnprocessableEntitySchema().loads(json.dumps({"message": e.messages.__str__()}))
        return content_negotiation(error, error.get("code"))

    try:
        return ratings_service.create_rating(new_rating, music_id)
    except NotFound:
        error = NotFoundSchema().loads("{}")
        return content_negotiation(error, error.get("code"))
    except UnprocessableEntity:
        error = UnprocessableEntitySchema().loads("{}")
        return content_negotiation(error, error.get("code"))
    except Exception:
        error = SomethingWentWrongSchema().loads("{}")
        return content_negotiation(error, error.get("code"))

@ratings.route('/<rating_id>', methods=['GET'])
@login_required
def get_rating(music_id, rating_id):
    """
    ---
    get:
      description: Getting a rating of a music
      parameters:
        - in: path
          name: music_id
          schema:
            type: uuidv4
          required: true
          description: UUID of music id
        - in: path
          name: rating_id
          schema:
            type: uuidv4
          required: true
          description: UUID of rating id
      responses:
        '200':
          description: Ok
          content:
            application/json:
              schema: Rating
            application/yaml:
              schema: Rating
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
          - musics
          - ratings
    """
    try:
        return ratings_service.get_rating(music_id, rating_id)
    except NotFound:
        error = NotFoundSchema().loads("{}")
        return content_negotiation(error, error.get("code"))
    except UnprocessableEntity:
        error = UnprocessableEntitySchema().loads("{}")
        return content_negotiation(error, error.get("code"))
    except Exception:
        error = SomethingWentWrongSchema().loads("{}")
        return content_negotiation(error, error.get("code"))

@ratings.route('/<rating_id>', methods=['DELETE'])
@login_required
def delete_rating(music_id, rating_id):
    """
    ---
    delete:
      description: Delete a rating
      parameters:
        - in: path
          name: music_id
          schema:
            type: uuidv4
          required: true
          description: UUID of music id
        - in: path
          name: rating_id
          schema:
            type: uuidv4
          required: true
          description: UUID of rating id
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
          - musics
          - ratings
    """
    try:
        return ratings_service.delete_rating(music_id, rating_id)
    except Forbidden:
        error = ForbiddenSchema().loads("{}")
        return content_negotiation(error, error.get("code"))
    except Exception:
        error = SomethingWentWrongSchema().loads("{}")
        return content_negotiation(error, error.get("code"))

@ratings.route('/<rating_id>', methods=['PUT'])
@login_required
def put_rating(music_id, rating_id):
    """
    ---
    put:
      description: Updating a rating
      parameters:
        - in: path
          name: music_id
          schema:
            type: uuidv4
          required: true
          description: UUID of music id
        - in: path
          name: rating_id
          schema:
            type: uuidv4
          required: true
          description: UUID of rating id
      requestBody:
        required: true
        content:
            application/json:
                schema: RatingUpdate
      responses:
        '200':
          description: Ok
          content:
            application/json:
              schema: Rating
            application/yaml:
              schema: Rating
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
        '500':
          description: Something went wrong
          content:
            application/json:
              schema: SomethingWentWrong
            application/yaml:
              schema: SomethingWentWrong
      tags:
          - musics
          - ratings
    """
    try:
        rating_update = RatingUpdateSchema().loads(json_data=request.data.decode('utf-8'))
    except ValidationError as e:
        error = UnprocessableEntitySchema().loads(json.dumps({"message": e.messages.__str__()}))
        return content_negotiation(error, error.get("code"))

    try:
        return ratings_service.modify_rating(music_id, rating_id, rating_update)
    except NotFound:
        error = NotFoundSchema().loads("{}")
        return content_negotiation(error, error.get("code"))
    except Forbidden:
        error = ForbiddenSchema().loads("{}")
        return content_negotiation(error, error.get("code"))
    except Exception:
        error = SomethingWentWrongSchema().loads("{}")
        return content_negotiation(error, error.get("code"))