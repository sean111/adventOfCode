FROM php:8.3-cli
COPY --from=composer:latest /usr/bin/composer /usr/local/bin/composer

RUN apt update && apt full-upgrade -y && apt install git zsh -y
RUN curl --proto '=https' --tlsv1.2 -sSf https://just.systems/install.sh | bash -s -- --to /usr/bin
RUN sh -c "$(curl -fsSL https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh)"
RUN curl https://squizlabs.github.io/PHP_CodeSniffer/phpcs.phar --output /usr/bin/phpcs && chmod +x /usr/bin/phpcs

RUN mkdir /code
WORKDIR /code
CMD ["zsh"]