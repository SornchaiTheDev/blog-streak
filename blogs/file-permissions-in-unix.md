---
Title: File Permissions in UNIX
PublishedDate: 05/01/2025
---

Have you ever wonder what is this when you run `ls -l` command ?

```sh
total 40
drwxr-xr-x@ 3 imdev  staff    96 Sep  9 17:37 configs
drwxr-xr-x@ 3 imdev  staff    96 Sep  9 17:30 docker
drwxr-xr-x@ 8 imdev  staff   256 Sep  9 17:50 domain
-rw-r--r--@ 1 imdev  staff  2211 Sep  9 17:32 go.mod
-rw-r--r--@ 1 imdev  staff  9553 Sep  9 17:32 go.sum
drwxr-xr-x@ 3 imdev  staff    96 Sep  9 17:59 internal
-rw-r--r--@ 1 imdev  staff   909 Sep  9 18:35 main.go
drwxr-xr-x@ 4 imdev  staff   128 Sep  9 17:44 utils
```

The first part of the output (`-rw-r--r--@`) use to tell us about the file permissions.But how can I know what does it mean you may ask :3.So let's take a deeper look about what does each character mean.

> You need to know about File Ownership first 

### File Ownership

In UNIX every file and directory has 3 ownership attributes

1. User (owner) : The person who own the file
2. Group : a group of users who have shared access to the file
3. Others : All other users in the system.

> Then you need to know about Permission Types

### Permission Types

There are 3 types of action that user can interact with the file

1. Read (r)
- File : Allow to view the file
- Directory : Allow to listing files in the directory

2. Write (w)
- File : Allow to write to the file
- Directory : Allow to create, deleting, or renaming files or directories inside the directory

3. Execute (x)
- File : Allow to execute the file
- Directory : Allow to entering the directory

Now I'm going back to what I mention above `-rw-r--r--@`. You may now know what are `r` and `w` which is **Read** and **Write**.
But what about `-` and `@`.

The first character tell us what is the type of the file `-` if a file, `d` is a directory, and `l` is a symbolic link.

The last character tell us about special permission of this file which is a **Extended Attributes** (in Macos)

> Extended Attributes store metadata about the file, beyond the standard Unix file attribute such as Storing quarantine information for files downloaded from the internet, Adding custom metadata to files, such as tags or comments etc.

So now we can transcript this file permission `-rw-r--r--@` to

- `-` this is a file
- `rw-` Owner of this file can **Read** and **Write** to it.
- `r--` Users in the group can only **Read** it.
- `r--` Other users in the system can only **Read** it.
- `@` Extended Attributes of the file.

This is the end of this blog. Good night 😪
