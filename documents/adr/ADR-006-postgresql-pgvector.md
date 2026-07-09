Decision

Use PostgreSQL as the primary datastore with pgvector for semantic search.

Reason

One operational database.

Native ACID transactions.

Vector retrieval without introducing another database.

Suitable for MVP and early production.

Tradeoff

Slightly lower vector search scalability than dedicated vector databases.

Migration path

Introduce a dedicated vector database only when scale justifies the added complexity.