## Windows Subsystem for Linux Installation Guide for Windows 10

- Windows users can install WSL on their machines. Linux commands run well on WSL Terminal. Follow [this guide](https://docs.microsoft.com/en-us/windows/wsl/install-win10)
(WSL Ubuntu Distro is prefered)

## Vim - Installation

- ## On Debian based Linux 
Execute the commands below in terminal - 

```
$ sudo apt-get update 
$ sudo apt-get install vim
```

Check the installation of vim by runing the **which** command:

```
$ which vim
```

It should print the location of Vim binary, which should look something like this `/usr/bin/vim`

- ## On RPM based Linux
Execute the commands below in terminal - 

```
$ su - 
$ yum install vim
```

Check the installation of vim by runing the **which** command:

```
$ which vim
```

It should print the location of Vim binary, which should look something like this `/usr/bin/vim`

- ## On Windows
  - Follow the same instructions as above but on a WSL terminal depending on the Linux distribution. (Ubuntu Distro is prefered. Follow Debain based Linux installation instructions)

## VSCode installation on WSL

- First, install [Remote - WSL](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-wsl) extension on your VSCode on Windows. 
- Run the following command to install VSCode remote server

  ```
  $ code .
  ```
- After the above step, start VSCode in the desired path in WSL by `code .`
