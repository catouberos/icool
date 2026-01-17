# icool

the queue machine at icool previously was able to pick-up and search for any youtube links, but now, it was borked.

luckily, they instead ship us a web UI to do queueing ourself, but it for some reasons was constraint into their "white-listed" channel, and sucks at searching.

so, using their exposed unsecured websocket connection, we could queue any youtube link into the system, by just sending the video ID instead. while doing that, why not also changing the image and title ourself too?

## running this stuff

```
go run main.go --host=<ip:host> --room=<room>
```

leaving ip and host as default should be fine, but as a karaoke noobs i'm not certain, so ymmv.
