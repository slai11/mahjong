# mahjong
[![Netlify Status](https://api.netlify.com/api/v1/badges/6633457d-79a8-4074-b460-a131fa8dc73d/deploy-status)](https://app.netlify.com/sites/tableswim/deploys)

A minimalist online multiplayer mahjong game for friends.

A hosted version of this app is available on [www.tableswim.netlify.app](https://tableswim.netlify.app/).

## Design 
* player makes decision -> minimal computer intervention
* no rules imposed (computer doesnt count points)

## Deployment
Frontend is deployed on Netlify under free tier. Backend is deployed via
docker-compose files in the smallest DigitalOcean droplet there is. 

References: https://www.digitalocean.com/community/tutorials/how-to-deploy-a-go-web-application-with-docker-and-nginx-on-ubuntu-18-04

## Contributing
All sorts of contributions are welcomed, in particular, styling (CSS/html) and
bugfix. 
