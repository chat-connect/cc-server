# Run Production
FROM node:18.12.0 AS production

WORKDIR /swagger
RUN apt-get update && apt-get install -y
RUN yarn global add @stoplight/prism-cli
EXPOSE 9000

# Run Development
FROM node:18.12.0 AS development

WORKDIR /swagger
RUN apt-get update && apt-get install -y
RUN yarn global add @stoplight/prism-cli
EXPOSE 9000
