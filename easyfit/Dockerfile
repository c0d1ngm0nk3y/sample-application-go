FROM openjdk:latest
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN ls -la
EXPOSE 8080
CMD ["java", "-jar", "fitnesse-20150814-standalone.jar", "-p", "8080"]
