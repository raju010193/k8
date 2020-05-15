FROM node:10.13.0 as build-stage
WORKDIR /app
run npm install -g @angular/cli
COPY package*.json /app/
RUN npm install
COPY ./ /app/
ARG configuration=production
RUN ng build -- --output-path=./dist/out

# Stage 1, based on Nginx, to have only the compiled app, ready for production with Nginx
FROM nginx:1.15
#Copy ci-dashboard-dist
COPY --from=build-stage /app/dist/out/ /usr/share/nginx/html
#Copy default nginx configuration
COPY ./nginx-custom.conf /etc/nginx/conf.d/default.conf
