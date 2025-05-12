
# Searchin - Search System

This project is designed to provide a simple and efficient search system by gathering data from the web, processing it on the backend, and displaying it via the frontend. 

## Project Structure

This application consists of three main components:

1. **Crawler**: This component scrapes data from the internet.
2. **Backend**: Handles requests and serves data from the database.
3. **Frontend**: Provides the user interface for interacting with the search system.

## Technologies

The project uses the following technologies:

- **Golang**: For building the backend API.
- **Docker**: To containerize the application for easy deployment and scalability.
- **PostgreSQL**: For storing data and interacting with the database.
- **Next.js**: For the frontend interface.

## Installation

To set up this project locally, follow these steps:

1. **Clone the repository**

```bash
git clone https://github.com/SarvarbekUP/Searchin_v1.git
cd Searchin_v1
```

2. **Set up the environment variables:**
Create a `.env` and setup like `.env.example`

3. **Run crawler, backend and frontend:**

**For crawler:**

cd crawler

```
go run main.go
```

**For backend:**

cd backend

```
go run cmd/server/main.go
```

**For frontend:**

cd fontend

```
npm run dev
```

## If you want to use docker:
1. **Clone the repository:**

```bash
git clone https://github.com/SarvarbekUP/Searchin_v1.git
cd Searchin_v1
```

2. **Set up the environment variables:**

Create a `.env` and setup like `.env.example`

3. **Run the Docker containers:**

Make sure Docker is installed on your machine. Then, build and start the containers with:

```bash
docker-compose up --build
```

4. **Access the application:**

Once the containers are up and running, access the app at [http://localhost:4000](http://localhost:4000).



## Usage

Once the application is running, follow these steps:

- Navigate to `http://localhost:4000` to interact with the search system.
- The frontend will display the results fetched from the backend, which in turn gets its data from the crawler and database.

## Contributing

Feel free to contribute to this project! Here's how you can help:

1. Fork the repository.
2. Create a new branch (`git checkout -b feature-name`).
3. Make your changes and commit them (`git commit -am 'Add new feature'`).
4. Push to the branch (`git push origin feature-name`).
5. Create a pull request.

If you encounter any bugs or have suggestions for improvements, please open an issue, and we will address it.

## License

This project is licensed under the [MIT License](LICENSE).

---

Thank you for exploring this project! Feel free to reach out if you have any questions or suggestions.
