disk_usage
==========

disk_usage is a utility that will walk down from a starting location on a filesystem and return 
pertinent usage data. This includes, space of the partition, inode data, upto the 10 largest 
cumulative directories, and upto the 20 largest files.

This is the version of disk_usage written in go.

### Installing disk_usage


  * Using Go:
    * Make sure you have [go](http://golang.org/) installed (for Mac users I strongly recommend HomeBrew: just `brew install go`.
    * Set your $GOPATH and $GOBIN variables.

    * Build from source:

      ```shell
      git clone git@github.com:josh-paul/go_disk_usage.git
      cd go_disk_usage
      go build -o $GOBIN/disk_usage
      $GOBIN/disk_usage TARGET
      ```
  * Download binary
    * Download correct binary from github releases (https://github.com/josh-paul/go_disk_usage/releases/)
    * Save file as disk_usage in a location in your system path and add executable permisions.

### Sample Output

```
$ disk_usage /var
86.79% available disk space on /var
Total: 67 GB, Used: 8.9 GB, Free: 58 GB

94.47% of total inodes are free.
Total: 4194304, Used: 231801, Free: 3962503

Total directory count of 275
The 10 largest directories are:
Size   Directory
185 MB /var/cache/apt/archives
154 MB /var/lib/apt/lists
92 MB  /var/cache/apt
38 MB  /var/lib/dpkg/info
13 MB  /var/log
7.9 MB /var/lib/app-info/icons/ubuntu-artful-universe/64x64
5.7 MB /var/lib/mlocate
5.5 MB /var/cache/cracklib
5.3 MB /var/cache/debconf
4.7 MB /var/backups

Total file count of 9401
The 20 largest files are:
Size   Modified                  File
47 MB  2017-10-19T05:56:35-07:00 /var/lib/apt/lists/us.archive.ubuntu.com_ubuntu_dists_artful_universe_binary-amd64_Packages
47 MB  2017-10-19T05:56:35-07:00 /var/lib/apt/lists/us.archive.ubuntu.com_ubuntu_dists_artful_universe_binary-i386_Packages
46 MB  2017-11-30T09:12:12-08:00 /var/cache/apt/pkgcache.bin
46 MB  2017-11-30T08:54:56-08:00 /var/cache/apt/srcpkgcache.bin
44 MB  2017-11-27T06:29:23-08:00 /var/cache/apt/archives/firefox_57.0+build4-0ubuntu0.17.10.6_amd64.deb
37 MB  2017-11-20T13:24:25-08:00 /var/cache/apt/archives/docker-ce_17.11.0~ce-0~ubuntu_amd64.deb
32 MB  2017-11-07T06:40:50-08:00 /var/cache/apt/archives/linux-image-extra-4.13.0-17-generic_4.13.0-17.20_amd64.deb
26 MB  2017-10-18T04:43:18-07:00 /var/lib/apt/lists/us.archive.ubuntu.com_ubuntu_dists_artful_universe_i18n_Translation-en
21 MB  2017-11-07T06:40:59-08:00 /var/cache/apt/archives/linux-image-4.13.0-17-generic_4.13.0-17.20_amd64.deb
11 MB  2017-11-16T05:44:03-08:00 /var/cache/apt/archives/libwebkit2gtk-4.0-37_2.18.3-0ubuntu0.17.10.1_amd64.deb
11 MB  2017-11-07T06:40:50-08:00 /var/cache/apt/archives/linux-headers-4.13.0-17_4.13.0-17.20_all.deb
8.2 MB 2017-11-07T00:08:24-08:00 /var/log/syslog.7.gz
7.9 MB 2017-10-19T03:03:14-07:00 /var/lib/apt/lists/us.archive.ubuntu.com_ubuntu_dists_artful_universe_dep11_icons-64x64.tar.gz
6.7 MB 2017-10-19T05:56:35-07:00 /var/lib/apt/lists/us.archive.ubuntu.com_ubuntu_dists_artful_main_binary-amd64_Packages
6.7 MB 2017-10-19T05:56:35-07:00 /var/lib/apt/lists/us.archive.ubuntu.com_ubuntu_dists_artful_main_binary-i386_Packages
5.7 MB 2017-11-30T09:41:46-08:00 /var/lib/mlocate/mlocate.db
5.2 MB 2017-10-18T11:39:40-07:00 /var/cache/cracklib/cracklib_dict.pwd
5.2 MB 2017-11-21T05:41:26-08:00 /var/cache/apt/archives/samba-libs_2%3a4.6.7+dfsg-1ubuntu3.1_amd64.deb
4.4 MB 2017-10-19T03:03:24-07:00 /var/lib/apt/lists/us.archive.ubuntu.com_ubuntu_dists_artful_universe_dep11_Components-amd64.yml.gz
4.1 MB 2017-11-16T05:44:02-08:00 /var/cache/apt/archives/libjavascriptcoregtk-4.0-18_2.18.3-0ubuntu0.17.10.1_amd64.deb
```
