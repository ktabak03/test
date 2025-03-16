FROM nginx
COPY default.conf /etc/nginx/conf.d/default.conf

ARG CONTAPI

RUN sed -i "s/api-go-1/$CONTAPI/g" /etc/nginx/conf.d/default.conf 
