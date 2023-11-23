
# Music Database Schema

## Tables and Fields

**1. Music**
- `id`: Integer, Primary Key
- `title`: String
- `genreId`: Integer, Foreign Key (references Genre)
- `artistId`: Integer, Foreign Key (references Artist)
- `albumId`: Integer, Foreign Key (references Album)

**2. Artist**
- `id`: Integer, Primary Key
- `name`: String

**3. Album**
- `id`: Integer, Primary Key
- `name`: String
- `artistId`: Integer, Foreign Key (references Artist)

**4. Genre**
- `id`: Integer, Primary Key
- `name`: String

## Relationships
- **Music** has many-to-one relationships with **Artist**, **Album**, and **Genre**.
- **Album** has a many-to-one relationship with **Artist**.

## Indexing and Constraints
- Foreign keys and frequently queried fields like `name` in Artist and Album are indexed for improved search performance.
- Data types and constraints (like NOT NULL) are appropriately applied to each field.
- Additional metadata fields like `creationDate` and `lastModifiedDate` can be considered for tracking.

## Notes
- The database schema is designed to minimize redundancy and optimize querying.
- Consider using a composite key in the Music table for `title` and `albumId` if uniqueness across albums is required.
- Regular review and updates of the schema may be necessary to accommodate new requirements.
