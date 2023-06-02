apt-get -y update &&
apt-get -y install certbot &&
apt-get -y install python3-certbot-nginx &&
apt-get -y install cron &&
crontab -l |
{ cat; echo "0 12 * * * /usr/bin/certbot renew --quiet"; } |
crontab - &&
crontab -l
# Please run below  command to run the .sh file
# chmod +x sslSetUp.sh