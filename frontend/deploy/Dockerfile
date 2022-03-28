FROM node:alpine
 
WORKDIR /app
 
COPY package*.json ./
 
RUN npm install
 
COPY . /app
 
EXPOSE 3333
 
CMD [ "npm", "run", "dev" ]