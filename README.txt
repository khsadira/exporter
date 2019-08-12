########
#README#
########

A general exporter to prometheus
I'll keep it upload with what my work will ask for

____AT THE MOMENT ONLY 24x7 AND MOZILLA_OBSERVATORY ARE IMPLEMENTED____

INSTALLATION:
If you need env variable like CLIENT_ID and similar stuff you have to set them like this
CLIENT_ID_XX
CLIENT_SECRET_XX
REFRESH_TOKEN_XX

XX is the name of the import structure, like for 24x7 it's CLIENT_SECRET_24x7 etc..
(you can find all these information by reading get_XXtoken.go)


-without docker:
1: make re
2: ./exporter


-with docker:
1: make re
2: make docker
3: docker run -d -p 8080:8080 -e CLIENT_ID_XX exporter
(you can put as much as you need env var by using -e VAR)


Then you can use it easily by using your navigator at localhost:8080 or `ip`:8080 if you want to ask from another computer


khsadira
