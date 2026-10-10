[instructions.md](https://github.com/user-attachments/files/33133270/instructions.md)
# Linux Installation Fest 2026
*October 10* | *2:00 PM Onwards* | *Nalanda NR111* 

<img width="1336" height="890" alt="image" src="https://gist.github.com/user-attachments/assets/0445a555-4426-448b-b236-53ec0687e0c4" />

So you have finally decided to get started on the journey of using a GNU/Linux distribution! Not sure which distro to choose? Here’s a quick summary of features to help you find the perfect fit!

**Ubuntu**\
If you’re new to Linux, Ubuntu is the perfect starting point. It’s easy to use and incredibly versatile, with a huge software library, a friendly community to help you out, and even comes with built-in gaming capabilities. Its also the easiest distro to set up dual booting with!

**Fedora**\
If you’re a developer or love having the latest tech, Fedora is the perfect choice! It’s always up-to-date with cutting-edge features, making testing a breeze. Plus, it offers a smooth installation process and supports secure booting for added peace of mind.

**Linux Mint**\
Switching from Windows? Linux Mint makes the transition a breeze with its familiar interface! It's stable and minimalist, so everything just works with fewer errors, letting you enjoy your experience right away.

> [!TIP]
> You can dual boot Linux alongside your Windows installation with no difficulties. We will help you set up dual boot on your laptops.

## Workshop Requirements

- **Download the latest version** of any one of the distributions of your choice: 
  - [Ubuntu 26.04 LTS](https://ubuntu.com/download/desktop)
  - [Fedora 43](https://www.fedoraproject.org/workstation/download)
  - [Linux Mint 22.2](https://linuxmint.com/download.php)
- **Download and install [Ventoy](https://www.ventoy.net/en/download.html) on your systems**: Ventoy is a tool that can be used to create bootable USB drives. It is recommended to download and install this software on your computer before the workshop.
- **Make sure you have 30 GB of empty SSD/HDD space**: Before installing a Linux distribution, it is essential to make sure that you have enough free space on your system. Having at least 30 GB of empty space on your SSD/HDD is recommended.
- **Bring a pen drive**: It is recommended to bring a pen drive with a capacity of at least 8 GB to create a bootable USB drive. Note that your pen drive will be formatted/erased, so back up any critical data to prevent further inconveniences.
- **Disable Secure Boot and Fast Boot**: Follow the below instructions to disable Secure Boot and Fast Boot, for a smooth installation process.

### Firmware and System Settings

Make sure you're familiar with how to access your laptop's firmware (BIOS) settings. When you start the computer, before your computer boots, pressing a special key (usually F2, F10, F12 or Del) opens the firmware settings. Once you are in your firmware settings, follow the on-screen instructions to navigate, and:

- Disable **Secure Boot**: Some Linux distributions may not install properly if Secure Boot is enabled.
- Disable **Fast Boot**: This prevents Windows from locking your drive, allowing smoother installation.

## How to Prevent Windows PIN Lockout When Toggling Secure Boot / BIOS Settings
When modifying system firmware settings (such as disabling Secure Boot for Linux dual-boot installations), Windows Hello PIN authentication often fails because the Trusted Platform Module (TPM) detects a boot integrity change.

To prevent getting locked out of Windows during future hardware or BIOS changes, follow these steps before making changes in UEFI/BIOS.

### 1. Remove PIN & Enable Password Sign-In

1. Open **Settings** by pressing `Win` + `I`.
2. Navigate to **Accounts** $\rightarrow$ **Sign-in options**.
3. Under **Additional settings**, toggle **OFF** the setting:  
   *"For improved security, only allow Windows Hello sign-in for Microsoft accounts on this device"*.
4. Expand **PIN (Windows Hello)** and click **Remove**.
5. Confirm removal when prompted.

> **Why this works:** Removing the PIN forces Windows to default to standard account password authentication, which is stored in the OS account database rather than locked inside the TPM PCR registers.

### 2. Suspend or Disable BitLocker Drive Encryption

If BitLocker or Windows Device Encryption is enabled, the TPM seal binds both the drive encryption key and Windows Hello secrets to the exact hardware/boot state.

1. Open the Start menu, search for **Manage BitLocker** (or **Device Encryption**), and open it.
2. Click **Suspend protection** (or turn off Device Encryption).
3. Confirm by clicking **Yes**.

> **Note:** Suspending BitLocker keeps protection temporarily disengaged across reboots without requiring a full decryption process, allowing seamless BIOS/Secure Boot toggles.

### Recommended Workflow for Installing Dual-Boot Linux

1. **In Windows:** Suspend BitLocker $\rightarrow$ Remove PIN $\rightarrow$ Ensure account password is known.
2. **In UEFI/BIOS:** Disable Secure Boot $\rightarrow$ Adjust Boot Priority to USB.
3. **Boot Linux USB:** Complete your partitioning and installation.
4. **Boot back into Windows:** Re-enable Windows Hello PIN if desired.
- Disable **Automatic Updates**: Temporarily disable Windows auto-updates to avoid interruptions while setting up dual boot.
- Disable **BitLocker**: Search for BitLocker in settings and then turn it off. **It may take a long time**.


