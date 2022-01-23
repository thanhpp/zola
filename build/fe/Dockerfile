FROM node:16-alpine

WORKDIR /client
COPY client/package.json .

RUN npm install

COPY ./client/public public/
COPY ./client/src src/
COPY ./client/.env.local .env.local
COPY ./client/.env.production .env.production

EXPOSE 3000

CMD ["npm", "start"]

