from marshmallow import Schema, fields, validates_schema, ValidationError


# Schéma artiste de sortie (renvoyé au front)
class ArtistSchema(Schema):
    id = fields.String(description="UUID")
    name = fields.String(description="Name")
    
    
    @staticmethod
    def is_empty(obj):
        return (not obj.get("id") or obj.get("id") == "") and \
               (not obj.get("name") or obj.get("name") == "")

# Schéma artiste de modification (name)
class ArtistUpdateSchema(ArtistSchema):
    # permet de définir dans quelles conditions le schéma est validé ou nom
    @validates_schema
    def validates_schemas(self, data, **kwargs):
        if not (("name" in data and data["name"] != "")):
            raise ValidationError("at least one of ['name'] must be specified")
        