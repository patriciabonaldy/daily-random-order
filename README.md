## Daily Order

**Setup:**
- Run follow command:
~~~bash
make all
~~~

### Creation a binary file

~~~bash
make build
~~~

### Testing

~~~bash
make test
~~~

### Run app

- Create a json file with the co-worker of team.
- Create a .env file and add the follow:
~~~bash
# Data Path
DATA_PATH=/path_of_json/team.json (the filename is not important)
# SLack
WEBHOOKURL=webhookurl
~~~

- run the follow command:
~~~bash
./daily-random
~~~