Current Structure

```
+-------------------+
|     main.go       |
| (entry point)     |
+-------------------+
          |
          v
+-------------------+
|    config.go      |<-------------------+
|  (Config struct)  |                    |
+-------------------+                    |
          |                              |
          v                              |
+-------------------+                    |
|    state.go       |                    |
|  (State struct)   |                    |
|  - holds Config   |                    |
|  + more soon      |                    |
+-------------------+                    |
          |                              |
          v                              |
+-------------------+                    |
|   commands.go     |                    |
| (Commands struct) |                    |
|  - Register()     |                    |
|  - Run()          |                    |
+-------------------+                    |
          |                              |
          v                              |
+-------------------+                    |
|   handlers.go     |                    |
| (HandlerLogin,    |                    |
|  etc.)            |                    |
+-------------------+                    |
                                         |
+-----------------------------------------+
|      ~/.gatorconfig.json (file)        |
|   (read/written by config.go)          |
+-----------------------------------------+
```
