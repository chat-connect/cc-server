# Run Production
FROM node:18.12.0-alpine AS production

WORKDIR /swagger
RUN apk upgrade --update && apk add bash && apk --no-cache add git
RUN yarn global add @stoplight/prism-cli
EXPOSE 8000

# Run Development
FROM node:18.12.0-alpine AS development

WORKDIR /swagger
RUN apk upgrade --update && apk add bash && apk --no-cache add git
RUN yarn global add @stoplight/prism-cli
EXPOSE 8000
