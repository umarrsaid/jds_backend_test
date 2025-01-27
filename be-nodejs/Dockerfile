# Use official Node.js image as base
FROM node:16

# Set working directory
WORKDIR /app

# Copy package.json and package-lock.json
COPY package*.json ./

# Install dependencies
RUN npm install
RUN npm install -g sequelize-cli

# Copy the rest of the application code
COPY . .

# Run database migrations and seeders
RUN sequelize db:create
RUN sequelize db:migrate
RUN sequelize db:seed:all

# Expose port 3000
EXPOSE 3000

# Run the app
CMD ["npm", "start"]