# mahjong
[![Netlify Status](https://api.netlify.com/api/v1/badges/6633457d-79a8-4074-b460-a131fa8dc73d/deploy-status)](https://app.netlify.com/sites/tableswim/deploys)

A minimalist online multiplayer mahjong game for friends. Best played with video/voice conferencing and a dash of socialising.

A hosted version of this app is available on [www.tableswim.netlify.app](https://tableswim.netlify.app/).

## Deployment
Frontend is deployed on Netlify under free tier. Backend is deployed via
docker-compose files in the smallest DigitalOcean droplet there is. 

References: https://www.digitalocean.com/community/tutorials/how-to-deploy-a-go-web-application-with-docker-and-nginx-on-ubuntu-18-04

## Local development
To run the web app on localhost:8080, you need Vue installed.
```
cd frontend
yarn serve
```

To run the backend on localhost:80, you need go installed.
```
go build
./mahjong
```

## Contributing
All sorts of contributions are welcomed, in particular, styling (CSS/html) and
bugfix. 

Features such as score/pay-out tabulation, chat-system and payments, which reduces the need
for voice/video communication between players will not be entertained. TableSwim
Mahjong is meant for friends to meet up digitally and play mahjong.
