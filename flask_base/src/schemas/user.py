from marshmallow import Schema, fields, validates_schema, ValidationError


# Schéma utilisateur de sortie (renvoyé au front)
class UserSchema(Schema):
    id = fields.String(description="UUID")
    inscription_date = fields.DateTime(description="Inscription date")
    username = fields.String(description="Username")
    email = fields.String(description="Email")
    
    @staticmethod
    def is_empty(obj):
        return (not obj.get("id") or obj.get("id") == "") and \
               (not obj.get("username") or obj.get("username") == "") and \
               (not obj.get("email") or obj.get("email") == "") and \
               (not obj.get("inscription_date") or obj.get("inscription_date") == "")


class BaseUserSchema(Schema):
    password = fields.String(description="Password")
    username = fields.String(description="Username")
    email = fields.String(description="Email")


# Schéma utilisateur de modification (name, username, password)
class UserUpdateSchema(BaseUserSchema):
    # permet de définir dans quelles conditions le schéma est validé ou nom
    @validates_schema
    def validates_schemas(self, data, **kwargs):
        if not (("email" in data and data["email"] != "") or
                ("username" in data and data["username"] != "") or
                ("password" in data and data["password"] != "")):
            raise ValidationError("at least one of ['email','username','password'] must be specified")
