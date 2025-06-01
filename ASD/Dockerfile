FROM ubuntu:20.04

# Set timezone to Asia/Seoul
ENV TZ=Asia/Seoul
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

# Set APP_HOME environment variable
ENV APP_HOME=/app/config/

# Copy the Go binary from the build output
COPY app /app/asd

# Copy the entrypoint script
COPY dist/script.sh /app/script.sh

# Set permissions for the binary and script
RUN chmod 0755 /app/asd
RUN chmod 0755 /app/script.sh

# Create config directory for APP_HOME
RUN mkdir -p /app/config/

# Set working directory
WORKDIR /app/

# Set entrypoint to the bash script
ENTRYPOINT ["/app/script.sh"]