FROM node:lts

WORKDIR /usr/src/app

COPY . .

RUN npm install -g @angular/cli
RUN npm install  # Use this instead of relying on local `node_modules`

EXPOSE 4280

CMD ["ng", "serve", "--host", "0.0.0.0", "--port", "4280"]
