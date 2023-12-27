from helpers import db
from models.music import Music


def get_music(title):
    return db.session.query(Music).filter(Music.title == title).first()


def get_music_from_id(id):
    return Music.query.get(id)


def add_music(Music):
    db.session.add(Music)
    db.session.commit()


def update_music(Music):
    existing_Music = get_music_from_id(Music.id)
    existing_Music.title = Music.title
    existing_Music.artistId = Music.artistId
    existing_Music.albumId = Music.albumId
    existing_Music.genreId = Music.genreId
    
    db.session.commit()


def delete_Music(id):
    db.session.delete(get_music_from_id(id))
    db.session.commit()
