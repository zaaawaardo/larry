# Larry 📁

**Larry** is a free and open-source file organizer written in Go.

Let **Larry** (from *The Amazing World of Gumball*) organize your files into neat folders based on their extensions

---

## Installation

* **Build from source**

```bash
git clone <repository-url>
cd larry
go build -o larry
```

Optionally, install it system-wide:

```bash
sudo cp larry /usr/local/bin/
```

---

## Usage

```bash
larry <directory>
```

Example:

```bash
larry ~/Downloads
```

Larry will show you the planned changes and ask for confirmation before moving any files.

---

## Features

* 📂 Organizes files by extension.
* ⚙️ Automatically creates a default configuration on first launch.
* 📁 Creates missing destination folders.
* 👀 Previews all changes before applying them.
* 🎨 Colored terminal output.
* ❓ Unknown file types are moved to an `other` folder.

---

## Configuration

Larry stores its configuration in:

```text
~/.config/larry/config.json
```

You can edit this file to customize how extensions are grouped.

---

## License

Larry is free and open source.

Feel free to use it, modify it, and contribute to the project.

---

 
