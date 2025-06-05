# /bin/bash
sudo apt update && sudo apt install -y nginx-full  && \
sudo systemctl start nginx && \
sudo systemctl enable nginx && \
sudo cp ~/nginx-layer4-load-balancer/nginx/nginx.conf /etc/nginx/nginx.conf && \
sudo systemctl restart nginx
