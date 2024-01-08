from marshmallow import Schema, fields, validates_schema, ValidationError


# Schéma genre de sortie (renvoyé au front)
class AlbumSchema(Schema):
    
    id = fields.String(description="UUID")
    name = fields.String(description="Name")
    artistId = fields.String(description="ArtistId")
    
    @staticmethod
    def is_empty(obj):
        return (not obj.get("id") or obj.get("id") == "") and \
               (not obj.get("name") or obj.get("name") == "") and \
               (not obj.get("artistId") or obj.get("artistId") == "")
    

# Schéma genre de modification (name)
class AlbumUpdateSchema(AlbumSchema):
    # permet de définir dans quelles conditions le schéma est validé ou nom
    @validates_schema
    def validates_schemas(self, data, **kwargs):
        if not (("name" in data and data["name"] != "") or ("artistId" in data and data["artistId"] != "")):
            raise ValidationError("at least one of ['name', 'artistId'] must be specified")
        