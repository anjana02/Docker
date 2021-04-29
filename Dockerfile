# sample Dockerfile for reference

# specify the base image

FROM nginx:alpine

# copy the path

COPY . /usr/share/nginx/html
