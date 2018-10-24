sudo apt-get update

sudo apt-get install vim -y

sudo apt-get git-core -y

sudo apt-get install screen

git --version

wget -qO - https://download.sublimetext.com/sublimehq-pub.gpg | sudo apt-key add -

sudo apt-get install apt-transport-https

echo "deb https://download.sublimetext.com/ apt/stable/" | sudo tee /etc/apt/sources.list.d/sublime-text.list

sudo apt-get update

sudo apt-get install sublime-text

sudo tar -C /usr/local -xzf go1.11.1.linux-amd64.tar.gz

sudo dpkg -i code_1.28.1-1539281690_amd64.deb

sudo apt-get update

sudo apt-get install postgresql postgresql-contrib -y

sudo apt-get install curl

export PATH=$PATH:/usr/local/go/bin
