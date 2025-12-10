Директория с разными небольшими объяснениями для глупого змия как и что работает. Своеобразная памятка для интересных особенностей языка.   
  
Сылки используемых приложений:  
https://go.microsoft.com/fwlink/?LinkID=760868 (Vsc в будущем переписать на скрипт)  
https://go.dev/dl/  
https://www.postgresql.org/download/linux/ubuntu/  
https://www.pgadmin.org/download/pgadmin-4-apt/ (Я беру оба варианта для работы через интерфейс и localhost)  
https://github.com/ohmyzsh/ohmyzsh/wiki/Installing-ZSH  
gitgub  
  
Настройка окружения змия на linux mint zara(22.2)(В будущем переделать в скрипт):  
sudo apt install postgresql ; curl -fsS https://www.pgadmin.org/static/packages_pgadmin_org.pub | sudo gpg --dearmor -o /usr/share/keyrings/packages-pgadmin-org.gpg && sudo sh -c 'echo "deb [signed-by=/usr/share/keyrings/packages-pgadmin-org.gpg] https://ftp.postgresql.org/pub/pgadmin/pgadmin4/apt/noble pgadmin4 main" > /etc/apt/sources.list.d/pgadmin4.list && apt update' && sudo apt install pgadmin4 ; apt install git ; apt install zsh  
  
sudo tar -C /usr/local -xzf ...  
export PATH=$PATH:/usr/local/go/bin  
  
Создание ssh-key для работы с github:    
ssh-keygen -t ed25519  
eval `ssh-agent`  
ssh-add ~/.ssh/id_ed25519  
touch ~/.ssh/config  
`Host github.com  
    HostName github.com  
    User git  
    IdentityFile ~/.ssh/id_ed25519  
    IdentitiesOnly yes  
    AddKeysToAgent yes`
Добавляем ключ в аккаунт  
ssh -T git@github.com
