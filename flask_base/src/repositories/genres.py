from helpers import db
from models.genre import Genre


def get_genre(name):
    return db.session.query(Genre).filter(Genre.name == name).first()

def get_genre_from_id(id):
    return Genre.query.get(id)

def add_genre(genre):
    db.session.add(genre)
    db.session.commit()
    
def update_genre(genre):    
    existing_genre = get_genre_from_id(genre.id)
    existing_genre.name = genre.name
    db.session.commit()

def delete_genre(id):
    db.session.delete(get_genre_from_id(id))
    db.session.commit()

