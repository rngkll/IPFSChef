# IPFSChef
IPFS stream

Install ffmpeg
mkdir stream


## Diagram

  +---------+     +--------------+     +------------+
  |  Input  +---->+  Transcoder  +---->+  IPFS add  |
  +---------+     +--------------+     +------------+
                                              |
       +--------------------------------------+
       |
       | +----------------------+
       +>+  Manifest generator  |
         +----------------------+
                   |                +------------------+
                   +--------------->+  pubsub publish  |
                                    +------------------+

Made with :rainbow: by JáquerEspeis
