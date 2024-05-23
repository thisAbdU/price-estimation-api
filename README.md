## Price Estimation API

This repository contains a robust REST API developed in Golang, designed to facilitate price estimations for users. The API adheres to the CRUD (Create, Read, Update, Delete) paradigm and leverages a PostgreSQL database for persistent storage. 

### Features

- **Create**: Add new price estimation data to the database.
- **Read**: Retrieve existing price estimation data.
- **Update**: Modify existing price estimation data.
- **Delete**: Remove price estimation data from the database.

### Technologies

- **Golang**: The primary language used for building the API, chosen for its performance and concurrency capabilities.
- **PostgreSQL**: The relational database system used for persistent storage, selected for its robustness and scalability.
- **REST**: The architectural style used for designing networked applications, ensuring a stateless and scalable API.

### Getting Started

1. **Clone the repository**:
   ```bash
   git clone https://github.com/yourusername/price-estimation-api.git
   cd price-estimation-api
   ```

2. **Set up the database**:
   - Ensure PostgreSQL is installed and running.
   - Create a new database for the project.

3. **Configure environment variables**:
   - Create a `.env` file with the necessary database connection details.

4. **Run the application**:
   ```bash
   go run main.go
   ```

### API Endpoints

- `POST /estimates`: Create a new price estimate.
- `GET /estimates`: Retrieve all price estimates.
- `GET /estimates/{id}`: Retrieve a specific price estimate by ID.
- `PATCH /estimates/{id}`: Update a specific price estimate by ID.
- `DELETE /estimates/{id}`: Delete a specific price estimate by ID.
