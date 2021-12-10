# docker-exer


### step 0 :installation

### ### install docker
sudo apt update 
sudo apt install docker.io 

### ### install docker-compose
curl -L "https://github.com/docker/compose/releases/download/1.29.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
chmod +x /usr/local/bin/docker-compose
ln -s /usr/local/bin/docker-compose /usr/bin/docker-compose

### step 0 done

### #########################################################

### step 1 :build-run

### run apps from file
sudo chmod 775 main
./main
visit via browser 127.0.0.1:5000 or curl 127.0.0.1:5000

### run apps from docker
- Build
sudo docker build -t api-demo:1.0.0 .
- Run
sudo docker run -p "5001:5000" -e MESSAGE='halo' -e PORT='5000' api-demo:1.0.0
visit via browser 127.0.0.1:5001 or curl 127.0.0.1:5002

### run apps from docker-compose
- Run
sudo docker-compose up
visit via browser 127.0.0.1:5002 or curl 127.0.0.1:5002

### step 1 :done

### #########################################################

### step 2:pull-run 

### run apps from docker-compose
- Run
sudo docker-compose up
visit via browser 127.0.0.1:9000\

ACCESS_KEY=minio
SECRET_KEY=minio123

### step 2:done 
