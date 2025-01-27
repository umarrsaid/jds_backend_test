FROM playcourt/nodejs:22-alpine

ARG NODE_ENV
ARG PORT
ARG DB_DIALECT
ARG DB_HOST
ARG DB_PORT
ARG DB_DATABASE
ARG DB_USERNAME
ARG DB_PASSWORD
ARG SECRET
ARG MINIO_URL
ARG MINIO_ACCESSKEY
ARG MINIO_SECRETKEY
ARG MINIO_BUCKET
ARG NEXTAUTH_URL
ARG NEXTAUTH_SECRET
ARG JWT_SECRET
ARG SEND_EMAIL_USER_ACCOUNT
ARG SEND_EMAIL_PASS_ACCOUNT


# Gunakan user root untuk menginstal dependensi
USER root

# Install Python dan buat symlink untuk python3
RUN apk add --no-cache python3 py3-pip && ln -sf python3 /usr/bin/python

# Pastikan izin yang benar untuk /usr/local
RUN chown -R user:root /usr/local

# Install Sequelize CLI secara global
RUN npm install -g sequelize-cli

# Ganti ke user 'user'
USER user

# Buat virtual environment untuk Python
RUN python3 -m venv /usr/src/app/.venv

# Aktifkan virtual environment dan setel variabel lingkungan
ENV PATH="/usr/src/app/.venv/bin:$PATH"
ENV VIRTUAL_ENV="/usr/src/app/.venv"

# Salin kode aplikasi dan pastikan izin yang benar
COPY --chown=user:root . /usr/src/app/

# Pastikan entrypoint.sh bisa dieksekusi
RUN ls -l /usr/src/app/entrypoint.sh # Tambahkan langkah ini untuk memeriksa keberadaan file
RUN chmod +x /usr/src/app/entrypoint.sh

# Setel direktori kerja
WORKDIR /usr/src/app

RUN touch .env

RUN echo "NODE_ENV=${NODE_ENV}" >> .env
RUN echo "PORT=${PORT}" >> .env
RUN echo "DB_DIALECT=${DB_DIALECT}" >> .env
RUN echo "DB_HOST=${DB_HOST}" >> .env
RUN echo "DB_PORT=${DB_PORT}" >> .env
RUN echo "DB_DATABASE=${DB_DATABASE}" >> .env
RUN echo "DB_USERNAME=${DB_USERNAME}" >> .env
RUN echo "DB_PASSWORD=${DB_PASSWORD}" >> .env
RUN echo "SECRET=${SECRET}" >> .env
RUN echo "MINIO_URL=${MINIO_URL}" >> .env
RUN echo "MINIO_ACCESSKEY=${MINIO_ACCESSKEY}" >> .env
RUN echo "MINIO_SECRETKEY=${MINIO_SECRETKEY}" >> .env
RUN echo "MINIO_BUCKET=${MINIO_BUCKET}" >> .env
RUN echo "NEXTAUTH_URL=${NEXTAUTH_URL}" >> .env
RUN echo "NEXTAUTH_SECRET=${NEXTAUTH_SECRET}" >> .env
RUN echo "JWT_SECRET=${JWT_SECRET}" >> .env
RUN echo "SEND_EMAIL_USER_ACCOUNT=${SEND_EMAIL_USER_ACCOUNT}" >> .env
RUN echo "SEND_EMAIL_PASS_ACCOUNT=${SEND_EMAIL_PASS_ACCOUNT}" >> .env

# Install dependensi Node.js
RUN npm install

# Ekspos port aplikasi
EXPOSE 3000

# Setel entrypoint dan perintah default
CMD ["sh", "-c", "sequelize db:migrate && npm start"]
