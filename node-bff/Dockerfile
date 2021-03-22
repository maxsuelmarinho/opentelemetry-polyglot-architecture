FROM node:12-alpine
#ENV NODE_ENV=production
RUN mkdir -p /home/node/app/node_modules && chown -R node:node /home/node/app
WORKDIR /home/node/app
COPY package*.json ./
USER node
#RUN npm ci --only=${NODE_ENV}
RUN npm i
COPY --chown=node:node . .
RUN npm run prod:build
EXPOSE 8080
#CMD [ "npm", "start" ]
CMD node ./dist/server.js
