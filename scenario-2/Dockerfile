FROM jenkins/jenkins:lts
ARG http_port=8082
ARG agent_port=50000
ARG JENKINS_HOME=/var/jenkins_home

VOLUME $JENKINS_HOME
EXPOSE ${http_port}
# will be used by attached slave agents:
EXPOSE ${agent_port}
