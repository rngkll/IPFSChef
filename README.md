# IPFSChef
IPFS stream

Install ffmpeg
mkdir stream

## Diagram
```
+---------+     +--------------+     +------------+
|  Input  +---->+  Transcoder  +---->+  IPFS add  |
+---------+     +-------+------+     +------+-----+
                        |                   |
                        |                   v
                        |            +------+---------------+
                        +----------->+  Manifest generator  |
                                     +----------+-----------+
                                                |
                                                v
                                     +----------+------------+
                                     |  ipfs pubsub publish  |
                                     +-----------------------+
```




Made with :rainbow: by JáquerEspeis
