# Kasir API

A simple Point of Sale (POS) API built with Go (Golang) and PostgreSQL. This API manages products and product categories.

## Prerequisites

- Go 1.20+
- PostgreSQL

## Setup

1.  **Clone the repository**
2.  **Install dependencies**:
    ```bash
    go mod tidy
    ```
3.  **Database Setup**:
    Create a PostgreSQL database and run the following SQL to create the necessary tables:

    ```sql
    CREATE TABLE categories (
        id SERIAL PRIMARY KEY,
        name VARCHAR(100) NOT NULL,
        description TEXT
    );

    CREATE TABLE products (
        id SERIAL PRIMARY KEY,
        name VARCHAR(100) NOT NULL,
        price INTEGER NOT NULL,
        stock INTEGER NOT NULL,
        category_id INTEGER,
        CONSTRAINT fk_category
            FOREIGN KEY(category_id)
            REFERENCES categories(id)
            ON DELETE SET NULL
    );
    ```

4.  **Environment Variables**:
    Create a `.env` file in the root directory:
    ```env
    PORT=8080
    DB_CONN=postgres://user:password@localhost:5432/dbname?sslmode=disable
    ```

5.  **Run the application**:
    ```bash
    go run main.go
    # or with Air for live reloading
    air
    ```

## API Endpoints

### Products

| Method | Endpoint | Description | Body (JSON) |
| :--- | :--- | :--- | :--- |
| `GET` | `/api/product` | Get all products | - |
| `POST` | `/api/product` | Create a new product | `{"name": "...", "price": 1000, "stock": 10, "category_id": 1}` |
| `GET` | `/api/product/{id}` | Get product by ID | - |
| `PUT` | `/api/product/{id}` | Update product by ID | `{"name": "...", "price": 1000, "stock": 10, "category_id": 1}` |
| `DELETE` | `/api/product/{id}` | Delete product by ID | - |

### Categories

| Method | Endpoint | Description | Body (JSON) |
| :--- | :--- | :--- | :--- |
| `GET` | `/api/category` | Get all categories | - |
| `POST` | `/api/category` | Create a new category | `{"name": "...", "description": "..."}` |
| `GET` | `/api/category/{id}` | Get category by ID | - |
| `PUT` | `/api/category/{id}` | Update category by ID | `{"name": "...", "description": "..."}` |
| `DELETE` | `/api/category/{id}` | Delete category by ID | - |
| `GET` | `/api/category_detail/{name}/detail` | Get category details with products | - |

### Example Responses

**GET /api/category_detail/Minuman/detail**

```json
{
  "id": 1,
  "name": "Minuman",
  "products": [
    {
      "id": 10,
      "name": "Es Teh",
      "price": 3000,
      "stock": 50,
      "category_id": 1
    },
    {
      "id": 11,
      "name": "Kopi",
      "price": 5000,
      "stock": 20,
      "category_id": 1
    }
  ]
}
```
