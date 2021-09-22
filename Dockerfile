FROM node:12.11 AS ANGULAR_BUILD
RUN npm install -g @angular/cli@8.3.12
RUN npm install -g yarn
COPY frontend /frontend
WORKDIR frontend
RUN yarn install && yarn build

FROM golang:1.13.1-alpine AS GO_BUILD
COPY backend /backend
WORKDIR /backend
RUN go build -o /go/bin/backend

FROM alpine:3.10
WORKDIR app
COPY --from=ANGULAR_BUILD /frontend/dist/webapp/* ./frontend/dist/webapp/
COPY --from=GO_BUILD /go/bin/server ./
RUN ls
CMD ./server