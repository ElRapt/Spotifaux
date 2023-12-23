from marshmallow import Schema, fields, validates_schema, ValidationError
from schemas.user import BaseUserSchema


# Schéma utilisateur de connexion
class UserLoginSchema(Schema):
    username = fields.String(description="Username", required=True)
    password = fields.String(description="Password", required=True)


# Schéma utilisateur d'enregistrement
class UserRegisterSchema(BaseUserSchema):
    # permet de définir dans quelles conditions le schéma est validé ou nom
    @validates_schema
    def validates_schemas(self, data, **kwargs):
        if "email" not in data or data["email"] == "" or \
                "username" not in data or data["username"] == "" or \
                "password" not in data or data["password"] == "":
            raise ValidationError("['email','username','password'] must all be specified")
