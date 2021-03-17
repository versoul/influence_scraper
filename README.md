# Influence scraper

### Build and run

The docker container can be build with the following command:

```shell
docker build -t influence_scraper .
```

You can list all the available images by typing the following command:

```shell
docker image ls
```

Type the following command to run the docker image:

```shell
docker run -d -p 9090:9090 influence_scraper
```

To access the quiz-server open http://<ip>:9090 where <ip> is the ip address of the docker host.

###Job to be done

Build an engagement rate calculator for Instagram using Golang and Vue.js.

More specifically:
- Basic web application using Vue.js boilerplate app. No need for styling.
- The user inputs a username and receives the average engagement rate for the profile. This is a simple proof-of-concept.
- No need to handle error/edge cases.
- No need to “retry” automatically if an IG request gets blocked (more info down there), I will retry manually.
- No need to “protect” the endpoint with a key or anything.
- No need to “cache” the data, you can fetch from IG in real-time.
- Shared via Github (keep the repo public). Will be pulled and run on my local machine. Please separate the frontend and backend in different folders, but keep as a monorepo.
- Backend and scraper built in Golang. It should expose an API endpoint in the format of your choice.

Examples
- You can try our own calculator here as an example. I repeat: you don’t need any styling.
- https://www.instagram.com/thetwobohemians/ -> 3.04%
- https://www.instagram.com/amoureuxdumonde/ -> 6.39%

Helpful information
- Engagement Rate = (likes + comments) / followers
- - Is calculated on a “per post” basis. You can get the average by averaging the engagement rate for 12 most recent posts.
- - Please display as a percentage.
- The latest 12 posts for a user can be found directly in the source HTML of a profile’s page (i.e https://www.instagram.com/thetwobohemians/). Look for the field “edge_owner_to_timeline_media”. You don’t need to render the page. You don’t need to have a logged in profile to fetch this information.
- As long as you are using a residential internet, IG should not block you. If you use any datacenter proxy, you will get blocked immediately. If your internet gets blocked, I can provide a residential proxy.
- - IG doesn’t throttle or block a request: it will redirect you to the login page when you get blocked.

