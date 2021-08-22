# ZzEmu
Z-80 zilog emulator in go <p>
ps:
<i>My only intention here is to learn the language go, using Z80 processor.</i><p>
It was my first language, when I was very, very young.<p>
Zilog assembler was a elegant language for more civilized times. ( :D )


## Setup
### Directory and git
```bash
# Download dependencies
sudo apt install golang git gitk meld vim tree

# setup git (used to download in github)
git config --global user.name 'John doe'
git config --global user.email 'johnDoe@gmail.com'
git config --global color.branch auto
git config --global color.diff auto
git config --global color.interactive auto
git config --global color.status auto
git config --global core.editor 'vim'
git config --global credential.helper 'cache --timeout=3600'

# Thirth part library
mkdir -p /home/develop/Projetos/golib

## Create project directory (optional, replaced by call '. ./enable_develop.sh' ib shell command line)
# mkdir -p /home/develop/Projetos/go/src
# mkdir -p /home/develop/Projetos/go/bin
# mkdir -p /home/develop/Projetos/go/pkg
```

## Envoirement files setup 
Edit the file ~/.bashrc and add to tail
```file 
export GOPATH=/home/develop/Projetos/golib
export PATH=$PATH:GOPATH/bin
## replaced by bash enable_develop.sh
# export GOPATH=$GOPATH:/home/develop/Projetos/go
```

Update envoirement
```bash
# exec the batch
source ./.bashrc
```

## Z80 dependencias para teste e extencoes do VScode
```bash
apt install pasmo

code --install-extension golang.go
code --install-extension GrapeCity.gc-excelviewer
code --install-extension mborik.z80-macroasm
code --install-extension mohsen1.prettify-json
code --install-extension QiaoJie.binary-viewer
```

# to create pathfix
go mod init github.com/EduardoPagotto/ZzEmu

## Testar:
DJNZ offset
JR offset

## oderm de push pop 

push
1 - high
2 - low

pop
1 - low
2 - high
