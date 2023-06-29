FROM node:18
WORKDIR /app

COPY package.json ./
# COPY package.json yarn.lock ./
RUN yarn install

CMD [ "yarn", "run", "dev" ]