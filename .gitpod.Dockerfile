FROM gitpod/workspace-full

# install ansible-galaxy
RUN sudo apt update && sudo apt install software-properties-common
RUN sudo add-apt-repository --yes --update ppa:ansible/ansible
RUN sudo apt install -y ansible
