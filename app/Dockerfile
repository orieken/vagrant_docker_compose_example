FROM node:0.12.5-onbuild

COPY . /src

RUN cd /src; npm install

EXPOSE 3000
CMD ["npm", "start"]