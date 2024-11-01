# Installing Golang in WSL Ubuntu
`wget https://dl.google.com/go/go1.23.2.linux-amd64.tar.gz`<br>
`sudo tar -xvf go1.18.3.linux-amd64.tar.gz`<br>
`sudo mv go /usr/local`<br>
`echo "export GOROOT=/usr/local/go" >> ~/.bashrc`<br>
`echo "export GOPATH=\$HOME/go" ~/.bashrc`<br>
`echo "export PATH=\$GOPATH/bin:\$GOROOT/bin:\$PATH" >> ~/.bashrc`<br>
`source ~/.bashrc`<br>