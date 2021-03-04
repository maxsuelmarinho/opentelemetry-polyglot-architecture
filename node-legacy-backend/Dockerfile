FROM node:12-alpine
ENV NODE_ENV=production
RUN mkdir -p /home/node/app/node_modules && chown -R node:node /home/node/app
WORKDIR /home/node/app
COPY package*.json ./
USER node
RUN npm ci --only=${NODE_ENV}
COPY --chown=node:node . .
EXPOSE 8080
CMD [ "node", "server.js" ]
