from helpers import db


# modèle de donnée pour la base de donnée genre
class Genre(db.Model):

    __tablename__ = 'genres'

    id = db.Column(db.String(255), primary_key=True)
    name = db.Column(db.String(255), nullable=False)

    def __init__(self, uuid, name):
        self.id = uuid
        self.name = name

    def is_empty(self):
        return (not self.id or self.id == "") and \
               (not self.name or self.name == "")

    @staticmethod
    def from_dict(obj):
        return Genre(obj.get("id"), obj.get("name"))
    