hashname
========

hashname generates hashes of files and renames the files.
This tool can be used to prevent duplicates in file collections.

Install with Go
---------------

    go get github.com/MatthiasSchild/hashname

Example
-------

    >> hashname --ext *.jpg
    Use hashing method: sha1
    Keep extensions: true
    IMG_3369.jpg -> 16e0f92f725e7652fc2fb256cf6fe30899047102.jpg
    IMG_3370.jpg -> 0f159d0dec228cabd7c0e8f6918715cf6c7df230.jpg
    IMG_3371.jpg -> c03dd48a31ab21b6639b2555d95667dc76e42f7a.jpg
    IMG_3372.jpg -> 26742d76483e45704054bff8f594d8e01a940513.jpg
    IMG_3373.jpg -> 6c7c4bb7b18c66cd8b18c6ed8443464fa4beebd0.jpg

Options
-------

- `method`: The hashing method to use; default is `sha1`
    - `sha1`, `sha256`, `sha512`, `md5`,
- `ext`: Keep the extension of the file
- `dry`: Execute a dry run without actually renaming the files
