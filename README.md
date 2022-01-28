# Icecast.xml Generator

This tool was created with the objective of facilitating the creation of *icecast.xml* files. This tool was created to internally operate the [perl19/icecast2](https://github.com/Felyp-Henrique/icecast2-docker) Docker image. 

The core of the application is simple and lightweight, the artifact generated in the compilation is small and therefore drastically reduces the size of the Docker image, unlike using Python or Ruby to generate the configurations. 

[!["Buy Me A Coffee"](https://www.buymeacoffee.com/assets/img/custom_images/yellow_img.png)](https://www.buymeacoffee.com/felyphenrique)

## Usage

For now it is only possible to create *icecast.xml* files through options, in the future maybe this will change and we will not only have this option, and maybe we can use **icegen** to manage the files and make tasks a little more difficult.

There is only one command called **new**, you can see it using the **help** command:

```bash
$ ./icegen 
icecast.xml generator

Usage:
   [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  new         new icecast.xml

Flags:
  -h, --help   help for this command

Use " [command] --help" for more information about a command.

```

### Command **new**

|Options        |Default value  |Description                                    |
|---------------|---------------|-----------------------------------------------|
|dirtemplates   |./templates    |Path of dir templates                          |
|numclients     |1000           |Set the limit of connection in your server     |
|numsources     |1              |Set the limit of source mount in your server   |
|queue          |524288         |Set queue size                                 |
|clitimeout     |30             |Timeout of client                              |
|hdrtimeout     |15             |Timeout of header                              |
|srctimeout     |10             |Timeout of source                              |
|burst          |65535          |Set burst size                                 |
|srcpass        |hackme         |Set the passoword of user source               |
|admin          |admin          |Set the name of administrator                  |
|adminpass      |hackme         |Set the password of administrator              |
|host           |localhost      |Set the hostname of server                     |
|port           |8000           |Set the port of server                         |

## Install by bin

Binaries available in the *releases* tab. [Go now and download](https://github.com/Felyp-Henrique/icegen/releases).

### Unix like

```bash
unzip icegen-<OS_TYPE>.zip

chmod +x icegen

icegen --help
```

### Windows

```powershell
# extract the files and run 
icegen.exe --help
```

## Install by source

First install the **Go language** and the **build-essential** package (if on Linux distributions). 

### Unix like

```bash
# with go get
go get "https://github.com/Felyp-Henrique/icegen.git"

# or with git clone
mkdir -p "$HOME/go/src/Felyp-Henrique"

git clone "https://github.com/Felyp-Henrique/icegen.git"

# and build

cd "$HOME/go/src/Felyp-Henrique/icegen"

make
```

### Windows

```powershell
# clone inside $GOPATH and access the icegen directory 

go get
go build
```

