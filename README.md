# Note

- Added database service as a extra works (to store the result).

- If you can not connect to the database (cloud.mongodb.com). Please change DNS in `/etc/resolv.conf` to `8.8.8.8` (for Linux)
- If it still failed, you can comment the `init()` method in the main.go and type `0` for storeDB option input

