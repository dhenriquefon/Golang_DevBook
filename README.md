# Golang DevBook

API e APP de rede social desenvolvida utilizando golang.

Todo codigo desenvolvido fazendo o curso: https://www.udemy.com/course/aprenda-golang-do-zero-desenvolva-uma-aplicacao-completa/

## Instalação

Como dependencia é necessário subir um banco de dados my sql.

### MYSQL SETUP

Instalar mysql e executar seguintes comandos

```unix
-- dando o start
sudo systemctl start mysql

-- habilitar caso queira que o mysql start sempre que ligar a maquina
sudo systemctl enable mysql

-- logar no banco de dados como usuario root
sudo mysql -u root -p

-- criar base de dados
CREATE DATABASE devbook;

-- usar uma database
USE devbook

-- criando usuario
CREATE USER 'golang'@'localhost' IDENTIFIED BY 'golang';

-- grants pros usuarios

GRANT ALL PRIVILEGES ON devbook.* TO 'golang'@'localhost';
```

## Build da Aplicaço

### Golang DevBook - SETUP

Executar script:

``bash
-- criacao de tabelas
/sql/sql.sql

-- popular tabelas
/sql/dados.sql
```

### Golang DevBook - SETUP

``bash
-- build api
cd /api/
go build

-- rodar api
cd /api/
./devbook-api

-- build webapp
cd /webapp/
go build

-- rodar api
cd /webapp/
./webapp
```


## Usando a Aplicação

Realizar login em:

http://localhost:3000/

usuario: usuario1
senha: 123

Ou pode criar seu usuario e começar a usar
