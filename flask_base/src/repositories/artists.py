from helpers import db
from models.artist import Artist


def get_artist(name):
    return db.session.query(Artist).filter(Artist.name == name).first()

def get_artist_from_id(id):
    return Artist.query.get(id)

def add_artist(artist):
    db.session.add(artist)
    db.session.commit()
    
def update_artist(artist):    
    existing_artist = get_artist_from_id(artist.id)
    existing_artist.name = artist.name
    db.session.commit()

def delete_artist(id):
    db.session.delete(get_artist_from_id(id))
    db.session.commit()

