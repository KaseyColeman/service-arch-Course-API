# Base image
FROM nginx

# Get rid of the default configuration file
RUN rm /etc/nginx/conf.d/default.conf

# Use our own configuration file
COPY nginx.conf /etc/nginx/conf.d/default.conf