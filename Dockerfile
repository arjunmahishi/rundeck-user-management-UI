FROM jordan/rundeck:latest

RUN mkdir -p /home/rundeck-userman/ui

ADD ./rundeck-userman /home/rundeck-userman/.
ADD ui /home/rundeck-userman/ui/.

RUN chmod +x /home/rundeck-userman/rundeck-userman
RUN /home/rundeck-userman/rundeck-userman

EXPOSE 8000 
EXPOSE 4440