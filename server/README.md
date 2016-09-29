# Acksin Server

## Database Tables

```sql
CREATE TABLE acksin_machines (
    username VARCHAR(255),
    name VARCHAR(255),
    id VARCHAR(255),
    created_at TIMESTAMP DEFAULT now()
);
```
