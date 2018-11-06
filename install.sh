sudo apt-get update

sudo apt-get install vim -y

sudo apt-get git-core -y

sudo apt-get install screen -y

git --version

wget -qO - https://download.sublimetext.com/sublimehq-pub.gpg | sudo apt-key add -

sudo apt-get install apt-transport-https -y

echo "deb https://download.sublimetext.com/ apt/stable/" | sudo tee /etc/apt/sources.list.d/sublime-text.list

sudo apt-get update

sudo apt-get install sublime-text -y

sudo tar -C /usr/local -xzf /Downloads/go1.11.1.linux-amd64.tar.gz

sudo apt-get update

sudo apt-get install postgresql postgresql-contrib -y

sudo apt-get install curl -y

sudo dpkg -i /Downloads/code_1.28.2-1539735992_amd64.deb -y

sudo apt-get install -f -y

sudo apt-get install golang-go

export PATH=$PATH:/usr/local/go/bin
