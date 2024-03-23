# Use the official PostgreSQL image as the base image
FROM postgres:alpine

# Environment variables
ENV POSTGRES_USER=myuser
ENV POSTGRES_PASSWORD=mypassword
ENV POSTGRES_DB=mydatabase

# Expose PostgreSQL default port
EXPOSE 5432
