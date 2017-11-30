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