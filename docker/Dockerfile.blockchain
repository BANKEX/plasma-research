FROM node:latest

#RUN apk add --update --no-cache make linux-headers netcat-openbsd git

RUN apt-get update && apt-get install -y netcat

WORKDIR /app/
ADD ./src/contracts/package.json ./src/contracts/package-lock.json ./src/contracts/Makefile /app/
RUN make node_modules
ADD ./src/contracts /app/
RUN mkdir -p /app/data/
RUN ./scripts/build.sh

EXPOSE 8545

CMD ./scripts/run.sh