# Windows Installation Guide

## Installing Python

You can download python 3 directly from the [Microsoft Store](https://www.microsoft.com/en-us/p/python-38/9mssztt1n39l?activetab=pivot:overviewtab).

To check whether it is install open Windows Powershell and running the following command.

```
python3
```

If the python console opens then you have successfully installed it.

[Alternate Guide without using Microsoft Store](https://matthewhorne.me/how-to-install-python-and-pip-on-windows-10/)

## Installing Flask

Installing python from the Microsoft store will automatically install pip3 after that you can just open Powershell and install flask by running the following command,

```
pip3 install flask
```

## Installing Postman

You can install Postman from [here](https://www.postman.com/downloads/).

## Installing Visual Studio Code

You can install Visual Studio code from [here](https://code.visualstudio.com/Download).

# Linux Installation Guide

## Installing Pip

Python is already installed in linux to install pip you need to run the following commands,

### To install pip on [Ubuntu](https://linuxconfig.org/ubuntu-linux-download), [Debian](https://linuxconfig.org/debian-linux-download), and [Linux Mint](https://linuxconfig.org/linux-mint-download)

```
sudo apt install python3-pip
```

### To install pip on [CentOS 8 (and newer)](https://linuxconfig.org/centos-linux-download), [Fedora](https://linuxconfig.org/fedora-linux-download), and [Red Hat](https://linuxconfig.org/red-hat-linux-download)

```
sudo dnf install python3		
```

### To install pip on CentOS 6 and 7, and older versions of Red Hat

```
sudo yum install epel-release
sudo yum install python-pip
```

### To install pip on [Arch Linux](https://linuxconfig.org/arch-linux-download) and [Manjaro](https://linuxconfig.org/manjaro-linux-download)

```
sudo pacman -S python-pip
```

### To install pip on [OpenSUSE](https://linuxconfig.org/opensuse-linux-download)

```
sudo zypper install python3-pip
```

## Installing Flask

After installing pip3 after that you can just open Terminal and install flask by running the following command,

```
pip3 install flask
```

## Installing Visual Studio Code

To install Visual Studio Code on Linux visit  [this site.](https://code.visualstudio.com/docs/setup/linux)

## Installing Postman

To install Postman open Terminal and run the following commands,

#### Install Postman in Debian and Ubuntu

```
sudo apt update
sudo apt install snapd
sudo snap install postman
```

#### Install Postman in Linux Mint

```
sudo rm /etc/apt/preferences.d/nosnap.pref
sudo apt update
sudo apt install snapd
sudo snap install postman
```

#### Install Postman in Fedora Linux

```
sudo dnf install snapd
sudo ln -s /var/lib/snapd/snap /snap
sudo snap install postman
```
