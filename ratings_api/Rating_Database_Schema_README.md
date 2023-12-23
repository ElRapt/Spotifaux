
# Music Database Schema

## Tables and Fields

**1. Rating**
- `id`: UUID, Primary Key
- `user_id`: UUID
- `song_id`: UUID
- `rating`: INT, Not Null
- `comment`: TEXT

## Indexing and Constraints
- Primary keys (id) are indexed for efficient querying.
- Data types and constraints (like NOT NULL) are applied to ensure data integrity.

## Notes
- The API is designed to handle CRUD operations for ratings.
- The database schema focuses on the essential fields required for rating functionality.
- Regular review and updates of the schema may be necessary to accommodate new requirements.
