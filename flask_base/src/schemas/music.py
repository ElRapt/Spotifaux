from marshmallow import Schema, fields, validates_schema, ValidationError


# Schéma musique de sortie (renvoyé au front)
class MusicSchema(Schema):
    id = fields.String(description="UUID")
    title = fields.String(description="Title")
    artistId = fields.String(description="ArtistId")
    genreId = fields.String(description="GenreId")
    albumId = fields.String(description="AlbumId")
    
    
    @staticmethod
    def is_empty(obj):
        return (not obj.get("id") or obj.get("id") == "") and \
               (not obj.get("title") or obj.get("title") == "") and \
               (not obj.get("artistId") or obj.get("artistId") == "") and \
               (not obj.get("genreId") or obj.get("genreId") == "") and \
               (not obj.get("albumId") or obj.get("albumId") == "")


class BaseMusicSchema(Schema):
    title = fields.String(description="Title")
    artistId = fields.String(description="ArtistId")
    genreId = fields.String(description="GenreId")
    albumId = fields.String(description="AlbumId")
    
class NewMusicSchema(BaseMusicSchema):
    @validates_schema
    def validates_schemas(self, data, **kwargs):
        if "title" not in data or data["title"] == "" or \
                "artistId" not in data or data["artistId"] == "" or \
                "genreId" not in data or data["genreId"] == "" or \
                "albumId" not in data or data["albumId"] == "":
            raise ValidationError("['title','artistId','genreId','albumId'] must all be specified")

# Schéma musique de modification (name, username, password)
class MusicUpdateSchema(MusicSchema):
    # permet de définir dans quelles conditions le schéma est validé ou nom
    @validates_schema
    def validates_schemas(self, data, **kwargs):
        if not (("title" in data and data["title"] != "") or
                ("artistId" in data and data["artistId"] != "") or
                ("genreId" in data and data["genreId"] != "") or
                ("albumId" in data and data["albumId"] != "")):
            raise ValidationError("at least one of ['title','artistId','genreId','albumId'] must be specified")
