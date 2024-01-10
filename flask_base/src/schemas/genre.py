from marshmallow import Schema, fields, validates_schema, ValidationError


# Schéma genre de sortie (renvoyé au front)
class GenreSchema(Schema):
    
    id = fields.String(description="UUID")
    name = fields.String(description="Name")
    
    @staticmethod
    def is_empty(obj):
        return (not obj.get("id") or obj.get("id") == "") and \
               (not obj.get("name") or obj.get("name") == "")
    

class BaseGenreSchema(Schema):
    name = fields.String(description="Name")


class NewGenreSchema(BaseGenreSchema):
    @validates_schema
    def validates_schemas(self, data, **kwargs):
        if "name" not in data or data["name"] == "":
            raise ValidationError("['name'] must all be specified")
        
# Schéma genre de modification (name)
class GenreUpdateSchema(GenreSchema):
    # permet de définir dans quelles conditions le schéma est validé ou nom
    @validates_schema
    def validates_schemas(self, data, **kwargs):
        if not (("name" in data and data["name"] != "")):
            raise ValidationError("at least one of ['name'] must be specified")
        