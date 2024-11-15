FROM node:alpine
WORKDIR /app

COPY **/package.json **/package-lock.json ./

RUN npm i 
COPY . .
EXPOSE 8091
CMD ["tail", "-f", "/dev/null"]