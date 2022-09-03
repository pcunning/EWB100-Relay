# EWB100-Relay

![Image of EWB100](https://i.postimg.cc/G35NPV6f/Screen-Shot-2022-09-02-at-6-54-59-PM.png)

> The TEAM badge (EWB100) brings a new level of affordability and portability to mobile voice, providing basic Push-to-Talk (PTT) communications that enables enterprises to extend the benefits of on-the-spot voice to new workers and workgroups inside the four walls.

## How it works

The EWB100 sends multicast voice packets to `239.192.2.2:5000`. These packets are collected and published to a nats channel to be recived and rebrocast across multiple networks. By default the app uses `demo.nats.io` but you can provide a connection sting with `-c=nats://demo.nats.io:4222`.
