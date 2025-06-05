# /bin/bash
sudo apt update && sudo apt install -y curl wget git  && \
wget https://go.dev/dl/go1.24.3.linux-amd64.tar.gz && \
sudo tar -C /usr/local -xzf go1.24.3.linux-amd64.tar.gz && \
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc && \
source ~/.bashrc && \
git clone https://github.com/OlyMahmudMugdho/nginx-layer4-load-balancer.git
