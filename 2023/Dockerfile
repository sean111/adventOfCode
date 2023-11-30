FROM php:8.3-cli
COPY --from=composer:latest /usr/bin/composer /usr/local/bin/composer
RUN mkdir /code
COPY . /code
WORKDIR /code
