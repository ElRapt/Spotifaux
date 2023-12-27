from helpers import db


# modèle de donnée pour la base de donnée musique
class Music(db.Model):
    __tablename__ = 'musics'

    id = db.Column(db.String(255), primary_key=True)
    title = db.Column(db.String(255), nullable=False)
    artistId = db.Column(db.String(255), nullable=False)
    albumId = db.Column(db.String(255), nullable=False)
    genreId = db.Column(db.String(255), nullable=False)

    def __init__(self, uuid, title, artistId, genreId, albumId):
        self.id = uuid
        self.title = title
        self.genreId = genreId
        self.artistId = artistId
        self.albumId = albumId

    def is_empty(self):
        return (not self.id or self.id == "") and \
               (not self.title or self.title == "") and \
               (not self.genreId or self.genreId == "") and \
               (not self.artistId or self.artistId == "") and \
               (not self.albumId or self.albumId == "")

    @staticmethod
    def from_dict(obj):
        return Music(obj.get("id"), obj.get("title"), obj.get("genreId"), obj.get("artistId"), obj.get("albumId"))
