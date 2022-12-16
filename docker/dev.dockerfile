

FROM --platform=linux/amd64 ubuntu:latest 

WORKDIR /app

ADD ./home/ ./home/
RUN apt update && apt install curl -y
ADD download.sh download.sh
RUN bash download.sh

CMD ["blsd","start", "--home", "home"]s