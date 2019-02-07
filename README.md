# rundeck-user-management-UI
An API layer and a web UI built on top of rundeck's `PropertyFileLoginModule`

# How to use
You can either download the pre-built binary or build it from this source.

- Build from source 
    ```bash
    $ git clone git@github.com:arjunmahishi/rundeck-user-management-UI.git
    $ cd rundeck-user-management-UI
    $ go build # need to have go installed
    $ sudo chmod 666 /etc/rundeck/realm.properties # Allow read/write access to the realm.properties file 
    $ ./rundeck-user-management-UI
    ```