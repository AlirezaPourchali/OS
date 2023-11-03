# Inodes


* inodes are numbers that are set to a file in linux fs and every file has 1 inode 

* inodes contain some metadata like size , permissions , the actual location on disk and other info

* `stat <file>` will give you those data

* inode 0 -> NULL

* inode 1 -> badblock

* inode 2 -> start of the fs hierarchy

Those files mentioned in the questions are mountpoints for other filesystems and thus they have `inode 1`  , they are files on another fs.    

/dev and /sys are virtual fs managed by kernel and they are not actually on-disk      

find all mountpoints with `findmnt` , all of them have inode 1.   



