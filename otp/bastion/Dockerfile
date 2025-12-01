# Start from our Ubuntu base
FROM debian:latest

RUN apt-get update && apt-get install -y \
    openssh-server \
    sudo \
    && rm -rf /var/lib/apt/lists/*

RUN useradd -m -s /bin/bash luigi
RUN usermod -aG sudo luigi

RUN passwd -l luigi

RUN mkdir -p /home/luigi/.ssh
RUN chown luigi:luigi /home/luigi/.ssh
RUN chmod 700 /home/luigi/.ssh

COPY bastion_key.pub /home/luigi/.ssh/authorized_keys

RUN chown luigi:luigi /home/luigi/.ssh/authorized_keys
RUN chmod 600 /home/luigi/.ssh/authorized_keys


RUN echo "\n# --- Custom Secure Settings ---" >> /etc/ssh/sshd_config && \
    echo "PermitRootLogin no" >> /etc/ssh/sshd_config && \
    echo "PasswordAuthentication no" >> /etc/ssh/sshd_config && \
    echo "PubkeyAuthentication yes" >> /etc/ssh/sshd_config && \
    echo "ChallengeResponseAuthentication no" >> /etc/ssh/sshd_config

EXPOSE 22

# Start the SSH service
CMD ["/usr/sbin/sshd", "-D"]