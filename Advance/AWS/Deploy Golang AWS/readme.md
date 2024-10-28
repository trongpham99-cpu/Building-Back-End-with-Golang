1. Prepare EC2 Instance
    - Launch an EC2 instance with Ubuntu 
      - T2.micro
      - Ubuntu 18.04
      - Security Group with HTTP/HTTPS and SSH ports open
    - Ensure HTTP/HTTPS (ports 80 and 443) and SSH (port 22) are open in your security group.
2. Install Necessary Packages
    - Update the package list and install the necessary packages.
    ```bash
    sudo apt update
    sudo apt install nginx golang
    ```
3. Run Your Golang Application
- Build your Golang app:
  ```bash
  go build -o myapp
  ```
4. Configure Nginx
  - Set Up Nginx Server Block: Create a configuration file for your domain:
    ```bash
    sudo nano /etc/nginx/sites-available/<your_domain>
    ```
  - Add the following configuration (update paths as needed for your project):
    ```bash
    server {
    listen 80;
    server_name <your_domain>;

    location / {
        proxy_pass http://localhost:8080; # Replace 8080 with your Golang app port
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
      }
    }
    ```
  - Enable the configuration:
    ```bash
    sudo ln -s /etc/nginx/sites-available/<your_domain> /etc/nginx/sites-enabled/
    sudo nginx -t  # Test the configuration
    sudo systemctl reload nginx
    ```
5. Obtain SSL Certificate 
For HTTPS, use Certbot to obtain an SSL certificate:
    ```bash
    sudo apt install certbot python3-certbot-nginx
    sudo certbot --nginx -d <your_domain>
    ```

- Run your app in the background, ideally using a process manager like systemd:
  ```bash
  ./myapp &
  ```

6. Access the Application
Now, your Golang application should be accessible at your domain. You can access it via HTTP or HTTPS, depending on your configuration.