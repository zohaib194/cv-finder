FROM node:14.17 AS ANGULAR_BUILD
RUN npm install -g @angular/cli@12.2.6
RUN npm install -g yarn
COPY frontend /frontend
WORKDIR frontend
RUN yarn install && yarn build

FROM golang:1.17.1-alpine AS GO_BUILD
COPY backend /backend
WORKDIR /backend
RUN go build -o /go/bin/backend

FROM alpine:3.10
WORKDIR app
COPY --from=ANGULAR_BUILD /frontend/dist/webapp/* ./frontend/dist/webapp/
COPY --from=GO_BUILD /go/bin/backend ./
RUN ls
CMD ./backend