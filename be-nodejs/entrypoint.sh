#!/bin/bash

# Menjalankan migrasi database Sequelize
sequelize db:migrate

# Menjalankan aplikasi
exec "$@"